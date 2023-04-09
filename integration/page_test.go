package integration

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/unknowntpo/page/domain"

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
		client := pageRepo.PrepareTestDatabase()
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

	When("test List Lifecycle", func() {
		var (
			pages = []string{"page1", "page2", "page3"}
		)

		const (
			listKey string = "my_list"
			userID  int64  = 123
		)

		It("should not failed", func() {
			Expect(newList(client, listKey, userID)).ShouldNot(HaveOccurred())
			// can not create same list twice
			Expect(newList(client, listKey, userID).Error()).To(ContainSubstring(domain.ErrListAlreadyExists.Error()))
			// set pages
			Expect(setPage(client, listKey, userID, pages)).ShouldNot(HaveOccurred())
			_, err := getHead(client, "non-exist-key", userID)
			Expect(err.Error()).To(ContainSubstring(domain.ErrListNotFound.Error()))

			head, err := getHead(client, listKey, userID)
			Expect(err).ShouldNot(HaveOccurred())

			gotPages, err := getPages(client, head, userID, len(pages))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(gotPages).To(Equal(pages))
		})
	})
})

/*
func main() {
	client := pageconnect.NewPageServiceClient(
		newInsecureClient(),
		address,
		connect.WithGRPC(),
	)

	// Create a new list
	listKey := "my_list"
	userID := int64(123)

	if err := newList(client, listKey, userID); err != nil {
		panic(fmt.Errorf("failed on newList: %v", err))
	}

	pages := []string{"page1", "page2", "page3"}
	if err := setPage(client, listKey, userID, pages); err != nil {
		panic(fmt.Errorf("failed on setPage: %v", err))
	}

	head, err := getHead(client, listKey, userID)
	if err != nil {
		panic(fmt.Errorf("failed on getHead: %v", err))
	}

	fmt.Println("got head", head)

	ps, err := getPage(client, head, userID, len(pages))
	if err != nil {
		panic(fmt.Errorf("failed on getPage: %v", err))
	}
	fmt.Println(ps)
}
*/

func newList(client pageconnect.PageServiceClient, listKey string, userID int64) error {
	req := connect.NewRequest(&pb.NewListRequest{
		ListKey: listKey,
		UserID:  userID,
	})
	_, err := client.NewList(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func setPage(client pageconnect.PageServiceClient, listKey string, userID int64, pageContents []string) error {
	// Set a page with dummy content

	setPageStream := client.SetPage(context.Background())

	for _, ps := range pageContents {
		setPageReq := &pb.SetPageRequest{
			UserID:      userID,
			ListKey:     listKey,
			PageContent: ps,
		}
		if err := setPageStream.Send(setPageReq); err != nil {
			if err != nil {
				switch {
				case err == io.EOF:
					goto END
				default:
					return err
				}
			}
		}
		_, err := setPageStream.Receive()
		if err != nil {
			switch {
			case err == io.EOF:
				goto END
			default:
				return err
			}
		}
	}

END:
	if err := setPageStream.CloseRequest(); err != nil {
		return err
	}
	if err := setPageStream.CloseResponse(); err != nil {
		return nil
	}

	return nil
}

func getHead(client pageconnect.PageServiceClient, listKey string, userID int64) (string, error) {
	// Get the head of the list
	getHeadReq := connect.NewRequest(&pb.GetHeadRequest{
		ListKey: listKey,
		UserID:  userID,
	})
	getHeadResp, err := client.GetHead(context.Background(), getHeadReq)
	if err != nil {
		return "", err
	}
	head := getHeadResp.Msg.PageKey
	return head, nil
}

func getPages(client pageconnect.PageServiceClient, headKey string, userID int64, numOfPage int) ([]string, error) {
	getPageStream := client.GetPage(context.Background())

	var gotPages []string

	cur := headKey
	for i := 0; i < numOfPage; i++ {
		req := &pb.GetPageRequest{
			PageKey: cur,
		}
		err := getPageStream.Send(req)
		if err != nil {
			switch {
			case err == io.EOF:
				goto END
			default:
				return nil, err
			}
		}
		res, err := getPageStream.Receive()
		if err != nil {
			switch {
			case err == io.EOF:
				goto END
			default:
				return nil, err
			}
		}
		gotPages = append(gotPages, res.PageContent)
		cur = res.Next
	}

END:
	if err := getPageStream.CloseRequest(); err != nil {
		return nil, err
	}
	if err := getPageStream.CloseResponse(); err != nil {
		return nil, err
	}

	return gotPages, nil
}
