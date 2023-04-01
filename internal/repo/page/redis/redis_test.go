package redis

import (
	"testing"

	"github.com/unknowntpo/page/domain"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestRedisRepo(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "RedisRepo")
}

var _ = ginkgo.Describe("PageRepo", func() {
	var repo domain.PageRepo

	ginkgo.BeforeEach(func() {
		client := PrepareTestDatabase()
		repo = NewPageRepo(client)
	})

	ginkgo.When("lala", func() {
		var (
			err error
		)
		ginkgo.BeforeEach(func() {
			_, err = repo.GetHead(domain.ListKey("fdsaf"))

			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		})
		ginkgo.It("should", func() {
			gomega.Expect(1 + 1).To(gomega.Equal(2))
		})
	})

})
