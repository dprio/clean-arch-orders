package resolvers

import "github.com/dprio/clean-arch-orders/internal/usecase"

type Resolvers struct {
	OrderResolver *Resolver
}

func NewGraphQLResolvers(useCases *usecase.UseCases) *Resolvers {
	return &Resolvers{
		OrderResolver: &Resolver{createOrderUseCase: useCases.CreateOrderUseCase},
	}

}
