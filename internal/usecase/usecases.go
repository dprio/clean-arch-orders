package usecase

import (
	"github.com/dprio/clean-arch-orders/internal/infrastructure/db"
	"github.com/dprio/clean-arch-orders/internal/usecase/createorder"
	"github.com/dprio/clean-arch-orders/pkg/events"
)

type UseCases struct {
	CreateOrderUseCase createorder.UseCase
}

func New(dbs *db.DBs, eventDispatcher events.EventDispatcherInterface) *UseCases {
	return &UseCases{
		CreateOrderUseCase: buildCreateOrderUseCase(dbs, eventDispatcher),
	}
}

func buildCreateOrderUseCase(dbs *db.DBs, eventDispatcher createorder.EventDispatcher) createorder.UseCase {
	eventCreator := events.NewEventCreator("order_created")
	return createorder.New(dbs.OrderRepository, eventDispatcher, eventCreator)
}
