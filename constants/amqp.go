package constants

const (
	AMQP_QUEUE_NAME = "learning-rabbitmq.queue"
	// AMQP_QUEUE_NAME = ""
	// direct exchange
	AMQP_DIRECT_EXCHANGE_NAME        = "learning-rabbitmq-direct.exchange"
	AMQP_DIRECT_EXCHANGE_ROUTING_KEY = "learning-rabbitmq-direct.routing-key"
	// fanout exchange
	AMQP_FANOUT_EXCHANGE_NAME = "learning-rabbitmq-fanout.exchange"
)
