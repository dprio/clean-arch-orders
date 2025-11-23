package resolvers

import "github.com/dprio/clean-arch-orders/internal/usecase/createorder"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	createOrderUseCase createorder.UseCase
}
