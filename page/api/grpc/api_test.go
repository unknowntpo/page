package page

import (
	"context"
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
var _ = Describe("PageAPI", func() {
	var (
		mux             *http.ServeMux
		server          *httptest.Server
		client          pageconnect.PageServiceClient
		mockPageUsecase *mock.MockPageUsecase
	)

	const (
		testListKey      domain.ListKey = "testListKey"
		dummyHeadPageKey domain.PageKey = "dummy"
	)
	BeforeEach(func() {
		// Set up mock usecase
		ctrl := gomock.NewController(GinkgoT())
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
})
