package api

import (
	"Shield/model"
	"Shield/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewMusic(c *gin.Context) {
	var data model.Music
	_ = c.ShouldBindJSON(&data)
	code := model.InsertMusic(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 获取所有照片
func Musics(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNumber"))

	data, code, total := model.ListMusics(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": tool.GetErrMsg(code),
	})
}

// 查询单张照片
func Music(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.MusicByID(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

func DeleteMusic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Print(id)
	code := model.DeleteMusic(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 获取推送文章
func GiftMusic(c *gin.Context) {
	gift, code := model.GiftMusic()

	c.JSON(http.StatusOK, gin.H{
		"data": gift,
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

//// 搜索音乐
//func FindMusic(c *gin.Context) {
//	musicList, code := model.FindMusic()
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": musicList,
//		"status":  code,
//		"message": tool.GetErrMsg(code),
//	})
//}
