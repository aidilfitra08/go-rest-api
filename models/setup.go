package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin?parseTime=true"), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})

	DB = database
}