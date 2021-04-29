package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB *gorm.DB
)

func InitDB()(err error) {
	db,err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/shield?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	_ = db.AutoMigrate(&User{}, &Article{}, &Channel{})

	if err == nil {
		log.Println("connect db success")
		DB = db
	}
	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {
	err := InitDB()
	checkError(err)
}