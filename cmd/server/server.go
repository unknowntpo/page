package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pageAPI "github.com/unknowntpo/page/internal/api/page/grpc"
	pageUcase "github.com/unknowntpo/page/internal/usecase/page"

	pb "github.com/unknowntpo/page/internal/api/page/grpc/page"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:4000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	pageUsecase := pageUcase.NewPageUsecase()

	pageServer := pageAPI.NewPageServer(pageUsecase)
	pb.RegisterPageServiceServer(grpcServer, pageServer)
	grpcServer.Serve(lis)
}
