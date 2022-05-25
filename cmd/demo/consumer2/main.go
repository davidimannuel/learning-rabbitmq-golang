package main

import (
	"errors"
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
		"demo.queue2", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	helpers.FailOnError(err, "failed to declare a queue")

	err = ch.ExchangeDeclare(
		"demo.exchange",     // name
		amqp.ExchangeFanout, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	helpers.FailOnError(err, "Failed to declare an exchange")

	routingKey := "key.error"
	err = ch.QueueBind(
		q.Name,          // queue name
		routingKey,      // routing key -- The meaning of a binding key (routing key) depends on the exchange type
		"demo.exchange", // exchange
		false,
		nil,
	)
	helpers.FailOnError(err, "failed to bind a queue")

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
	log.Println("listener to queue", q.Name, "routing key", routingKey)
	go func() {
		for d := range msgs {
			log.Printf("Received delivery tag %d with message: %s", d.DeliveryTag, d.Body)
			// log.Println("process finished")
			// err := CheckMessages(string(d.Body))
			// if err != nil {
			// 	d.Reject(true)
			// } else {
			// 	d.Ack(false)
			// }
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func CheckMessages(msg string) error {
	if msg == "success" {
		return nil
	}
	return errors.New("invalid message")
}
