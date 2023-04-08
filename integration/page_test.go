package integration

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/unknowntpo/page/domain"
	"github.com/unknowntpo/page/domain/mock"
	"github.com/unknowntpo/page/infra"

	pageAPI "github.com/unknowntpo/page/page/api/grpc"
	pageRepo "github.com/unknowntpo/page/page/repo/redis"
	pageUcase "github.com/unknowntpo/page/page/usecase"

	"github.com/unknowntpo/page/gen/proto/page/pageconnect"

	pb "github.com/unknowntpo/page/gen/proto/page"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPageIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PageIntegration")
}

// Ref: https://github.com/bufbuild/connect-demo/blob/main/main_test.go
var _ = Describe("PageIntegration", Ordered, func() {
	var (
		mux    *http.ServeMux
		server *httptest.Server
		client pageconnect.PageServiceClient
	)

	const (
		testListKey      domain.ListKey = "testListKey"
		dummyHeadPageKey domain.PageKey = "dummy"
	)

	BeforeEach(func() {
		client := infra.NewRedisClient()
		repo := pageRepo.NewPageRepo(client)
		pageUsecase := pageUcase.NewPageUsecase(repo)
		pageServer := pageAPI.NewPageServer(pageUsecase)

		// Start mux
		mux = http.NewServeMux()
		mux.Handle(pageconnect.NewPageServiceHandler(
			pageServer,
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
	})

	When("NewList is called", func() {
		var (
			err    error
			res    *connect.Response[pb.NewListResponse]
			userID int64
		)
		JustBeforeEach(func() {
			res, err = client.NewList(context.Background(), connect.NewRequest(&pb.NewListRequest{
				ListKey: string(testListKey),
				UserID:  userID,
			}))
		})
		Context("normal", func() {
			BeforeEach(func() {
				userID = 33
			})
			It("should return OK", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.Msg.Status).To(Equal("OK"))
			})
		})
		Context("userID not valid", func() {
			BeforeEach(func() {
				userID = 0
			})
			It("should return invalid argument", func() {
				Expect(err.Error()).To(ContainSubstring(pageAPI.ErrInvalidUserID.Error()))
			})
		})

	})

	When("GetHead is called", func() {
		var (
			err    error
			res    *connect.Response[pb.GetHeadResponse]
			userID int64
		)
		BeforeEach(func() {
			userID = 33
			res, err = client.GetHead(context.Background(), connect.NewRequest(&pb.GetHeadRequest{
				ListKey: string(testListKey),
				UserID:  userID,
			}))
			Expect(err).ShouldNot(HaveOccurred())
		})
		Context("normal", func() {
			It("should return expected PageKey", func() {
				Expect(res.Msg.PageKey).To(Equal(string(dummyHeadPageKey)))
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
		)

		const (
			userID int64 = 33
			round  int   = 3
		)

		BeforeEach(func() {
			stream = client.SetPage(context.Background())
		})
		Context("call stream.Send for three times", func() {
			BeforeEach(func() {
				for i := 0; i < round; i++ {
					err := stream.Send(
						&pb.SetPageRequest{
							UserID:      userID,
							ListKey:     string(testListKey),
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
