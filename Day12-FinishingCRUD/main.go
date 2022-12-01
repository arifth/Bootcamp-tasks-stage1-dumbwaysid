package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pengenalan_golang/connection"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
)

// https://github.com/Torao-Law/B42-CH1-ST1---Day-2-Routing

type Card struct {
	Id                int
	Project_name      string
	Start_date        time.Time
	End_date          time.Time
	Start_date_Parsed string
	End_date_Parsed   string
	Durasi            string
	Desc              string
	// handle checkbox with fix length array , not slices causing error when validating process in template
	Tech [4]string
	Img  string
}

// array dng index id dan isi interface cards
// Global Variabel Penampung DB

var Cards = []Card{}

// RENDER DATA DENGAN AWALAN HURUF KAPITAL

func main() {

	// list of endpoints :
	// "/",
	// "/addProject"
	// "/contact"
	//

	// konnek db dlu baru instantiate route dng mux
	connection.ConnectDB()

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
	// send to form
	// route.HandleFunc("/updatecard/{id}", updatecard).Methods("GET")
	// Delete Course
	route.HandleFunc("/deletecard/{id}", deletecard).Methods("GET")

	// api handle to update card
	route.HandleFunc("/updating/{id}", updating).Methods("GET")

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

	// query := "SELECT Project_name,Start_date,End_date,Durasi,Desc,Tech,Img FROM tb_courses"
	query := "SELECT * FROM tb_courses"
	// query := "SELECT id, Project_name, Start_date, End_date, Durasi, Desc, Tech, Img FROM tb_courses"

	dataCards, err := connection.Conn.Query(context.Background(), query)

	// cek apakah query berhasil
	if err != nil {
		fmt.Println("tidak bisa kueri tabel" + err.Error())
		return
	}

	var queried []Card

	// method return true ,used for loop inside query if there are more than one value
	for dataCards.Next() {

		var each = Card{}

		err := dataCards.Scan(&each.Id, &each.Project_name, &each.Start_date, &each.End_date, &each.Durasi, &each.Desc, &each.Tech, &each.Img)
		if err != nil {
			fmt.Println("Message:" + err.Error())
		} else {
			//parse the date from db
			parsed := Card{
				Id:           each.Id,
				Project_name: each.Project_name,
				Durasi:       each.Durasi,
				Desc:         each.Desc,
				Tech:         each.Tech,
				Img:          each.Img,
			}

			// pp.Println(parsed)
			queried = append(queried, parsed)

		}

		// pp.Println(queried)
		// queried = append(queried, each)
	}
	// Cards = append(Cards, queried...)

	// pp.Println(queried)

	response := map[string]interface{}{
		"Cards": queried,
	}

	// pp.Print(response)

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
	// parse file template from html
	tmpt, _ := template.ParseFiles("public/projectform.html")
	w.Header().Set("Content-type:", "text/html")

	// if err != nil {
	// 	fmt.Println("error pakk")
	// 	return
	// }

	tmpt.Execute(w, nil)

}
func updating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("public/updatecard.html")
	// ERROR HANDLING RENDER HTML TEMPLATE
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	} else {
		ID, _ := strconv.Atoi(mux.Vars(r)["id"])
		updateData := Card{}
		// GET PROJECT BY ID FROM POSTGRESQL
		err = connection.Conn.QueryRow(context.Background(), `SELECT * FROM public.tb_courses WHERE "id" = $1`, ID).Scan(&updateData.Id, &updateData.Project_name, &updateData.Start_date, &updateData.End_date, &updateData.Durasi, &updateData.Desc, &updateData.Tech, &updateData.Img)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message : " + err.Error()))
			return
		}

		updateData = Card{
			Id:                updateData.Id,
			Project_name:      updateData.Project_name,
			Start_date_Parsed: ReturnDate(updateData.Start_date),
			End_date_Parsed:   ReturnDate(updateData.End_date),
			Desc:              updateData.Desc,
			Tech:              updateData.Tech,
			Img:               updateData.Img,
		}

		response := map[string]interface{}{
			"updateData": updateData,
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, response)
	}

}

