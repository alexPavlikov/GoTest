package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"D:/GoFiles/myFirstWeb/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html", "static/header.html", "static/footer.html")
	if err != nil {
		log.Fatal("indexHandler", err)
	}
	tmpl.ExecuteTemplate(w, "index", nil)
	//	fmt.Fprintf(w, "Hello")
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/write.html", "static/header.html", "static/footer.html")
	if err != nil {
		log.Fatal("writeHandler", err)

	}
	tmpl.ExecuteTemplate(w, "write", nil)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Print(id, title, content)
}

func main() {
	fmt.Println("Listen on port:")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)
	http.ListenAndServe(":9090", nil)
}
