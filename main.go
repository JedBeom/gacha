package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
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

	mux.HandleFunc("/admin/user_list", userList)
	mux.HandleFunc("/admin", adminPage)
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("オハヨ, /\\ |_ | <-> <>!")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("ダンス グッナイ☆スターズ")
		os.Exit(1)
	}()

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		generateHTML(w, nil, "404")
		return
	}
	var tdata TemplateData
	_, tdata.CurrentUser, _ = sessionAndUser(r)
	generateHTML(w, tdata, "layout", "content")
}
