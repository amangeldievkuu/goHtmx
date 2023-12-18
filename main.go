package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct{
	Title string
	Director string
}

func h1(w http.ResponseWriter, r *http.Request){
	//main template
	tmpl := template.Must(template.ParseFiles("index.html"))

	//films data
	films := map[string][]Film{
		"Films": {
			{Title: "The Idiots", Director: "Brendon Moreno"},
			{Title: "Vlastelin Kolec", Director: "Justin Geidji"},
			{Title: "Boika",  Director: "Kolby Kovington"},
		},
	}
	//excute
	tmpl.Execute(w, films)
}

func h2(w http.ResponseWriter, r *http.Request){
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}

func h3(w http.ResponseWriter, r *http.Request){

}


func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	http.HandleFunc("/delete-item", h3)
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}