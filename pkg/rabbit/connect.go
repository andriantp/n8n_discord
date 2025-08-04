package rabbit

import (
	"context"
	"time"

	"github.com/streadway/amqp"
)

func (r *amqpStruct) Connect(ctx context.Context) error {
	conn, err := amqp.DialConfig(r.setting.Host,
		amqp.Config{Heartbeat: 30 * time.Second},
	)
	if err != nil {
		return err
	}
	r.conn = conn

	r.ch, err = r.conn.Channel()
	if err != nil {
		return err
	}

	err = r.binding()
	if err != nil {
		return err
	}

	r.notifyClose = r.conn.NotifyClose(make(chan *amqp.Error))
	return nil
}

func (r *amqpStruct) Disconnect(ctx context.Context) error {
	err := r.ch.Close()
	if err != nil {
		return err
	}
	err = r.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *amqpStruct) IsClose() chan *amqp.Error {
	return r.notifyClose
}

func (r *amqpStruct) binding() error {
	_, err := r.ch.QueueDeclare(
		r.setting.Que, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	for _, routingkey := range r.setting.RoutingKey {
		err = r.ch.QueueBind(
			r.setting.Que, //  queue
			routingkey,    // routing key
			"amq.topic",   //  exchange name
			false,         // no-wait
			nil,           // arguments
		)
		if err != nil {
			return err
		}
	}

	return nil
}
