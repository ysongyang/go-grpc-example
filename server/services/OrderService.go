package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (*OrderService) NewOrder(ctx context.Context, orderReq *OrderRequest) (*OrderResponse, error) {
	fmt.Println(orderReq.OrderModel)
	resp := &OrderResponse{
		Status:  "success",
		Message: "success",
	}
	return resp, nil
}
