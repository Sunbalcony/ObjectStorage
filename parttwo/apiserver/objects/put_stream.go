package objects

import (
	"fmt"
	"io"
	"net/http"
)

// PutStream 一个结构体，内含一个io.pipeWriter的指针和一个error的channel
//writer用于实现Write方法，c用于把一个goroutine传输数据过程中发生的错误传回主线程
type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

// NewPutStream 用于生成一个PutStream的结构体。它用io.pipe创建了一堆reader和writer，类型分别是*io.pipeReader和*io.pipeWriter
//他们是管道互联的，写入的writer的内容可以从reader中读出来
func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		client := http.Client{}
		response, err := client.Do(request)
		if err != nil && response.StatusCode != http.StatusOK {
			fmt.Println(err)
			err = fmt.Errorf("dataServer return http code %d", response.StatusCode)

		}
		c <- err
	}()
	return &PutStream{writer, c}
}

func (w *PutStream) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)

}

func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c

}
