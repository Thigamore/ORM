package main

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Id    int    `db:"id"`
	Age   int    `db:"age"`
	Email string `db:"email"`
}

func main() {
	fmt.Println("Initializing Db")

	db, err := sql.Open("mysql", "Vscode:@tcp(127.0.0.1:3306)/test_learning")
	if err != nil {
		panic(err)
	}

	rs, err := db.Query("SELECT * FROM user")

	thom := user{
		Id:    0,
		Age:   17,
		Email: "a@gmail.com",
	}

	fmt.Println(reflect.TypeOf(thom).Field(0).Tag.Get("db"))

	for rs.Next() {
		var user user
		if err != nil {
			panic(err)
		}
		fmt.Println(user)

	}
}
