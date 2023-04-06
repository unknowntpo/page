package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	pageAPI "github.com/unknowntpo/page/page/api/grpc"
	pageUcase "github.com/unknowntpo/page/page/usecase"

	"github.com/unknowntpo/page/gen/proto/page/pageconnect"
)

const addr = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"page.PageService",
		// protoc-gen-connect-go generates package-level constants
		// for these fully-qualified protobuf service names, so you'd more likely
		// reference userv1.UserServiceName and groupv1.GroupServiceName.
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	pageUsecase := pageUcase.NewPageUsecase()
	pageServer := pageAPI.NewPageServer(pageUsecase)

	path, handler := pageconnect.NewPageServiceHandler(pageServer)
	mux.Handle(path, handler)

	log.Printf("Starting server at %s\n", addr)
	http.ListenAndServe(
		addr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
