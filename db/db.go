package db

import (
	"gosearch/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitDb()  {
	// replace password
	var dsn = "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if(err != nil){
		panic(err)
	}

	DB = *db
	db.AutoMigrate(&model.User1{})
}