package rabbit

import (
	"context"

	"github.com/streadway/amqp"
)

type Setting struct {
	Host       string
	Tag        string
	Que        string
	RoutingKey map[string]string
}

type amqpStruct struct {
	setting Setting

	conn        *amqp.Connection
	ch          *amqp.Channel
	notifyClose chan *amqp.Error
	msg         <-chan amqp.Delivery
}

type RabbitI interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	IsClose() chan *amqp.Error

	Publish( key string, payload []byte) error
}

func NewRabbit(setting Setting) RabbitI {
	return &amqpStruct{
		setting: setting,
	}
}
