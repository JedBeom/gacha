package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/auth", authHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/logout", logout)

	mux.HandleFunc("/register", register)
	mux.HandleFunc("/join", join)

	mux.HandleFunc("/admin", adminPage)
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		generateHTML(w, nil, "404")
		return
	}
	_, user, err := sessionAndUser(r)
	var tdata TemplateData
	if err == nil {
		tdata.User = user
	}
	generateHTML(w, tdata, "layout", "content")
}
