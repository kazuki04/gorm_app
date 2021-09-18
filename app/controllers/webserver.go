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
	article = article.Create()
	fmt.Println(article.ID)
	http.Redirect(w, r, fmt.Sprintf("/show/%d", int(article.ID)), http.StatusFound)
}

func showHandler(w http.ResponseWriter, r *http.Request, article models.Article) {
	templates.ExecuteTemplate(w, "show.html", article)
}

func editHandler(w http.ResponseWriter, r *http.Request, article models.Article) {
	templates.ExecuteTemplate(w, "edit.html", article)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, models.Article)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(m)
		id, _ := strconv.Atoi(m[2])
		fmt.Println(id)
		article := models.FindArticle(id)
		fn(w, r, article)
	}

}

func StartWebServer() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new/", newHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/show/", makeHandler(showHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil)
}
