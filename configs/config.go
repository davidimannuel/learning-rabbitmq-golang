package configs

import "fmt"

var Amqp = AmqpConfig{
	Host:     "127.0.0.1",
	Port:     5672,
	User:     "guest",
	Password: "guest",
}

type AmqpConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func (a AmqpConfig) GetConnectionURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		a.User,
		a.Password,
		a.Host,
		a.Port,
	)
}
