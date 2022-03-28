package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-to-do-list/config"
	"go-to-do-list/models"
	"go-to-do-list/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DBURL(config.SetupDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.TodoItem{})

	//setup routes
	r := routes.BuildRouter()

	// running
	r.Run()

}
