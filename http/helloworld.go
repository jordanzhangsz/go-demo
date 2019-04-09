package main

import (
	"fmt"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Request=%v", r)
	if r.Method == "GET" {
	   fmt.Fprintln(w, "Helloworld!!")
	}
}

func main() {
	http.HandleFunc("/", Helloworld)
	fmt.Println("Server is started at port 8080 ...")
	http.ListenAndServe(":8080", nil)
}
