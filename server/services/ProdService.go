package services

import (
	"context"
)

type ProdService struct {
}

func (*ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}
