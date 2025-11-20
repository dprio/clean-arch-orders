package handlers

import (
	"github.com/dprio/clean-arch-orders/internal/infrastructure/web/handlers/orderhandler"
	"github.com/dprio/clean-arch-orders/internal/usecase"
)

type Handlers struct {
	CreateOrderHandler *orderhandler.OrderHandler
}

func New(useCases usecase.UseCases) *Handlers {
	return &Handlers{
		CreateOrderHandler: orderhandler.New(useCases.CreateOrderUseCase),
	}
}
