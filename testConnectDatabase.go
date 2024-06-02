package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	type User struct {
		ID       uint
		Username string
		Password string
	}

	dsn := "root:@tcp(127.0.0.1:3306)/shield?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	//user := User{Username: "LINGDINGYANG", Password: "nicheng1002", ID: 2}
	//a := User{Username: "MOSHENGDEXIAWU", Password: "nicheng1002", ID: 3}
	//result := db.Create(&user) //保存数据
	//db.Select("Username", "Password", "ID").Create(&a)
	//fmt.Print(result)
	_ = err

	var user User
	db.First(&user)
	fmt.Println(user.Username)

	var users []User
	db.Find(&users)
	//for  a := 0; a< len(users); a++ {
	//	fmt.Println(users[a].Username)
	//}
	//fmt.Println(users)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": users,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
