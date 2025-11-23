package service

import (
	"context"

	"github.com/dprio/clean-arch-orders/internal/infrastructure/grpc/pb"
	"github.com/dprio/clean-arch-orders/internal/usecase/createorder"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase createorder.UseCase
}

func NewOrderService(createOrderUseCase createorder.UseCase) pb.OrderServiceServer {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := createorder.Input{
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(ctx, dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
