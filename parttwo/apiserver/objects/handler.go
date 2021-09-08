package objects

import (
	"fmt"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(os.Getwd())
	if r.Method == http.MethodGet {
		get(w, r)
		return
	}
	if r.Method == http.MethodPut {
		put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

}

//func put(w http.ResponseWriter, r *http.Request) {
//	file, err := os.Create("./" + "objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
//	if err != nil {
//		fmt.Println(err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	defer file.Close()
//	io.Copy(file, r.Body)
//
//}
//
//func get(w http.ResponseWriter, r *http.Request) {
//	file, err := os.Open("./" + "objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
//	if err != nil {
//		fmt.Println(err.Error())
//		w.WriteHeader(http.StatusNotFound)
//		return
//	}
//	defer file.Close()
//	io.Copy(w, file)
//
//}
