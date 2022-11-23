package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	fmt.Print("Hello from go")

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world "))
	})

	Port := "5000"

	fmt.Print("server running on port" + Port)

	http.ListenAndServe("localhost:"+Port, route)

}
