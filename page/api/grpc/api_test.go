package page

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
			validListKey        = "validListKey"
			alreadyExistListKey = "alreadyExistListKey"
			validUserID         = 33
			invalidUserID       = 0
			invalidListKey      = ""
		)
		JustBeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				NewList(gomock.Any(), userID, gomock.Any()).
				DoAndReturn(func(ctx context.Context, userID int64, listKey domain.ListKey) error {
					switch string(listKey) {
					case validListKey:
						// normal
						return nil
					case alreadyExistListKey:
						return domain.ErrListAlreadyExists
					default:
						return domain.ErrInternal
					}
				}).AnyTimes()
			res, err = client.NewList(context.Background(), connect.NewRequest(&pb.NewListRequest{
				ListKey: listKey,
				UserID:  userID,
			}))
		})
		Context("normal", func() {
			BeforeEach(func() {
				userID = validUserID
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
				listKey = invalidListKey
			})
			It("should return invalid argument", func() {
				Expect(err.Error()).To(ContainSubstring(domain.ErrInvalidListKey.Error()))
			})
		})
		When("userID not valid", func() {
			BeforeEach(func() {
				userID = invalidUserID
				listKey = validListKey
			})
			It("should return invalid userID error", func() {
				Expect(err.Error()).To(ContainSubstring(domain.ErrInvalidUserID.Error()))
			})
		})
		When("list already exist", func() {
			BeforeEach(func() {
				userID = validUserID
				listKey = alreadyExistListKey
			})
			It("should return invalid userID error", func() {
				Expect(err.Error()).To(ContainSubstring(domain.ErrListAlreadyExists.Error()))
			})
		})
	})

	Describe("GetHead", func() {
		var (
			res                 *connect.Response[pb.GetHeadResponse]
			userID              int64
			listKey             string
			err                 error
			expectedHeadPageKey = domain.GeneratePageKey(time.Now())
		)
		const (
			existListKey domain.ListKey = "existListKey"
		)

		BeforeEach(func() {
			userID = 33
			mockPageUsecase.
				EXPECT().
				GetHead(gomock.Any(), userID, gomock.Any()).
				DoAndReturn(func(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
					if listKey != existListKey {
						return "", domain.ErrListNotFound
					}
					return expectedHeadPageKey, nil
				}).AnyTimes()
		})
		JustBeforeEach(func() {
			res, err = client.GetHead(context.Background(), connect.NewRequest(&pb.GetHeadRequest{
				ListKey: listKey,
				UserID:  userID,
			}))
		})
		Context("normal", func() {
			BeforeEach(func() {
				listKey = string(existListKey)
			})
			It("should return expected PageKey", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.Msg.PageKey).To(Equal(string(expectedHeadPageKey)))
			})
		})
		Context("list not found", func() {
			BeforeEach(func() {
				listKey = "notExistList"
			})
			It("should return domain.ErrListNotFound", func() {
				Expect(err.Error()).To(ContainSubstring(domain.ErrListNotFound.Error()))
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

	Describe("GetPage", func() {
		var (
			res     *pb.GetPageResponse
			userID  int64
			listKey string
			pageKey string
			err     error
			stream  *connect.BidiStreamForClient[pb.GetPageRequest, pb.GetPageResponse]
			//			errors              []error
			//			s
		)
		const (
			existListKey domain.ListKey = "existListKey"
			existPagekey domain.PageKey = "existPageKey"
			nextPageKey  domain.PageKey = "nextPageKey"
			pageContent                 = "dummy"
			round                       = 3
		)

		BeforeEach(func() {
			userID = 33
			mockPageUsecase.
				EXPECT().
				GetPage(gomock.Any(), userID, gomock.Any(), gomock.Any()).
				DoAndReturn(func(ctx context.Context, userID int64, listKey domain.ListKey, pageKey domain.PageKey) (domain.Page, error) {
					if listKey != existListKey {
						return domain.Page{}, domain.ErrListNotFound
					}
					if pageKey != existPagekey {
						return domain.Page{}, domain.ErrPageNotFound
					}
					return domain.Page{Key: existPagekey, Content: pageContent, Next: nextPageKey}, nil
				}).AnyTimes()
		})
		JustBeforeEach(func() {
			stream = client.GetPage(context.Background())

			for i := 0; i < round; i++ {
				e := stream.Send(
					&pb.GetPageRequest{
						UserID:  userID,
						ListKey: listKey,
						PageKey: pageKey,
					})
				if e != nil {
					switch e {
					case io.EOF:
						goto END
					default:
						err = e
						return
					}
				}

				// Wait until we got response
				res, e = stream.Receive()
				if e != nil {
					switch e {
					case io.EOF:
						goto END
					default:
						err = e
						goto END
					}
				}
			}

		END:
			Expect(stream.CloseRequest()).ShouldNot(HaveOccurred())
			Expect(stream.CloseResponse()).ShouldNot(HaveOccurred())
		})
		Context("normal", func() {
			BeforeEach(func() {
				pageKey = string(existPagekey)
				listKey = string(existListKey)
			})
			It("should return expected PageKey", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.Key).To(Equal(string(existPagekey)))
				Expect(res.PageContent).To(Equal(pageContent))
				Expect(res.Next).To(Equal(string(nextPageKey)))
			})
		})
		Context("list not found", func() {
			BeforeEach(func() {
				listKey = "notExistList"
				pageKey = string(existPagekey)
			})
			It("should return domain.ErrListNotFound", func() {
				Expect(err.Error()).To(ContainSubstring(domain.ErrListNotFound.Error()))
			})
		})
		Context("existed list, but page not found", func() {
			BeforeEach(func() {
				listKey = string(existListKey)
				pageKey = "notExistPage"
			})
			It("should return domain.ErrPageNotFound", func() {
				Expect(err.Error()).To(ContainSubstring(domain.ErrPageNotFound.Error()))
			})
		})
	})
})
