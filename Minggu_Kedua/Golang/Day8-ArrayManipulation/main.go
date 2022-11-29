package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// https://github.com/Torao-Law/B42-CH1-ST1---Day-2-Routing

type Card struct {
	Project_name string
	Start_date   string
	End_date     string
	Durasi       string
	Desc         string
	Tech         [4]string
	Img          string
}

// array dng index id dan isi interface cards

var Cards = []Card{}

// RENDER DATA DENGAN AWALAN HURUF KAPITAL

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

	// base route
	route.HandleFunc("/", home).Methods("GET")
	// contact page
	route.HandleFunc("/contact", contact).Methods("GET")
	// add Course
	route.HandleFunc("/addproject", addProject).Methods("GET")
	// api handle form value & redirect to /
	route.HandleFunc("/sendform", sendform).Methods("POST")
	Port := "5000"

	fmt.Print("server running on port" + Port)

	http.ListenAndServe("localhost:"+Port, route)

}

func home(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type:", "text/html")
	tmpt, err := template.ParseFiles("public/index.html")

	// cek apakah ada error ketika parsing html , jika ada kembalikan error ke console dan return kosong
	if err != nil {
		res.Write([]byte("request error:" + err.Error()))
		//    kembalikan kosong
		return
	}

	response := map[string]interface{}{
		"Cards": Cards,
	}

	// fmt.Println(response)

	res.WriteHeader(http.StatusOK)
	// jika tidak ada error , maka eksekusi tmplate, tulis parsed html ke response obj
	tmpt.Execute(res, response)

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

func sendform(w http.ResponseWriter, r *http.Request) {

	// fmt.Print("hallo form sendform")
	// parse request form and body,lalu catch error dan return http error 5xx cuy

	err := r.ParseMultipartForm(1024)

	if err != nil {
		log.Fatal(err)
	} else {
		project_name := r.FormValue("name")
		start_date := r.FormValue("start-date")
		end_date := r.FormValue("end-date")
		desc := r.FormValue("desc")
		img := r.FormValue("img")
		nodejs := r.PostForm.Get("nodejs")
		java := r.PostForm.Get("java")
		react := r.PostForm.Get("react")
		ts := r.PostForm.Get("ts")

		item := Card{
			Project_name: project_name,
			Start_date:   formatDate(start_date),
			End_date:     formatDate(end_date),
			Durasi:       getDuration(start_date, end_date),
			Desc:         desc,
			Img:          img,
			Tech:         [4]string{nodejs, java, react, ts},
		}

		fmt.Print(item.Tech)
		Cards = append(Cards, item)

		// fmt.Print(len(Cards))
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}

func getDuration(start_date string, end_date string) string {
	layout := "2006-01-02"
	date1, _ := time.Parse(layout, start_date)
	date2, _ := time.Parse(layout, end_date)

	hasil := date2.Sub(date1).Hours() / 23
	var durasi string

	if hasil > 30 {
		if (hasil / 30) < 1 {
			durasi = "1 bulan"
		} else {
			durasi = strconv.Itoa(int(hasil)/30) + "Bulan"
		}

	}
	return durasi

}

func formatDate(date string) string {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, date)

	hasil := t.Format("02 January 2006")
	return hasil
}

// func handleDuration(start string, end string, r *http.Request) string {
// 	r.ParseForm()
// 	start = r.FormValue("start-date")
// 	end = r.FormValue("end-date")
