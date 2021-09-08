package main

import (
	"fmt"
	"net/http"
	"parttwo/apiserver/heartbeat"
	"parttwo/apiserver/locate"
	"parttwo/apiserver/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	fmt.Println("server listen 1235")
	http.ListenAndServe(":1235", nil)

}
