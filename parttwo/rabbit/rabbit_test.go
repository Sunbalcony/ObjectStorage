package rabbit

import (
	"fmt"
	"testing"
	"time"
)

func TestRabbitMQ_Publish(t *testing.T) {

	queue3 := New(Host)
	defer queue3.Close()
	for {
		time.Sleep(time.Millisecond)
		queue3.Publish("test", "123")

	}

}

func TestRecv(t *testing.T) {
	queue := New(Host)
	defer queue.Close()
	queue.Bind("test")

	consume := queue.Consume()
	//message := <-consume
	//var actual interface{}
	//err := json.Unmarshal(message.Body, &actual)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(actual)
	for i := range consume {
		fmt.Println(string(i.Body))
	}

}
