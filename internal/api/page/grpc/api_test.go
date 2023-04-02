package page_test

import (
	"context"
	"testing"

	page "github.com/unknowntpo/page/internal/api/page/grpc"
	"github.com/unknowntpo/page/internal/domain"
	"github.com/unknowntpo/page/internal/domain/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRedisRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisRepo")
}

var _ = Describe("PageAPI", func() {
	var (
		api             domain.PageAPI
		mockPageUsecase *mock.MockPageUsecase
		gotPageKey      domain.PageKey
	)

	const (
		testListKey      domain.ListKey = "testListKey"
		dummyHeadPageKey domain.PageKey = "dummy"
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockPageUsecase = mock.NewMockPageUsecase(ctrl)
		api = page.NewPageAPI(mockPageUsecase)
	})

	When("GetHead is called", func() {
		var (
			err error
		)
		BeforeEach(func() {
			mockPageUsecase.
				EXPECT().
				GetHead(context.Background(), testListKey).
				Return(dummyHeadPageKey, nil).Times(1)
		})
		BeforeEach(func() {
			gotPageKey, err = api.GetHead(context.Background(), testListKey)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return expected value", func() {
			Expect(gotPageKey).To(Equal(dummyHeadPageKey))
		})
	})
})
