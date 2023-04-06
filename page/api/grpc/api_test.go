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

	const (
		testListKey      domain.ListKey = "testListKey"
		dummyHeadPageKey domain.PageKey = "dummy"
		userID           int64          = 33
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

	When("NewList is called", func() {
		var (
			err error
			res *connect.Response[pb.NewListResponse]
		)
		BeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				NewList(gomock.Any(), userID, testListKey).
				Return(nil).Times(1)
		})
		BeforeEach(func() {
			res, err = client.NewList(context.Background(), connect.NewRequest(&pb.NewListRequest{
				ListKey: string(testListKey),
				UserID:  userID,
			}))
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return OK", func() {
			Expect(res.Msg.Status).To(Equal("OK"))
		})
	})

	When("GetHead is called", func() {
		var (
			err error
			res *connect.Response[pb.GetHeadResponse]
		)
		BeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				GetHead(gomock.Any(), testListKey).
				Return(dummyHeadPageKey, nil).Times(1)
		})
		BeforeEach(func() {
			res, err = client.GetHead(context.Background(), connect.NewRequest(&pb.GetHeadRequest{
				ListKey: string(testListKey),
				UserID:  33,
			}))
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return expected PageKey", func() {
			Expect(res.Msg.PageKey).To(Equal(string(dummyHeadPageKey)))
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
		)
		const round = 3

		BeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				SetPage(gomock.Any(), userID, testListKey, gomock.Any()).
				DoAndReturn(func(ctx context.Context, userID int64, listKey domain.ListKey, page domain.Page) (domain.PageKey, error) {
					gotPages = append(gotPages, page)
					return dummyHeadPageKey, nil
				}).Times(round)
		})
		BeforeEach(func() {
			stream = client.SetPage(context.Background())
			Expect(err).ShouldNot(HaveOccurred())
		})
		Context("call stream.Send for three times", func() {
			BeforeEach(func() {
				for i := 0; i < round; i++ {
					err := stream.Send(&pb.SetPageRequest{UserID: userID, ListKey: string(testListKey), PageContent: pages[i].Content})
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
