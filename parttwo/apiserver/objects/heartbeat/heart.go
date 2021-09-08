package heartbeat

import (
	"math/rand"
	"parttwo/rabbit"
	"strconv"
	"sync"
	"time"
)

var dataServers = make(map[string]time.Time)

var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbit.New("amqp://storage:storage@mid.low.im:5672")
	defer q.Close()
	q.Bind("apiServer")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

//定时扫描清除超过10秒没收到心跳消息的数据服务节点
func removeExpiredDataServer() {
	for {
		time.Sleep(time.Second * 10)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}

}

// GetDataServers 遍历并返回当前所有的数据服务节点
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}

}

// ChooseRandomDataServer 从当前数据节点任意选择一个返回，如果当前数据节点为空，则返回空字符串
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]

}
