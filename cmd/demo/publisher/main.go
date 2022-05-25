package main

import (
	"learning-rabbitmq/configs"
	"learning-rabbitmq/helpers"

	"github.com/streadway/amqp"
)

func main() {
	amqpConn, err := amqp.Dial(configs.Amqp.GetConnectionURL())
	helpers.FailOnError(err, "failed connect AMQP")
	defer amqpConn.Close()

	ch, err := amqpConn.Channel()
	helpers.FailOnError(err, "failed connect channel")
	defer ch.Close()

	err = ch.Publish(
		"demo.exchange",     // exchange
		"key.asdasd.asdsad", // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("from publisher success"),
		})
	helpers.FailOnError(err, "Failed to publish a message")
}
