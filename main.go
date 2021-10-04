package main

import (
	"gombit-publisher/config"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(config.URL)
	__(err)
	defer conn.Close()

	ch, err := conn.Channel()
	__(err)
	defer ch.Close()

	que, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	__(err)

	body := "Hello World!"

	err = ch.Publish(
		"",
		que.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	__(err)

}

func __(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
