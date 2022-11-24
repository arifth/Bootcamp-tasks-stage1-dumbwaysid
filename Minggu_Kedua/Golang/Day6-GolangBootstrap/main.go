package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// https://github.com/Torao-Law/B42-CH1-ST1---Day-2-Routing

func main() {
	route := mux.NewRouter()

	// handle static files inside views folders with prefix

	route.PathPrefix("/Public").Handler(http.StripPrefix("public", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	// route.HandleFunc('/listBlog',listBlog()).Methods("GET")
	// route.HandleFunc('/BlogDetail',BlogDetail()).Methods("GET")
	Port := "5000"

	fmt.Print("server running on port" + Port)

	http.ListenAndServe("localhost:"+Port, route)

}

func home(req http.ResponseWriter, res *http.Request) {
	req.Header().Set("Content-type:", "text/html")
	tmpt, err := template.ParseFiles("public/index.html")

	// cek apakah ada error ketika parsing html , jika ada kembalikan error ke console dan return kosong
	if err != nil {
		req.Write([]byte("request error cuy:" + err.Error()))
		//    kembalikan kosong
		return
	}

	// jika tidak ada error , maka eksekusi tmplate, tulis parsed html ke response obj
	tmpt.Execute(req, nil)

}

// func listBlog() {

// }

// func BlogDetail() {

// }
