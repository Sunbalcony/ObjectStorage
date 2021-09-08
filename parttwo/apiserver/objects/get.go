package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"parttwo/apiserver/locate"
	"strings"
)

//object是一个字符串，它代表对象的名字
//我们首先调用locate.Locate定位这个对象，如果返回的数据节点为空字符串，则返回定位失败的错误，否则就调用NewGetStream返回其结果
func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	fmt.Println("数据节点:",server)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return NewGetStream(server, object)

}

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	fmt.Println("文件名:",object)
	stream, err := getStream(object)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)

}
