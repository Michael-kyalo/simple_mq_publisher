package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	println("simple RabbitMQ Publisher")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("connected closed")

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello, World!"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(" [x] Sent 'Hello, World!'")
}
