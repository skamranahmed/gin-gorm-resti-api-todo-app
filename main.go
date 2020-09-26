package main

import (
	"Volumes/Kamran/go-lang-todo-app/Config"
	"Volumes/Kamran/go-lang-todo-app/Models"
	"Volumes/Kamran/go-lang-todo-app/Routes"
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {

	// Creating a connection to DB
	Config.DB, err = gorm.Open("mysql", Config.DBUrl(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup routes
	r := Routes.SetupRouter()

	// running
	r.Run()
}
