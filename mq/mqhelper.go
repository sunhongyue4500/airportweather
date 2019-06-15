package mq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 消息队列连接字符串
var mqURL string
var conn *amqp.Connection
var mqch amqp.Channel

func FailOnError(err error, msg string) {
	if err != nil {

	}
}

func ConnMQ(usr string, pwd string, host string, port int) error {
	mqURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", usr, pwd, host, port)
	var err error
	conn, err = amqp.Dial(mqURL)
	return err
}

func Send2MQ(msg []byte) {
	mqch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer mqch.Close()

	err = mqch.ExchangeDeclare(
		"topic_exchange1_publish", // name
		"topic",                   // type
		false,                     // durable
		false,                     // auto-deleted
		false,                     // internal
		false,                     // no-wait
		nil,                       // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	err = mqch.Publish(
		"topic_exchange1_publish", // exchange
		"AirportWeather",          // routing key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	FailOnError(err, "Failed to publish a message")

	log.Printf("发送到消息队列: %s", msg)
}
