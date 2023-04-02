package page

import (
	"testing"

	"github.com/unknowntpo/page/internal/domain"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestRedisRepo(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "RedisRepo")
}

var _ = ginkgo.Describe("PageRepo", func() {
	var api domain.PageAPI

	ginkgo.BeforeEach(func() {
		api = NewPageAPI()
	})

	ginkgo.When("GetPage is called", func() {

	})
})