func deletecard(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// pp.Println("test 1")

	// DELETE PROJECT BY ID AT POSTGRESQL
	_, err := connection.Conn.Exec(context.Background(), `DELETE FROM public.tb_courses WHERE "id" = $1`, Id)
	// ERROR HANDLING DELETE PROJECT AT POSTGRESQL
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}
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
		start_date_prs, _ := time.Parse("2006-01-02", r.FormValue("start-date"))
		end_date_prs, _ := time.Parse("2006-01-02", r.FormValue("end-date"))
		end_date := r.FormValue("end-date")
		desc := r.FormValue("desc")
		// img := r.FormValue("img")
		duration := getDuration(start_date, end_date)
		nodejs := r.PostForm.Get("nodejs")
		java := r.PostForm.Get("java")
		react := r.PostForm.Get("react")
		ts := r.PostForm.Get("ts")
		Tech := [4]string{nodejs, java, react, ts}
		// Img := "test.png"

		query := `INSERT INTO public.tb_courses("Project_name", "Start_date","End_date", "Durasi", "Desc", "Tech" ,"Img" )VALUES ($1, $2, $3, $4, $5, $6, $7)`

		// queri insert into postgress cuy
		_, err = connection.Conn.Exec(context.Background(), query, project_name, start_date_prs, end_date_prs, duration, desc, Tech, "test.png")

		if err != nil {
			fmt.Println("tidak bisa kueri tabel" + err.Error())
			return

		}
		// _, err = connection.Conn.Exec(context.Background(),
		//  `INSERT INTO public.tb_project()
		// 	VALUES ( $1, $2, $3, $4, $5, $6)`, ProjectName, ProjectStartDate, ProjectEndDate, ProjectDescription, ProjectTechnologies, ProjectImage)
		// // ERROR HANDLING INSERT PROJECT TO POSTGRESQL

		// item := Card{
		// 	Id:           Id,
		// 	Project_name: project_name,
		// 	// Start_date:   formatDate(start_date),
		// 	// End_date:     formatDate(end_date),
		// 	Durasi: getDuration(start_date, end_date),
		// 	Desc:   desc,
		// 	Img:    img,
		// 	Tech:   [4]string{nodejs, java, react, ts},
		// }

		// fmt.Print(item.Tech)
		// Cards = append(Cards, item)

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
		ID, _ := strconv.Atoi(mux.Vars(r)["id"])
		updateData := Card{}
		// GET PROJECT BY ID FROM POSTGRESQL
		err = connection.Conn.QueryRow(context.Background(), `SELECT * FROM public.tb_courses WHERE "id" = $1`, ID).Scan(&updateData.Id, &updateData.Project_name, &updateData.Start_date, &updateData.End_date, &updateData.Durasi, &updateData.Desc, &updateData.Tech, &updateData.Img)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message : " + err.Error()))
			return
		}

		updateData = Card{
			Id:                updateData.Id,
			Project_name:      updateData.Project_name,
			Start_date_Parsed: ReturnDate(updateData.Start_date),
			End_date_Parsed:   ReturnDate(updateData.End_date),
			Durasi:            updateData.Durasi,
			Desc:              updateData.Desc,
			Tech:              updateData.Tech,
			Img:               updateData.Img,
		}

		pp.Println(updateData)

		response := map[string]interface{}{
			"updateData": updateData,
		}

		w.WriteHeader(http.StatusOK)
		template.Execute(w, response)

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

// format date type to golang
func ParseDate(InputDate time.Time) string {

	Formated := InputDate.Format("02 January 2006")

	return Formated
}

// format date to potgres
func ReturnDate(InputDate time.Time) string {

	Formated := InputDate.Format("2006-01-02")

	return Formated
}
