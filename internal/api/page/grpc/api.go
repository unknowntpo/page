package page

import (
	"context"
	"io"
	"log"

	"github.com/pkg/errors"
	"github.com/unknowntpo/page/internal/domain"

	pb "github.com/unknowntpo/page/internal/api/page/grpc/page"
)

func NewPageServer(uCase domain.PageUsecase) *pageServer {
	return &pageServer{useCase: uCase}
}

type pageServer struct {
	pb.UnimplementedPageServiceServer
	pages   map[string]*pb.Page
	useCase domain.PageUsecase
}

func (s *pageServer) GetHead(req *pb.GetHeadRequest, stream pb.PageService_GetHeadServer) error {
	// TODO: Implement the logic to get the head page key
	// For example:
	var listKey domain.ListKey
	pageKey, err := s.useCase.GetHead(context.Background(), listKey)
	if err != nil {
		return errors.Wrap(err, "failed on api.useCase.GetHead")
	}

	// Send the page key through the stream
	if err := stream.Send(&pb.PageKey{Key: string(pageKey)}); err != nil {
		return err
	}

	return nil
}

func (s *pageServer) GetPage(stream pb.PageService_GetPageServer) error {
	// Receive the page keys from the stream and send the corresponding pages
	for {
		pageKey, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Println("got pageKey: ", pageKey)

		// TODO: Implement the logic to get the page
		// For example:
		page := &pb.Page{
			Title:   "Page Title",
			Content: "Page Content",
		}

		// Send the page through the stream
		if err := stream.Send(page); err != nil {
			return err
		}
	}
}

func (s *pageServer) SetPage(stream pb.PageService_SetPageServer) error {
	// Receive the pages from the stream and set them
	for {
		pg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// TODO: Implement the logic to set the page
		// For example:
		pageKey := "ABC"
		log.Printf("Setting page %s: %s", pageKey, pg.Content)

		// Send the page key through the stream
		if err := stream.SendAndClose(&pb.PageKey{Key: pageKey}); err != nil {
			return err
		}
	}
}
