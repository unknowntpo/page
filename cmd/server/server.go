package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/unknowntpo/playground-2022/go/grpc/page/page"
)

type pageServer struct {
	pb.UnimplementedPageServiceServer
	pages map[string]*pb.Page
}

func (s *pageServer) GetHead(req *pb.GetHeadRequest, stream pb.PageService_GetHeadServer) error {
	// TODO: Implement the logic to get the head page key
	// For example:
	pageKey := "123"

	// Send the page key through the stream
	if err := stream.Send(&pb.PageKey{Key: pageKey}); err != nil {
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

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:4000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	pb.RegisterPageServiceServer(grpcServer, &pageServer{})
	grpcServer.Serve(lis)
}
