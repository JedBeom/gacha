package main

import (
	"net/http"
	"time"

	"github.com/JedBeom/gacha/data"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		http.Redirect(w, r, "/login?msg=err", 401)
	}

	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			http.Redirect(w, r, "login?msg=err", 401)
		}

		cookie := http.Cookie{
			Name:     "_yayoiori",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)

	} else {
		http.Redirect(w, r, "/login?msg=err", 302)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("msg")
	if message == "err" {
		message = "Email or Password Wrong."
	} else if message == "registered" {
		message = "Register Complete! Please Login."
	} else {
		message = ""
	}
	var tdata TemplateData
	tdata.Message = message
	_, user, err := sessionAndUser(r)
	if err == nil {
		tdata.CurrentUser = user
	}
	generateHTML(w, tdata, "layout", "login")
}

func register(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("msg")
	if message == "err" {
		message = "Problem occured. Please retry after checking your infomation."
	} else {
		message = ""
	}

	var tdata TemplateData
	tdata.Message = message
	generateHTML(w, tdata, "layout", "register")
}

func join(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := data.User{}
	user.Email = r.PostFormValue("email")
	user.Name = r.PostFormValue("name")
	user.StudentID = r.PostFormValue("student_id")
	user.Password = r.PostFormValue("password")

	err := user.Create()
	if err != nil {
		http.Redirect(w, r, "/register?msg=err", 302)
		return
	}

	http.Redirect(w, r, "/login?msg=registered", 302)
	return
}

func logout(w http.ResponseWriter, r *http.Request) {
	sess, _, err := sessionAndUser(r)
	if err == nil {
		cookie, err := r.Cookie("_yayoiori")
		if err != nil {
			return
		}
		sess.Remove()
		cookie.Expires = time.Unix(0, 0)
		http.SetCookie(w, cookie)
	}

	http.Redirect(w, r, "/", 302)

	return
}
