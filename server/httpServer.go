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

	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
