package heartbeat

import (
	"parttwo/rabbit"
	"time"
)

// StartHeartbeat 向MQ中apiServers的exchange发送本节点监听地址
func StartHeartbeat() {
	q := rabbit.New("amqp://storage:storage@mid.low.im:5672")
	defer q.Close()
	for {
		q.Publish("apiServers", "1.2.3.4")
		time.Sleep(time.Second * 5)

	}

}

