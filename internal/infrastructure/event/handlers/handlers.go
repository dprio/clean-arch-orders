package handlers

import (
	"github.com/dprio/clean-arch-orders/pkg/events"
	"github.com/streadway/amqp"
)

func CreateAndRegisterEventHandlers(dispatcher events.EventDispatcherInterface) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	CreateAndRegisterOrderCreatedHandler(ch, dispatcher)
}
