package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	. "server/services"
)

func main() {
	/*creds, err := credentials.NewServerTLSFromFile("key/server.crt", "key/server_no_passwd.key")
	if err != nil {
		fmt.Println(err)
	}
	gRpcService := grpc.NewServer(grpc.Creds(creds))*/

	/*cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	gRpcService := grpc.NewServer(grpc.Creds(creds))*/

	gRpcService := grpc.NewServer()

	RegisterProdServiceServer(gRpcService, new(ProdService))   //注册商品服务
	RegisterOrderServiceServer(gRpcService, new(OrderService)) //注册订单服务
	RegisterUserServiceServer(gRpcService, new(UserService))   //注册用户服务
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("grpc-server is running...")
	err = gRpcService.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}

}
