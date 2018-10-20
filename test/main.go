package main

import (
	"fmt"

	"github.com/JedBeom/gacha/data"
)

func main() {
	user := data.User{}

	user.Email = "solkblte@icloud.com"
	user.Password = "bjhbeom0908"
	user.Name = "Jed Beom"
	err := user.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Uuid)
}
