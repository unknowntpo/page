package redis

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/unknowntpo/page/domain"
	mock "github.com/unknowntpo/page/domain/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRedisRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisRepo")
}

var _ = Describe("PageRepo", func() {
	var repo domain.PageRepo

	BeforeEach(func() {
		client := PrepareTestDatabase()
		repo = NewPageRepo(client)
	})

	When("SetAndGet to a List", func() {
		var (
			pages   []domain.Page
			listKey domain.ListKey
		)
		const (
			userID int64 = 33
		)
		BeforeEach(func() {
			listKey = domain.ListKey("testList")
			pages = []domain.Page{
				{
					// Key      PageKey
					// Articles []Article
					// NextPage PageKey
					Key:      domain.GeneratePageKey(),
					Articles: mock.GenerateDummyArticles(3),
				},
				{
					// Key      PageKey
					// Articles []Article
					// NextPage PageKey
					Key:      domain.GeneratePageKey(),
					Articles: mock.GenerateDummyArticles(3),
				},
				{
					// Key      PageKey
					// Articles []Article
					// NextPage PageKey
					Key:      domain.GeneratePageKey(),
					Articles: mock.GenerateDummyArticles(3),
				},
			}
		})
		When("Call SetPage for every page", func() {
			var (
				// gotPages []domain.Page
				gotPage domain.Page
				// gotHead  domain.PageKey
				err error
			)
			BeforeEach(func() {
				// set pages to list
				for _, p := range pages {
					Expect(repo.SetPage(context.Background(), userID, listKey, p)).ShouldNot(HaveOccurred())
				}

				// gotHead, err = repo.GetHead(context.Background(), userID, listKey)
				// Expect(err).ShouldNot(HaveOccurred())

				// curPageKey := gotHead
				// for i := 0; i < len(pages); i++ {
				// 	gotPage, err = repo.GetPage(context.Background(), curPageKey)
				// 	Expect(err).ShouldNot(HaveOccurred())
				// 	gotPages = append(gotPages, gotPage)
				// }
			})
			It("every page should be set in FIFO order", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(gotPage).To(Equal(pages))
			})
		})
	})
})

var _ = Describe("PingPong", func() {
	var client *redis.Client

	BeforeEach(func() {
		client = PrepareTestDatabase()
	})

	When("PING", func() {
		var (
			err    error
			stsCmd *redis.StatusCmd
		)
		BeforeEach(func() {
			stsCmd = client.Ping(context.Background())
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should", func() {
			Expect(stsCmd.Result()).To(Equal("PONG"))
		})
	})
})
