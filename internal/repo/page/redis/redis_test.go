package redis

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
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

	ginkgo.When("SetAndGet to a List", func() {
		var (
			err     error
			pages   []domain.Page
			listKey domain.ListKey
		)
		ginkgo.BeforeEach(func() {
			listKey = domain.ListKey("testList")
			pages = []domain.Page{
				{
					// Key      PageKey
					// Articles []Article
					// NextPage PageKey
					Key:      domain.GeneratePageKey(),
					Articles: domain.GenerateDummyArticles(3),
				},
				{
					// Key      PageKey
					// Articles []Article
					// NextPage PageKey
					Key:      domain.GeneratePageKey(),
					Articles: domain.GenerateDummyArticles(3),
				},
				{
					// Key      PageKey
					// Articles []Article
					// NextPage PageKey
					Key:      domain.GeneratePageKey(),
					Articles: domain.GenerateDummyArticles(3),
				},
			}
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		})
		ginkgo.When("Call SetPage for every page", func() {
			var (
				gotPages []domain.Page
				gotPage  domain.Page
				gotHead  domain.PageKey
				err      error
			)
			ginkgo.BeforeEach(func() {
				gotHead, err = repo.GetHead(context.Background(), listKey)
				gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

				curPageKey := gotHead
				for i := 0; i < len(pages); i++ {
					gotPage, err = repo.GetPage(context.Background(), curPageKey)
					gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
					gotPages = append(gotPages, gotPage)
				}
			})
			ginkgo.It("every page should be set in FIFO order", func() {
				gomega.Expect(gotPage).To(gomega.Equal(pages))
			})
		})
	})
})

var _ = ginkgo.Describe("PingPong", func() {
	var client *redis.Client

	ginkgo.BeforeEach(func() {
		client = PrepareTestDatabase()
	})

	ginkgo.When("PING", func() {
		var (
			err    error
			stsCmd *redis.StatusCmd
		)
		ginkgo.BeforeEach(func() {
			stsCmd = client.Ping(context.Background())
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
		})
		ginkgo.It("should", func() {
			gomega.Expect(stsCmd.Result()).To(gomega.Equal("PONG"))
		})
	})
})
