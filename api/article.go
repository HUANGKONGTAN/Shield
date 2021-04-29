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

func NewArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.InsertArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 查询所有文章
func Articles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNumber"))

	var code int
	var total int64
	var data [] *viewModel.ViewArticle

	id, _ := c.GetQuery("channelId")
	if id == "" {
		data, code, total = model.ListArticles(pageSize, pageNum)
	}else {
		channelId, _ := strconv.Atoi(id)
		data, code, total = model.ListArticlesByChannel(channelId, pageSize, pageNum)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": tool.GetErrMsg(code),
	})
}

// 查询单个文章信息
func Article(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, code := model.ArticleByID(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": tool.GetErrMsg(code),
	})
}

// 更新文章
func UpdateArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// DeleteArt 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Print(id)
	code := model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}

// 获取推送文章
func GiftArticle(c *gin.Context) {
	gift, code := model.GiftArticle()

	c.JSON(http.StatusOK, gin.H{
		"data": gift,
		"status":  code,
		"message": tool.GetErrMsg(code),
	})
}
