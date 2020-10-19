package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"server/services"
)

func main() {
	gwmux := runtime.NewServeMux()
	gRpcEndPoint := "localhost:8081"
	ctx := context.Background()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx, gwmux, gRpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(ctx, gwmux, gRpcEndPoint, opt)

	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	fmt.Println("http-server is running...")
	err = httpServer.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
