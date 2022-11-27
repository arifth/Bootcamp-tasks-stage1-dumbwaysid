package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// https://github.com/Torao-Law/B42-CH1-ST1---Day-2-Routing

func main() {

	// list of endpoints :
	// "/",
	// "/addProject"
	// "/contact"
	//

	route := mux.NewRouter()

	// serve static files such as css,js and img

	fs := http.FileServer(http.Dir("./public/assets/"))

	route.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))

	// handle static files inside views folders with prefix

	route.PathPrefix("/Public").Handler(http.StripPrefix("public", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/addproject", addProject).Methods("GET")
	Port := "5000"

	fmt.Print("server running on port" + Port)

	http.ListenAndServe("localhost:"+Port, route)

}

func home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type:", "text/html")
	tmpt, err := template.ParseFiles("public/index.html")

	// cek apakah ada error ketika parsing html , jika ada kembalikan error ke console dan return kosong
	if err != nil {
		res.Write([]byte("request error cuy mas:" + err.Error()))
		//    kembalikan kosong
		return
	}

	// jika tidak ada error , maka eksekusi tmplate, tulis parsed html ke response obj
	tmpt.Execute(res, nil)

}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpt, err := template.ParseFiles("public/contact.html")

	if err != nil {
		fmt.Println("error")
	}

	tmpt.Execute(w, nil)

}

func addProject(w http.ResponseWriter, r *http.Request) {
	tmpt, _ := template.ParseFiles("public/projectform.html")
	w.Header().Set("Content-type:", "text/html")
	// if err != nil {
	// 	fmt.Println("error pakk")
	// 	return
	// }

	tmpt.Execute(w, nil)

}
