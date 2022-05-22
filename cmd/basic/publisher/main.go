package main

import (
	"learning-rabbitmq/configs"
	"learning-rabbitmq/constants"
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

	q, err := ch.QueueDeclare(
		constants.AMQP_QUEUE_NAME, // name
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	helpers.FailOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("from publisher"),
		})
	helpers.FailOnError(err, "Failed to publish a message")
}
