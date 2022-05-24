package main

import (
	"learning-rabbitmq/configs"
	"learning-rabbitmq/constants"
	"learning-rabbitmq/helpers"

	"github.com/streadway/amqp"
)

func main() {
	// open connection
	amqpConn, err := amqp.Dial(configs.Amqp.GetConnectionURL())
	helpers.FailOnError(err, "failed connect AMQP")
	defer amqpConn.Close()

	// open channel
	ch, err := amqpConn.Channel()
	helpers.FailOnError(err, "failed connect channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		constants.AMQP_FANOUT_EXCHANGE_NAME, // name
		amqp.ExchangeFanout,                 // type
		false,                               // durable
		false,                               // auto-deleted
		false,                               // internal
		false,                               // no-wait
		nil,                                 // arguments
	)
	helpers.FailOnError(err, "failed to declare a queue")

	err = ch.Publish(
		constants.AMQP_FANOUT_EXCHANGE_NAME, // exchange
		"",                                  // routing key
		false,                               // mandatory
		false,                               // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("from publisher fanout exchange"),
		})
	helpers.FailOnError(err, "Failed to publish a message")
}
