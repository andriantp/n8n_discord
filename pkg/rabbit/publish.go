package rabbit

import (
	"time"

	"github.com/streadway/amqp"
)

func (r *amqpStruct) Publish(key string, payload []byte) error {
	routingkey := r.setting.RoutingKey[key]
	err := r.ch.Publish(
		"amq.topic", // exchange
		routingkey,  // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			Timestamp:    time.Now().UTC(),
			ContentType:  "application/json",
			Body:         payload,
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		return err
	}

	return nil
}
