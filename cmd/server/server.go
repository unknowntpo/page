package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pageAPI "github.com/unknowntpo/page/page/api/grpc"
	pageUcase "github.com/unknowntpo/page/page/usecase"

	"github.com/unknowntpo/page/gen/page/api/grpc/page/pageconnect"
)

// func main() {
// 	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:4000"))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	var opts []grpc.ServerOption

// 	grpcServer := grpc.NewServer(opts...)
// 	reflection.Register(grpcServer)

// 	pageUsecase := pageUcase.NewPageUsecase()

// 	pageServer := pageAPI.NewPageServer(pageUsecase)
// 	pb.RegisterPageServiceServer(grpcServer, pageServer)
// 	grpcServer.Serve(lis)
// }

func main() {
	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:4000"))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// var opts []grpc.ServerOption

	mux := http.NewServeMux()

	pageUsecase := pageUcase.NewPageUsecase()
	pageServer := pageAPI.NewPageServer(pageUsecase)

	path, handler := pageconnect.NewPageServiceHandler(pageServer)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
