package main

import (
	"fmt"
	"log"
	"net/http"
	"parttwo/dataserver/objects"
	"parttwo/dataserver/objects/heartbeat"
	"parttwo/dataserver/objects/locate"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Println("server listen 1234")
	log.Fatal(http.ListenAndServe(":1234", nil))

}
