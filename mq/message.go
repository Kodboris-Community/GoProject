package mq

import (
	"github.com/streadway/amqp"
	"log"
	_ "log"
)

func ConnectToRabbitmq(mqUrl string) (string, error) {
	conn, err := amqp.Dial(mqUrl)
	if err != nil {
		log.Fatalf("Unable to connect to rabbitmq!! %v\n", err)
	}
	// Release mq connection
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return "Fail to create a Channel", err
	}
	defer ch.Close()
	return "Connection to rabbitmq is good!!!!\n", nil
}
