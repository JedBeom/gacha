package main

import "net/http"

func adminPage(w http.ResponseWriter, r *http.Request) {
	_, user, err := sessionAndUser(r)
	if err == nil && user.IsAdmin {
		var tdata TemplateData
		tdata.User = user
		generateHTML(w, tdata, "layout", "admin")
	} else {
		generateHTML(w, nil, "404")
	}
	return
}
