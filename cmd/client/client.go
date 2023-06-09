package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/bufbuild/connect-go"
	pb "github.com/unknowntpo/page/gen/proto/page"
	"github.com/unknowntpo/page/gen/proto/page/pageconnect"
	"golang.org/x/net/http2"
)

const (
	address = "http://127.0.0.1:8080"
)

func newInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				// If you're also using this client for non-h2c traffic, you may want
				// to delegate to tls.Dial if the network isn't TCP or the addr isn't
				// in an allowlist.
				return net.Dial(network, addr)
			},
			// Don't forget timeouts!
		},
	}
}

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

func newList(client pageconnect.PageServiceClient, listKey string, userID int64) error {
	newListReq := connect.NewRequest(&pb.NewListRequest{
		ListKey: listKey,
		UserID:  userID,
	})
	newListResp, err := client.NewList(context.Background(), newListReq)
	if err != nil {
		return err
	}
	fmt.Println("newListResp", newListResp)
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
		setPageResp, err := setPageStream.Receive()
		if err != nil {
			switch {
			case err == io.EOF:
				goto END
			default:
				return err
			}
		}
		fmt.Printf("Page set with key: %s\n", setPageResp.PageKey)
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
	fmt.Printf("Head of list is page key: %s\n", head)
	return head, nil
}

func getPage(client pageconnect.PageServiceClient, headKey string, userID int64, numOfPage int) ([]string, error) {
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
