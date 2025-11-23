package usecase

import (
	"github.com/dprio/clean-arch-orders/internal/domain/eventtype"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/db"
	"github.com/dprio/clean-arch-orders/internal/usecase/createorder"
	"github.com/dprio/clean-arch-orders/internal/usecase/getorders"
	"github.com/dprio/clean-arch-orders/pkg/events"
)

type UseCases struct {
	CreateOrderUseCase createorder.UseCase
	GetOrdersUseCase   getorders.UseCase
}

func New(dbs *db.DBs, eventDispatcher events.EventDispatcherInterface) *UseCases {
	return &UseCases{
		CreateOrderUseCase: buildCreateOrderUseCase(dbs, eventDispatcher),
		GetOrdersUseCase:   getorders.New(dbs.OrderRepository),
	}
}

func buildCreateOrderUseCase(dbs *db.DBs, eventDispatcher createorder.EventDispatcher) createorder.UseCase {
	eventCreator := events.NewEventCreator(eventtype.OrderCreated)
	return createorder.New(dbs.OrderRepository, eventDispatcher, eventCreator)
}
