package controllers

import (
	"fmt"
	"gorm_app/app/models"
	"gorm_app/config"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

var templates = template.Must(template.ParseFiles("app/views/edit.html", "app/views/index.html", "app/views/new.html", "app/views/show.html"))
var validPath = regexp.MustCompile("^/(edit|save|view|show)/([a-zA-Z0-9]+)$")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new.html", nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	article := &models.Article{Title: r.FormValue("title"), Body: []byte(body)}
	article.Create()
	templates.ExecuteTemplate(w, "index.html", nil)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	id, _ := strconv.Atoi(m[2])
	article := models.FindArticle(id)
	templates.ExecuteTemplate(w, "show.html", article)
}

func StartWebServer() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new/", newHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/show/", showHandler)
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil)
}
