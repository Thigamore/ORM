package test

import (
	orm "github.com/Thigamore/ORM/src"
)

type User struct {
	Table *orm.Table
	id    int    `json:"id"`
	email string `json:"email"`
}

var table = &orm.Table{Name: "user"}

func CreateUser()
