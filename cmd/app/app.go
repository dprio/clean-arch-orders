package app

import (
	"github.com/dprio/clean-arch-orders/internal/infrastructure/config"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/db"
	eventhandlers "github.com/dprio/clean-arch-orders/internal/infrastructure/event/handlers"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/graph/graphqlserver"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/graph/resolvers"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/grpc/grpcserver"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/grpc/service"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/web/handlers"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/web/webserver"
	"github.com/dprio/clean-arch-orders/internal/usecase"
	"github.com/dprio/clean-arch-orders/pkg/events"
)

type App struct {
	useCases      *usecase.UseCases
	webServer     *webserver.WebServer
	grpcServer    *grpcserver.GRPCServer
	graphQlServer *graphqlserver.GraphQLServer
}

func New() *App {
	conf := config.New()

	dataBase := db.New(conf.DB)

	eventDispatcher := events.NewEventDispatcher()
	eventhandlers.CreateAndRegisterEventHandlers(eventDispatcher)

	useCases := usecase.New(dataBase, eventDispatcher)

	handlers := handlers.New(*useCases)

	webServer := createWebServer(conf, handlers)

	grpcServices := service.NewGRPCServices(useCases)

	grpcServer := grpcserver.New(grpcServices)

	graphQLResolvers := resolvers.NewGraphQLResolvers(useCases)

	graphqlServer := graphqlserver.New("8081", graphQLResolvers)

	return &App{
		useCases:      useCases,
		webServer:     webServer,
		grpcServer:    grpcServer,
		graphQlServer: graphqlServer,
	}

}

func createWebServer(conf *config.Config, handls *handlers.Handlers) *webserver.WebServer {
	webServer := webserver.New(conf.Web)

	webServer.AddHandler("POST", "/order", handls.CreateOrderHandler.Create)

	return webServer
}

func (app *App) Start() error {
	go app.webServer.Start()
	go app.grpcServer.Start()
	return app.graphQlServer.Start()
}
