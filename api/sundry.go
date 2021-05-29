package api

import (
	"Shield/model"
	"Shield/tool"
	"Shield/viewModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewSundry(c *gin.Context) {
	var data model.Sundry
	_ = c.ShouldBindJSON(&data)
	code := model.InsertSundry(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 查询所有文章
func Sundries(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNumber"))

	var code int
	var total int64
	var data [] *viewModel.ViewSundry

	id, _ := c.GetQuery("channelId")
	if id == "" {
		data, code, total = model.ListSundries(pageSize, pageNum)
	}else {
		channelId, _ := strconv.Atoi(id)
		data, code, total = model.ListSundriesByChannel(channelId, pageSize, pageNum)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": tool.GetErrMsg(code),
	})
}

// 查询单个文章信息
func Sundry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.SundryByID(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 更新文章
func UpdateSundry(c *gin.Context) {
	var data model.Sundry
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateSundry(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// DeleteArt 删除文章
func DeleteSundry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Print(id)
	code := model.DeleteSundry(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 获取推送文章
func GiftSundry(c *gin.Context) {
	gift, code := model.GiftSundry()

	c.JSON(http.StatusOK, gin.H{
		"data": gift,
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 搜索文章
func FindSundry(c *gin.Context) {
	keyWord := c.Query("keyWord")
	sundries, code := model.FindSundry(keyWord)

	c.JSON(http.StatusOK, gin.H{
		"data": sundries,
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

func LikeSundry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	fmt.Println(id)
	code := model.LikeSundry(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

func ReadSundry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	code := model.ReadSundry(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}