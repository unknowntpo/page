package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/unknowntpo/page/domain"
	mock "github.com/unknowntpo/page/domain/mock"
	"github.com/unknowntpo/page/pkg/debug"
	"github.com/unknowntpo/page/pkg/errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRedisRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RedisRepo")
}

var _ = Describe("PageRepo", func() {
	var (
		repo   domain.PageRepo
		client *redis.Client
	)

	BeforeEach(func() {
		client = PrepareTestDatabase()
		repo = NewPageRepo(client)
	})

	Context("NewList is called", func() {
		var (
			listKey domain.ListKey
		)
		const (
			userID int64 = 33
		)
		BeforeEach(func() {
			listKey = domain.ListKey("testList")
			Expect(repo.NewList(context.Background(), userID, listKey)).ShouldNot(HaveOccurred())
		})
		It("should initialize data structures for list", func() {
			assertFn := func() {
				// get content of `<listKey>-meta:<userID>`, make sure head, tail, nextCandidate is there
				res, err := client.HGetAll(
					context.Background(),
					string(domain.GenerateListMetaKeyByUserID(listKey, userID)),
				).Result()
				Expect(err).ShouldNot(HaveOccurred())
				fmt.Println("got res", res)

				head, ok := res["head"]
				Expect(ok).To(BeTrue())
				Expect(head).To(Equal(""))

				tail, ok := res["tail"]
				Expect(ok).To(BeTrue())
				Expect(tail).To(Equal(""))

				nextCandidate, ok := res["nextCandidate"]
				Expect(ok).To(BeTrue())
				Expect(nextCandidate).NotTo(Equal(""))
			}
			assertFn()
		})
	})

	Context("SetPage is called", func() {
		Context("NewList hasn't been called", func() {
			var (
				listKey domain.ListKey
				p       domain.Page
			)
			const (
				userID int64 = 33
			)
			BeforeEach(func() {
				listKey = domain.ListKey("testList")
				// Expect(repo.NewList(context.Background(), userID, listKey)).ShouldNot(HaveOccurred())
			})
			It("should return ResourceNotFound error", func() {
				_, err := repo.SetPage(context.Background(), userID, listKey, p)
				Expect(errors.KindIs(err, errors.ResourceNotFound)).To(BeTrue())
			})
		})
		Context("NewList has been called before", func() {
			var (
				listKey domain.ListKey
			)
			const (
				userID int64 = 33
			)
			BeforeEach(func() {
				listKey = domain.ListKey("testList")
				Expect(repo.NewList(context.Background(), userID, listKey)).ShouldNot(HaveOccurred())
			})
			When("SetPage is called", func() {
				var (
					pages []domain.Page
				)
				BeforeEach(func() {
					pages = []domain.Page{
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
					for i := range pages {
						pageKey, err := repo.SetPage(context.Background(), userID, listKey, pages[i])
						Expect(err).ShouldNot(HaveOccurred())
						// Set the pageKey back
						pages[i].Key = pageKey
					}
				})

				It("related data structures should be set", func() {
					assertListMeta := func() {
						// get content of `<listKey>-meta:<userID>`, make sure head, tail, nextCandidate is there
						res, err := client.HGetAll(
							context.Background(),
							string(domain.GenerateListMetaKeyByUserID(listKey, userID)),
						).Result()
						Expect(err).ShouldNot(HaveOccurred())

						// head should be set to first element of pages
						head, ok := res["head"]
						Expect(ok).To(BeTrue())
						Expect(head).To(Equal(string(pages[0].Key)))

						// tail should be set to last element of pages
						tail, ok := res["tail"]
						Expect(ok).To(BeTrue())
						Expect(tail).To(Equal(string(pages[len(pages)-1].Key)))

						nextCandidate, ok := res["nextCandidate"]
						Expect(ok).To(BeTrue())
						Expect(nextCandidate).NotTo(Equal(""))
					}

					assertPageList := func() {
						// get content of `<listKey>-meta:<userID>`, make sure head, tail, nextCandidate is there
						rangeOpts := &redis.ZRangeBy{
							Min: "0",
							Max: "+inf",
						}
						res, err := client.ZRangeByScore(
							context.Background(),
							string(domain.GenerateListKeyByUserID(listKey, userID)),
							rangeOpts,
						).Result()
						Expect(err).ShouldNot(HaveOccurred())

						fmt.Println("want pageList: ", debug.Debug(pages))

						fmt.Println("pageList: ", debug.Debug(res))
						for i := 0; i < len(pages); i++ {
							Expect(res[i]).To(Equal(string(pages[i].Key)))
						}
					}
					assertListMeta()
					assertPageList()
				})
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
