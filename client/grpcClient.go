package main

import (
	. "client/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
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

	//prodClient := NewProdServiceClient(conn)
	//orderClient := NewOrderServiceClient(conn)
	userClient := NewUserServiceClient(conn)
	ctx := context.Background()

	/*resp, err := prodClient.GetProdStock(ctx, &ProdRequest{
		ProdId:  1,
		ProArea: ProAreas_B,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ProdStock)*/

	/*resp, err := prodClient.GetProdStocks(ctx, &QuerySize{Size: 10})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ProdList)*/

	/*resp, err := prodClient.GetProdInfo(ctx, &ProdRequest{
		ProdId:  1,
	})*/

	/*t := timestamp.Timestamp{Seconds: time.Now().Unix()}

	resp, err := orderClient.NewOrder(ctx, &OrderModel{
		OrderId:    10001,
		OrderNo:    "20201019091712313",
		OrderMoney: 99.99,
		OrderTime:  &t,
	})

	if err != nil {
		log.Fatal(err)
	}*/

	/*
		//服务端流
			var i int32
			req := UserScoreRequest{}
			req.Users = make([]*UserInfo, 0)
			for i = 1; i <= 100; i++ {
				req.Users = append(req.Users, &UserInfo{UserId: i})
			}
			resp, err := userClient.GetUserScore(ctx, &req)
			fmt.Println(resp.Users)*/

	/*respStream, err := userClient.GetUserScoreByServerStream(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := respStream.Recv()
		if err == io.EOF { //读取结束后
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Users)
	}*/

	/*
		//客户端流
		respStream, err := userClient.GetUserScoreByClientStream(ctx)
		if err != nil {
			log.Fatal(err)
		}
		var i int32

		for j := 1; j <= 3; j++ {
			req := UserScoreRequest{}
			req.Users = make([]*UserInfo, 0)
			for i = 1; i <= 5; i++ {
				req.Users = append(req.Users, &UserInfo{UserId: i})
			}

			err := respStream.Send(&req)
			if err != nil {
				log.Fatal(err)
			}
		}
		resp, _ := respStream.CloseAndRecv()
		fmt.Println(resp.Users)*/

	//双向流
	respStream, err := userClient.GetUserScoreByTws(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var i int32
	var uid int32 = 1
	for j := 1; j <= 3; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i = 1; i <= 5; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: uid})
			uid++
		}

		err := respStream.Send(&req)
		if err != nil {
			log.Fatal(err)
		}
		res, err := respStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(res.Users)
	}

}
