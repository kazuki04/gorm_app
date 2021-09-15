package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"gorm_app/config"
)

var templates = template.Must(template.ParseFiles("app/views/edit.html", "app/views/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplates() {

}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	templates.ExecuteTemplate(w, "view.html", nil)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
		}
		fn(w, r, m[2])
	}
}

func StartWebServer() error {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	return http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), nil)
}
