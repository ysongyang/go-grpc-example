package services

import (
	"context"
	"fmt"
)

type ProdService struct {
}

func (*ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.ProArea {
	case ProAreas_A:
		stock = 100
		break
	case ProAreas_B:
		stock = 201
		break
	case ProAreas_C:
		stock = 31
		break
	}
	fmt.Println(req)
	return &ProdResponse{ProdStock: stock}, nil
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

func (*ProdService) GetProdInfo(context.Context, *ProdRequest) (*ProdModel, error) {
	ret := &ProdModel{ProdId: 101, ProdName: "测试商品", ProdPrice: 20.5}
	return ret, nil
}
