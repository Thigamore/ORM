package main

import (
	"github.com/Thigamore/ORM/src"
	db "github.com/Thigamore/ORM/src/db"
)

func main() {
	conf, err := src.GetConfig()
	if err != nil {
		panic(err)
	}
	err = db.InitDb("bookings", conf)
	if err != nil {
		panic(err)
	}
}
