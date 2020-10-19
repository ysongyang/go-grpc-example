package services

import (
	"context"
	"io"
	"log"
	"time"
)

type UserService struct {
}

//普通方法
func (*UserService) GetUserScore(ctx context.Context, req *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range req.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{
		Users: users,
	}, nil
}

//服务端流
func (*UserService) GetUserScoreByServerStream(req *UserScoreRequest,
	stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for i, user := range req.Users {
		user.UserScore = score
		score++
		users = append(users, user)
		//每10条发送
		if (i+1)%9 == 0 && i > 0 {
			err := stream.Send(&UserScoreResponse{
				Users: users,
			})
			if err != nil {
				return err
			}
			//清空切片
			users = (users)[0:0]
		}
		time.Sleep(time.Second * 1)
	}
	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{
			Users: users,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//客户端流
func (*UserService) GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&UserScoreResponse{
				Users: users,
			})
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
	}
	return nil
}

// 双向流
func (*UserService) GetUserScoreByTws(stream UserService_GetUserScoreByTwsServer) error {

	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF { //接收完了
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
		err = stream.Send(&UserScoreResponse{
			Users: users,
		})
		if err != nil {
			log.Println(err)
			return err
		}
		//清空切片
		users = (users)[0:0]
	}

	return nil
}
