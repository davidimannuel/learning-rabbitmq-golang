package main

import (
	"learning-rabbitmq/configs"
	"learning-rabbitmq/helpers"
	"log"

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
		"demo.queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	helpers.FailOnError(err, "failed to declare a queue")

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

	// err = ch.QueueBind(
	// 	q.Name, // queue name
	// 	constants.AMQP_DIRECT_EXCHANGE_ROUTING_KEY, // routing key -- The meaning of a binding key (routing key) depends on the exchange type
	// 	constants.AMQP_DIRECT_EXCHANGE_NAME,        // exchange
	// 	false,
	// 	nil,
	// )
	// helpers.FailOnError(err, "failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helpers.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	log.Println("listener to queue", q.Name)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
