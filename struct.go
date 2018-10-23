package main

import (
	"github.com/JedBeom/gacha/data"
)

type TemplateData struct {
	CurrentUser data.User
	Message     string
}

type AdminTemplate struct {
	CurrentUser data.User
	Users       []data.User
}
