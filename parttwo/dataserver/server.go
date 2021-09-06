package main

import (
	"fmt"
	"log"
	"net/http"
	"parttwo/dataserver/objects"
	"parttwo/dataserver/objects/heartbeat"
)

func main() {
	go heartbeat.StartHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Println("server listen 1234")
	log.Fatal(http.ListenAndServe(":1234", nil))

}
