package locate

import (
	"fmt"
	"os"
	"parttwo/rabbit"
	"strconv"
)

// Locate 返回文件是否存在
//用os.Stat访问磁盘上对应的文件名，用os.isNotExist判断文件是否存在，如果存在则定位成功返回true，负责false
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)

}

func StartLocate() {
	q := rabbit.New(rabbit.Host)
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	local, _ :=os.Getwd()
	fmt.Println("数据节点当前文件目录",local)
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(local + "\\objects\\" + object) {
			q.Send(msg.ReplyTo, "127.0.0.1:1234")
		}

	}

}
