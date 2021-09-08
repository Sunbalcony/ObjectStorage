package main

import (
	"fmt"
	"net/http"
	"parttwo/apiserver/locate"
	"parttwo/apiserver/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/",locate.Handler)
	fmt.Println("server listen 1234")
	http.ListenAndServe(":1235", nil)

}
