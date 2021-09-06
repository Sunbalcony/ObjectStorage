package main

import (
	"fmt"
	"net/http"
	"partone/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Println("server listen 1234")
	http.ListenAndServe(":1234", nil)

}
