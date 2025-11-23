package service

import (
	"github.com/dprio/clean-arch-orders/internal/infrastructure/grpc/pb"
	"github.com/dprio/clean-arch-orders/internal/usecase"
)

type GRPCServices struct {
	OrderService pb.OrderServiceServer
}

func NewGRPCServices(useCases *usecase.UseCases) *GRPCServices {
	return &GRPCServices{
		OrderService: NewOrderService(useCases.CreateOrderUseCase),
	}
}
