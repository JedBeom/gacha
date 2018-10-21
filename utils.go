package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/JedBeom/gacha/data"
)

func generateHTML(w http.ResponseWriter, content interface{}, fileName ...string) {
	var files []string
	for _, file := range fileName {
		files = append(files, fmt.Sprintf("tmpl/%s.tmpl", file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", content)
}

func sessionAndUser(w http.ResponseWriter, r *http.Request) (user data.User, session data.Session, err error) {
	cookie, err := r.Cookie("_yayoiori")
	if err == nil {
		session = data.Session{Uuid: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("Invaild session")
		}
		user, err = data.UserByEmail(session.Email)
	}
	return
}
