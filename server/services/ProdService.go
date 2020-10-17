package services

import (
	"context"
)

type ProdService struct {
}

func (*ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}

// @title 实现GetProdStocks方法
func (*ProdService) GetProdStocks(ctx context.Context, in *QuerySize) (*ProdRespList, error) {

	prodRes := []*ProdResponse{
		&ProdResponse{ProdStock: 20},
		&ProdResponse{ProdStock: 21},
		&ProdResponse{ProdStock: 22},
		&ProdResponse{ProdStock: 23},
		&ProdResponse{ProdStock: 24},
	}
	return &ProdRespList{ProdList: prodRes}, nil
}
