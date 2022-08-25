package main

import (
	"GoFiles/myFirstWeb/models"
	"crypto/rand"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	posts map[string]*models.Post
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html", "static/header.html", "static/footer.html")
	if err != nil {
		log.Fatal("indexHandler", err)
	}

	fmt.Println(posts)

	tmpl.ExecuteTemplate(w, "index", posts)
	//	fmt.Fprintf(w, "Hello")
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/write.html", "static/header.html", "static/footer.html")
	if err != nil {
		log.Fatal("writeHandler", err)

	}

	//fmt.Println(posts)

	tmpl.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/write.html", "static/header.html", "static/footer.html")
	if err != nil {
		log.Fatal("editHandler", err)
	}

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
	}
	tmpl.ExecuteTemplate(w, "write", post)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}
	delete(posts, id)
	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listen on port:")

	posts = make(map[string]*models.Post, 0)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deletePostHandler)
	http.HandleFunc("/SavePost", savePostHandler)
	http.ListenAndServe(":9090", nil)
}

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
