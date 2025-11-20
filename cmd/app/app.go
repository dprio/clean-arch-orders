package app

import (
	"github.com/dprio/clean-arch-orders/internal/infrastructure/config"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/db"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/web/handlers"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/web/webserver"
	"github.com/dprio/clean-arch-orders/internal/usecase"
	"github.com/dprio/clean-arch-orders/pkg/events"
)

type App struct {
	useCases  *usecase.UseCases
	webServer *webserver.WebServer
}

func New() *App {
	conf := config.New()

	dataBase := db.New(conf.DB)

	eventDispatcher := events.NewEventDispatcher()

	useCases := usecase.New(dataBase, eventDispatcher)

	handlers := handlers.New(*useCases)

	webServer := createWebServer(conf, handlers)

	return &App{
		useCases:  useCases,
		webServer: webServer,
	}

}

func createWebServer(conf *config.Config, handls *handlers.Handlers) *webserver.WebServer {
	webServer := webserver.New(conf.Web)

	webServer.AddHandler("POST", "/order", handls.CreateOrderHandler.Create)

	return webServer
}

func (app *App) Start() error {
	return app.webServer.Start()
}
