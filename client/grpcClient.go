package main

import (
	. "client/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	/*creds, err := credentials.NewClientTLSFromFile("key/server.crt", "my.com")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	*/

	/*cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})*/

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	prodClient := NewProdServiceClient(conn)

	ctx := context.Background()

	/*prodResp, err := prodClient.GetProdStock(context.Background(), &ProdRequest{
		ProdId: 1,
	})*/

	resp, err := prodClient.GetProdStocks(ctx, &QuerySize{Size: 10})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ProdList)

}
