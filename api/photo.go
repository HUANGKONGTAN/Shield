package api

import (
	"Shield/model"
	"Shield/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewPhoto(c *gin.Context) {
	var data model.Photo
	_ = c.ShouldBindJSON(&data)
	code := model.InsertPhoto(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 获取所有照片
func Photos(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNumber"))

	data, code, total := model.ListPhotos(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": tool.GetErrMsg(code),
	})
}

// 查询单张照片
func Photo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.PhotoByID(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

func DeletePhoto(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Print(id)
	code := model.DeletePhoto(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 获取推送文章
func GiftPhoto(c *gin.Context) {
	gift, code := model.GiftPhoto()

	c.JSON(http.StatusOK, gin.H{
		"data": gift,
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}
