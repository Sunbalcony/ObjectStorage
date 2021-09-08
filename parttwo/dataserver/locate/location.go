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
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}

		if Locate("../" + "objects/" + object) {
			fmt.Println("../" + "objects/" + object)
			q.Send(msg.ReplyTo, "1.2.3.4")
		}

	}

}
