package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"parttwo/apiserver/heartbeat"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, err := storeObject(r.Body, object)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(c)

}

func putStream(object string) (*PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find avaliable dataServer")
	}
	return NewPutStream(server, object), nil
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, err := putStream(object)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	io.Copy(stream, r)
	return http.StatusOK, nil
}
