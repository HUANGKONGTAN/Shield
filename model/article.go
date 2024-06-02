package model

import (
	"Shield/tool"
	"Shield/viewModel"
	"fmt"
)

type Article struct {
	BaseModel
	Title      string `gorm:"index"`
	Content    string `gorm:"type:longtext"`
	AuthorID   int    `gorm:"default:4"`
	ChannelID  int    `gorm:"default:1"`
	Published  bool
	ReadAmount int `gorm:"default:0"`
	LikeAmount int `gorm:"default:0"`
}

// 新增文章
func InsertArticle(data *Article) int {
	fmt.Print(data)
	err := DB.Create(&data).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func UpdateArticle(article *Article) int {
	var oldArticle Article
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["content"] = article.Content

	err := DB.Model(&oldArticle).Where("id = ? ", article.ID).Updates(&maps).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func ListArticlesByChannel(id int, pageSize int, pageNumber int) ([]*viewModel.ViewArticle, int, int64) {
	var articles []*viewModel.ViewArticle
	var total int64

	err := DB.Model(&Article{}).Select("article.title, "+
		"article.id, article.content, article.read_amount, "+
		"article.like_amount, article.created_at, "+
		"channel.name as channel, user.nick_name as author").
		Joins("left join channel on article.channel_id = channel.id").
		Joins("left join user on article.author_id = user.id").
		Where("channel_id =?", id).
		Limit(pageSize).Offset((pageNumber - 1) * pageSize).Scan(&articles).Error
	DB.Model(&Article{}).Where("channel_id =?", id).Count(&total)
	if err != nil {
		return nil, tool.ERROR_CHANNEL_NOT_EXIST, 0
	}
	return articles, tool.SUCCSE, total
}

func ListArticles(pageSize int, pageNum int) ([]*viewModel.ViewArticle, int, int64) {
	var articles []*viewModel.ViewArticle
	var total int64

	err := DB.Model(&Article{}).Select("article.title, " +
		"article.id, article.read_amount, " +
		"article.like_amount, article.created_at "). // +
		///"channel.name as channel, user.nick_name as author").
		//Joins("left join channel on article.channel_id = channel.id").
		Joins("left join user on article.author_id = user.id").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&articles).Error
	DB.Model(&Article{}).Count(&total)
	if err != nil {
		return nil, tool.ERROR, 0
	}
	return articles, tool.SUCCSE, total
}

func ArticleByID(articleID interface{}) (*viewModel.ViewArticle, int) {
	var article *viewModel.ViewArticle
	err := DB.Model(&Article{}).Select("article.title, "+
		"article.content, article.read_amount, "+
		"article.like_amount, "+
		"user.nick_name as author").
		Joins("left join user on article.author_id = user.id").
		Where("article.id = ?", articleID).First(&article).Error
	if err != nil {
		return nil, tool.ERROR_ARTICLE_NOT_EXIST
	}
	return article, tool.SUCCSE
}

func DeleteArticle(id int) int {
	var article Article
	err := DB.Model(&Article{}).Where("id = ? ", id).Delete(&article).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func GiftArticle() (Article, int) {

	var gift Article
	err := DB.Last(&Article{}).Select("ID, Title").Scan(&gift).Error
	if err != nil {
		return gift, tool.ERROR
	}
	return gift, tool.SUCCSE
}

func FindArticle(keyWord string) ([]*viewModel.ViewArticle, int) {
	var articles []*viewModel.ViewArticle
	err := DB.Model(&Article{}).Select("article.title, "+
		"article.id, article.read_amount, "+
		"article.like_amount, article.created_at, "+
		"channel.name as channel, user.nick_name as author").
		Joins("left join channel on article.channel_id = channel.id").
		Joins("left join user on article.author_id = user.id").
		Where("article.title LIKE  ? ", "%"+keyWord+"%").Scan(&articles).Error
	if err != nil {
		return nil, tool.ERROR
	}
	return articles, tool.SUCCSE
}

func LikeArticle(id int) int {
	var article Article
	DB.Model(&Article{}).Select("like_amount").
		Where("id =  ? ", id).First(&article)
	err := DB.Model(&Article{}).
		Where("id =  ? ", id).Update("like_amount", article.LikeAmount+1).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func ReadArticle(id int) int {
	var article Article
	DB.Model(&Article{}).Select("read_amount").
		Where("id =  ? ", id).First(&article)
	err := DB.Model(&Article{}).
		Where("id =  ? ", id).Update("read_amount", article.ReadAmount+1).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}
