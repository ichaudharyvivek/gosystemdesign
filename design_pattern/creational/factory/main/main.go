package main

import (
	"design-patterns/creational/factory"
	"fmt"
)

func main() {
	db := factory.NewDBFactory("mysql", "localhost:3306")
	if db == nil {
		fmt.Println("DB not initialized")
	}

	db.Connect()
	db.ExecuteQuery("SELECT * FROM users")
}
