package main

import (
	"fmt"
	"net/http"

	"github.com/JedBeom/gacha/data"
)

func adminPage(w http.ResponseWriter, r *http.Request) {
	_, user, err := sessionAndUser(r)
	if err == nil && user.IsAdmin {
		tdata := AdminTemplate{CurrentUser: user}
		generateHTML(w, tdata, "layout", "admin")
	} else {
		generateHTML(w, nil, "404")
	}
	return
}

func userList(w http.ResponseWriter, r *http.Request) {
	_, user, _ := sessionAndUser(r)
	if user.IsAdmin {

		orderBy := r.FormValue("order_by")
		orderBy = orderBy

		users, err := data.ListUsers("id")
		fmt.Println(users)
		if err != nil || len(users) == 0 {
			generateHTML(w, nil, "404")
		}

		tdata := AdminTemplate{
			CurrentUser: user,
			Users:       users,
		}
		generateHTML(w, tdata, "layout", "user_list")
		return
	} else {
		generateHTML(w, nil, "404")
		return
	}
}
