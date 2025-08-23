package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"www.github.com/Aryan1365/day3/Config"
	models "www.github.com/Aryan1365/day3/Models"
	"www.github.com/Aryan1365/day3/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DBURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&models.Student{}, &models.Mark{})
	r := Routes.SetUpRouter()
	r.Run(":8080")
}
