package createorder

import (
	"context"

	"github.com/dprio/clean-arch-orders/internal/domain"
)

type (
	OrderRepository interface {
		Save(ctx context.Context, order *domain.Order) (domain.Order, error)
	}

	EventDispatcher interface {
		Dispatch(ctx context.Context, event any) error
	}

	UseCase interface {
		Execute(ctx context.Context, input Input) (Output, error)
	}
)
