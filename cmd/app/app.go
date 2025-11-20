package app

import (
	"github.com/dprio/clean-arch-orders/internal/infrastructure/config"
	"github.com/dprio/clean-arch-orders/internal/infrastructure/db"
	"github.com/dprio/clean-arch-orders/internal/usecase"
	"github.com/dprio/clean-arch-orders/pkg/events"
)

type App struct {
	useCases *usecase.UseCases
}

func New() *App {
	conf := config.New()

	dataBase := db.New(conf.DB)

	eventDispatcher := events.NewEventDispatcher()

	useCases := usecase.New(dataBase, eventDispatcher)

	return &App{
		useCases: useCases,
	}

}

func (app *App) Start() {}
