package api

import (
	"Shield/model"
	"Shield/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"strconv"
)

func NewUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := model.InsertUser(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

func VerifyUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code := model.VerifyUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

func UserById(c *gin.Context) {
	id, _ := c.GetQuery("id")
	var userId int
	userId, _ = strconv.Atoi(id)
	data, code := model.UserByID(userId)
	fmt.Print()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data": data,
		"message": tool.GetErrMsg(code),
	})
}

// 更新用户
func UpdateUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateUser(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 认证登陆
func Auth(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get("user")

	var code int
	fmt.Println(user)
	if user == nil {
		code = tool.NO_AUTHED
	}else {
		code = tool.AUTHED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}
