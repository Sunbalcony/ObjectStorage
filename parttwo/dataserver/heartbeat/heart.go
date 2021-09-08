package heartbeat

import (
	"fmt"
	"parttwo/rabbit"
	"time"
)

// StartHeartbeat 向MQ中apiServers的exchange发送本节点监听地址
func StartHeartbeat() {
	q := rabbit.New(rabbit.Host)
	defer q.Close()
	i := 0
	for {
		q.Publish("apiServers", "127.0.0.1:1234")
		time.Sleep(time.Second * 5)
		i += 1
		fmt.Println("心跳发送中", i)

	}

}
