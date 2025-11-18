package createorder

import (
	"context"

	"github.com/dprio/clean-arch-orders/internal/domain"
	"github.com/google/uuid"
)

type createOrder struct {
	repository      OrderRepository
	eventDispatcher EventDispatcher
}

func New(repository OrderRepository, eventDispatcher EventDispatcher) UseCase {
	return &createOrder{
		repository:      repository,
		eventDispatcher: eventDispatcher,
	}
}

func (co *createOrder) Execute(ctx context.Context, input Input) (Output, error) {
	order, err := domain.NewOrder(uuid.NewString(), input.Price, input.Tax)
	if err != nil {
		return Output{}, err
	}

	if err = order.CaluculateFinalPeice(); err != nil {
		return Output{}, err
	}

	savedOrder, err := co.repository.Save(ctx, order)
	if err != nil {
		return Output{}, err
	}

	out := NewOutput(savedOrder)

	if err = co.eventDispatcher.Dispatch(ctx, out); err != nil {
		return Output{}, err
	}

	return out, nil
}
