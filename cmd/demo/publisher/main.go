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

	// make sure queue already declare
	q, err := ch.QueueDeclare(
		"demo.queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	helpers.FailOnError(err, "Failed to declare a queue")

	// err = ch.ExchangeDeclare(
	// 	constants.AMQP_QUEUE_NAME, // name
	// 	amqp.ExchangeDirect,       // type
	// 	true,                      // durable
	// 	false,                     // auto-deleted
	// 	false,                     // internal
	// 	false,                     // no-wait
	// 	nil,                       // arguments
	// )
	// helpers.FailOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key , we can publish without declare queue first, QueueDeclare just make sure queue already there
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("from publisher"),
		})
	helpers.FailOnError(err, "Failed to publish a message")
}
