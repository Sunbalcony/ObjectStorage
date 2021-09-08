package locate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parttwo/rabbit"
	"strconv"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	fmt.Println(info)
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bytes, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)

}

// Locate 定位对象  传入需要定位对象的文件名，他会创建一个新的消息队列，并向dataServers exchange群发这个对象名字的定位信息
func Locate(name string) string {
	q := rabbit.New("amqp://storage:storage@mid.low.im:5672")
	defer q.Close()
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		//设置超时机制，避免无休止的等待，1s后消息队列关闭
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s

}

// Exist 检查Locate结果是否为空字符串来判定对象是否存在
func Exist(name string) bool {
	return Locate(name) != ""

}
