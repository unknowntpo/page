package page

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	connect "github.com/bufbuild/connect-go"
	"github.com/unknowntpo/page/domain"
	"github.com/unknowntpo/page/domain/mock"
	pb "github.com/unknowntpo/page/gen/proto/page"
	"github.com/unknowntpo/page/gen/proto/page/pageconnect"
	"github.com/unknowntpo/page/pkg/errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPageAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PageAPI")
}

// Ref: https://github.com/bufbuild/connect-demo/blob/main/main_test.go
var _ = Describe("PageAPI", Ordered, func() {
	var (
		ctrl            *gomock.Controller
		mux             *http.ServeMux
		server          *httptest.Server
		client          pageconnect.PageServiceClient
		mockPageUsecase *mock.MockPageUsecase
	)

	BeforeEach(func() {
		// Set up mock usecase
		ctrl = gomock.NewController(GinkgoT())
		mockPageUsecase = mock.NewMockPageUsecase(ctrl)

		// Start mux
		mux = http.NewServeMux()
		mux.Handle(pageconnect.NewPageServiceHandler(
			NewPageServer(mockPageUsecase),
		))
		// Start test server
		server = httptest.NewUnstartedServer(mux)
		server.EnableHTTP2 = true
		server.StartTLS()
	})

	BeforeEach(func() {
		client = pageconnect.NewPageServiceClient(
			server.Client(),
			server.URL,
			connect.WithGRPC(),
		)
	})

	AfterEach(func() {
		server.Close()
		ctrl.Finish()
	})

	Describe("NewList", func() {
		var (
			err     error
			res     *connect.Response[pb.NewListResponse]
			userID  int64
			listKey string
		)
		const (
			validListKey  = "validListKey"
			validUserID   = 33
			invalidUserID = ""
		)
		JustBeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				NewList(gomock.Any(), userID, gomock.Any()).
				DoAndReturn(func(ctx context.Context, userID int64, listKey domain.ListKey) error {
					if string(listKey) == validListKey {
						return nil
					}
					return errors.New(errors.ResourceAlreadyExist, "already exist")
				}).AnyTimes()
			res, err = client.NewList(context.Background(), connect.NewRequest(&pb.NewListRequest{
				ListKey: listKey,
				UserID:  userID,
			}))
		})
		Context("normal", func() {
			BeforeEach(func() {
				userID = 33
				listKey = validListKey
			})
			It("should return OK", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.Msg.Status).To(Equal("OK"))
			})
		})
		When("listKey is not valid", func() {
			BeforeEach(func() {
				userID = validUserID
				listKey = invalidUserID
			})
			It("should return invalid argument", func() {
				Expect(err.Error()).To(ContainSubstring(ErrInvalidListKey.Error()))
			})
		})
		When("userID not valid", func() {
			BeforeEach(func() {
				userID = 0
				listKey = validListKey
			})
			It("should return invalid argument", func() {
				Expect(err.Error()).To(ContainSubstring(ErrInvalidUserID.Error()))
			})
		})
	})

	Describe("GetHead", func() {
		var (
			err     error
			res     *connect.Response[pb.GetHeadResponse]
			userID  int64
			listKey string
		)
		BeforeEach(func() {
			userID = 33
			mockPageUsecase.
				EXPECT().
				GetHead(gomock.Any(), userID, gomock.Any()).
				// FIXME: return head by input
				Return(gomock.Any(), nil).Times(1)
			res, err = client.GetHead(context.Background(), connect.NewRequest(&pb.GetHeadRequest{
				ListKey: listKey,
				UserID:  userID,
			}))
			Expect(err).ShouldNot(HaveOccurred())
		})
		Context("normal", func() {
			It("should return expected PageKey", func() {
				Expect(res.Msg.PageKey).To(Equal(listKey))
			})
		})
	})
	When("SetPage is called", func() {
		var (
			err    error
			stream *connect.BidiStreamForClient[pb.SetPageRequest, pb.SetPageResponse]
			pages  = []domain.Page{
				{
					Content: mock.GenerateRandomString(3),
				},
				{
					Content: mock.GenerateRandomString(3),
				},
				{
					Content: mock.GenerateRandomString(3),
				},
			}
			gotPages []domain.Page
			listKey  string
		)

		const (
			userID int64 = 33
			round  int   = 3
		)

		BeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				SetPage(gomock.Any(), userID, gomock.Any(), gomock.Any()).
				DoAndReturn(func(ctx context.Context, userID int64, listKey domain.ListKey, page domain.Page) (domain.PageKey, error) {
					gotPages = append(gotPages, page)
					// FIXME: should return correct result
					return "", nil
				}).Times(round)

			stream = client.SetPage(context.Background())
		})
		Context("call stream.Send for three times", func() {
			BeforeEach(func() {
				for i := 0; i < round; i++ {
					err := stream.Send(
						&pb.SetPageRequest{
							UserID:      userID,
							ListKey:     listKey,
							PageContent: pages[i].Content,
						})
					if err == io.EOF {
						break
					}
					Expect(err).ShouldNot(HaveOccurred())
					// Wait until we got response
					_, err = stream.Receive()
					if err == io.EOF {
						break
					}
				}
				Expect(stream.CloseRequest()).ShouldNot(HaveOccurred())
			})
			It("gotPages should be equal to pages", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(gotPages).To(Equal(pages))
			})
		})
	})
})
