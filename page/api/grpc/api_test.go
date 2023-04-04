package page_test

import (
	"context"
	"testing"

	"github.com/unknowntpo/page/internal/domain"
	"github.com/unknowntpo/page/internal/domain/mock"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPageAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PageAPI")
}

var _ = Describe("PageAPI", func() {
	var (
		api             *page.PageServer
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
		api = page.NewPageServer(mockPageUsecase)
	})

	When("api.GetHead is called", func() {
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
		It("should return expected PageKey", func() {
			Expect(gotPageKey).To(Equal(dummyHeadPageKey))
		})
	})
})
