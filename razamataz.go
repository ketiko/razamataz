package main

import (
	"fmt"
	"net/http"
)

func DoRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.WriteHeader(200)
	for key, values := range r.Form {
		w.Write([]byte(fmt.Sprintf("%s: %s \n", key, values)))
	}
}

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/", DoRequest)
	http.ListenAndServe(":9090", nil)
	fmt.Println("ending server")
}
