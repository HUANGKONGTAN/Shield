package api

import (
	"Shield/model"
	"Shield/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewChannel(c *gin.Context) {
	var data model.Channel
	_ = c.ShouldBindJSON(&data)

	code := model.InsertChannel(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 查询所有频道
func Channels(c *gin.Context) {
	var code int
	var total int64
	var data [] *model.Channel

	data, code, total = model.ListChannels()

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": tool.GetErrMsg(code),
	})
}

//// 查询频道信息
//func Channel(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//	data, code := model.ChannelByID(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": tool.GetErrMsg(code),
//	})
//}

// 更新频道
func UpdateChannel(c *gin.Context) {
	var data model.Channel
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateChannel(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 删除频道
func DeleteChannel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteChannel(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

