package model

import (
	"Shield/tool"
	"Shield/viewModel"
	"fmt"
)

type Sundry struct {
	BaseModel
	Title string `gorm:"index"`
	Content string `gorm:"type:longtext"`
	AuthorID int
	ChannelID int
	Published bool
	ReadAmount int `gorm:"default:0"`
	LikeAmount int `gorm:"default:0"`
}

// 新增文章
func InsertSundry(data *Sundry) int {
	fmt.Print(data)
	err := DB.Create(&data).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func UpdateSundry( sundry *Sundry) int {
	var oldSundry Sundry
	var maps = make(map[string]interface{})
	maps["title"] = sundry.Title
	maps["content"] = sundry.Content

	err := DB.Model(&oldSundry).Where("id = ? ", sundry.ID).Updates(&maps).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func ListSundriesByChannel(id int, pageSize int, pageNumber int) ([]*viewModel.ViewSundry, int, int64) {
	var sundries []*viewModel.ViewSundry
	var total int64

	err := DB.Model(&Sundry{}).Select("sundry.title, " +
		"sundry.id, sundry.content, sundry.read_amount, " +
		"sundry.like_amount, sundry.created_at, " +
		"channel.name as channel, user.nick_name as author").
		Joins("left join channel on sundry.channel_id = channel.id").
		Joins("left join user on sundry.author_id = user.id").
		Where("channel_id =?", id).
		Limit(pageSize).Offset((pageNumber-1)*pageSize).Scan(&sundries).Error
	DB.Model(&Sundry{}).Where("channel_id =?", id).Count(&total)
	if err != nil {
		return nil, tool.ERROR_CHANNEL_NOT_EXIST, 0
	}
	return sundries, tool.SUCCSE, total
}

func ListSundries(pageSize int, pageNum int) ([]*viewModel.ViewSundry, int, int64) {
	var sundries []*viewModel.ViewSundry
	var total int64

	err := DB.Model(&Sundry{}).Select("sundry.title, " +
		"sundry.id, sundry.content, sundry.read_amount, " +
		"sundry.like_amount, sundry.created_at, " +
		"channel.name as channel, user.nick_name as author").
		Joins("left join channel on sundry.channel_id = channel.id").
		Joins("left join user on sundry.author_id = user.id").
		Limit(pageSize).Offset((pageNum-1)*pageSize).Scan(&sundries).Error
	DB.Model(&Sundry{}).Count(&total)
	if err != nil {
		return nil, tool.ERROR, 0
	}
	return sundries, tool.SUCCSE, total
}

func SundryByID(sundryID interface{})(*viewModel.ViewSundry, int) {
	var sundry *viewModel.ViewSundry
	err := DB.Model(&Sundry{}).Select("sundry.title, " +
		"sundry.content, sundry.read_amount, " +
		"sundry.like_amount, " +
		"user.nick_name as author").
		Joins("left join user on sundry.author_id = user.id").
		Where("sundry.id = ?", sundryID).First(&sundry).Error
	if err != nil {
		return nil, tool.ERROR_ARTICLE_NOT_EXIST
	}
	return sundry, tool.SUCCSE
}

func DeleteSundry(id int) int {
	var sundry Sundry
	err := DB.Model(&Sundry{}).Where("id = ? ", id).Delete(&sundry).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func GiftSundry() (Sundry,  int) {
	var gift Sundry
	err := DB.Last(&Sundry{}).Select("ID, Title").Scan(&gift).Error
	if err != nil {
		return gift, tool.ERROR
	}
	return gift, tool.SUCCSE
}

func FindSundry(keyWord string)([]*viewModel.ViewSundry , int){
	var sundries []*viewModel.ViewSundry
	err := DB.Model(&Sundry{}).Select("sundry.title, " +
		"sundry.id, sundry.read_amount, " +
		"sundry.like_amount, sundry.created_at, " +
		"channel.name as channel, user.nick_name as author").
		Joins("left join channel on sundry.channel_id = channel.id").
		Joins("left join user on sundry.author_id = user.id").
		Where("sundry.title LIKE  ? ", "%"+keyWord+"%").Scan(&sundries).Error
	if err != nil {
		return nil, tool.ERROR
	}
	return sundries, tool.SUCCSE
}

func LikeSundry(id int) (int) {
	var sundry Sundry
	DB.Model(&Sundry{}).Select("like_amount").
		Where("id =  ? ", id).First(&sundry)
	err := DB.Model(&Sundry{}).
		Where("id =  ? ", id).Update("like_amount", sundry.LikeAmount+1).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}


func ReadSundry (id int) (int) {
	var sundry Sundry
	DB.Model(&Sundry{}).Select("read_amount").
		Where("id =  ? ", id).First(&sundry)
	err := DB.Model(&Sundry{}).
		Where("id =  ? ", id).Update("read_amount", sundry.ReadAmount+1).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}