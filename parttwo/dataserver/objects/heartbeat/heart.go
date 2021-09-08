package heartbeat

import (
	"parttwo/rabbit"
	"time"
)

func StartHeartbeat() {
	q := rabbit.New("amqp://storage:storage@mid.low.im:5672")
	defer q.Close()
	for {
		q.Publish("apiServers", "1.2.3.4")
		time.Sleep(time.Second * 5)

	}

}
