package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

const Host = "amqp://storage:storage@mid.low.im:5672"

type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

// New 函数用于创建一个新的rabbitmq.RabbitMQ结构体
//该结构体的Bind方法可以将自己的消息队列和一个exchange绑定，所有发往该exchange的消息都能在自己的消息队列中被接收到
func New(s string) *RabbitMQ {
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	queue, err := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = queue.Name
	return mq

}

func (q *RabbitMQ) Bind(exchange string) {
	err := q.channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil)
	if err != nil {
		panic(err)
	}
	q.exchange = exchange
}

// Send 可以往某个消息队列发数据
func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = q.channel.Publish("", queue, false, false, amqp.Publishing{ReplyTo: q.Name, Body: []byte(str)})
	if err != nil {
		panic(err)
	}

}

// Publish 向exchange发消息
func (q *RabbitMQ) Publish(exchange string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = q.channel.Publish(exchange, "", false, false, amqp.Publishing{ReplyTo: q.Name, Body: []byte(str)})
	if err != nil {
		panic(err)
	}

}

// Consume 生成一个接收消息的go channel
func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, err := q.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	return c
}

func (q *RabbitMQ) Close() {
	q.channel.Close()

}
