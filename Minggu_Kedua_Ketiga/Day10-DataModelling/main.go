package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/k0kubun/pp"

	"github.com/gorilla/mux"
)

// https://github.com/Torao-Law/B42-CH1-ST1---Day-2-Routing

type Card struct {
	Project_name string
	Start_date   string
	End_date     string
	Durasi       string
	Desc         string
	// handle checkbox with fix length array , not slices causing error when validating process in template
	Tech [4]string
	Img  string
}

// array dng index id dan isi interface cards

var Cards = []Card{
	{
		Project_name: "Kursus Golang sampai Botak",
		Start_date:   "11-10-2022",
		End_date:     "11-11-2022",
		Durasi:       "6 Bulan",
		Desc:         "lorem10",
		Tech:         [4]string{"nodejs", "java", "react", "ts"},
		Img:          "",
	},
	{
		Project_name: "Kursus Python sampai tipes ",
		Start_date:   "1-10-2022",
		End_date:     "1-12-2022",
		Durasi:       "3 Bulan",
		Desc:         "lorem10",
		Tech:         [4]string{"nodejs", "java", "react", "ts"},
		Img:          "",
	},
}

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
	// update Course
	route.HandleFunc("/updatecard/{id}", updatecard).Methods("GET")
	// Delete Course
	route.HandleFunc("/deletecard/{id}", deletecard).Methods("GET")

	// api handle form value & redirect to /
	route.HandleFunc("/sendform", sendform).Methods("POST")
	// handle page detail card
	route.HandleFunc("/detailcard/{id}", detailcard).Methods("GET")

	Port := "5500"

	fmt.Print("server running on port" + Port + "\n")

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

	pp.Print(response)

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

func updatecard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hallo from updatecard")
	tmpl, err := template.ParseFiles("public/updatecard.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {
		var UpdateData = Card{}
		// parse and assign index query url to variable
		index, _ := strconv.Atoi(mux.Vars(r)["index"])

		for i, data := range Cards {
			if index == i {
				UpdateData = Card{
					Project_name: data.Project_name,
					Start_date:   data.Start_date,
					End_date:     data.End_date,
					Durasi:       data.Durasi,
					Desc:         data.Desc,
					Tech:         data.Tech,
					Img:          "",
				}
				// notice this, still unresolve
				Cards = append(Cards[:index], Cards[index+1:]...)
			}
		}
		data := map[string]interface{}{
			"updateCard": UpdateData,
		}
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)
	}

}

func deletecard(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	// ProjectList = append(ProjectList[:index], ProjectList[index+1:]...)
	Cards = append(Cards[:index], Cards[index+1:]...)

	fmt.Println(Cards)

	http.Redirect(w, r, "/", http.StatusFound)

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

func detailcard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	template, err := template.ParseFiles("public/detailcard.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message:" + err.Error()))
		// return kosong spy tidak eksekusi baris kode dibawahnya
		return
	} else {
		// variable buat nampung isi detail cards
		var Container = Card{}
		// variable index parse id query ke string
		index, _ := strconv.Atoi(mux.Vars(r)["id"])

		for id, data := range Cards {
			// cek apakah id valid atau tidak,still doesnt work -_-
			if index != id {
				w.WriteHeader(http.StatusInternalServerError)
			}
			// cek apakah index sesuai query di URL
			if index == id {
				// masukkan data ke variabel penampung
				Container = Card{
					Project_name: data.Project_name,
					Start_date:   data.Start_date,
					End_date:     data.End_date,
					Durasi:       data.Durasi,
					Desc:         data.Desc,
					Tech:         data.Tech,
					Img:          "",
				}
			}
		}

		fmt.Println(Container.Tech)

		// buat variable map dengan key string interface kosong utk data passing ke template

		data := map[string]interface{}{
			"detailCards": Container,
		}

		// ksh tau browser request berhasil
		w.WriteHeader(http.StatusOK)
		// eksekusi template
		template.Execute(w, data)

		// fmt.Println(id)
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
