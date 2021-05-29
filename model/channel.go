package model

import (
	"Shield/tool"
)

type Channel struct {
	BaseModel
	Name string
	Describe string
	Type int
}

// 新增频道
func InsertChannel(data *Channel) int {
	err := DB.Create(&data).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func UpdateChannel(id int, channel *Channel) int {
	var oldChannel Channel
	var maps = make(map[string]interface{})
	maps["name"] = channel.Name
	maps["describe"] = channel.Describe

	err := DB.Model(&oldChannel).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func ListChannels(Type string) ([]*Channel, int, int64) {
	var channels []*Channel
	var total int64

	var TypeCode int
	if Type == "article" {
		TypeCode = 0
	}else if Type == "sundry" {
		TypeCode = 1
	}
	err := DB.Where("type = ? ", TypeCode).Find(&channels).Error
	DB.Model(&channels).Count(&total)
	if err != nil {
		return nil, tool.ERROR, 0
	}
	return channels, tool.SUCCSE, total
}

func ChannelByID(channelID interface{})(*Channel, int) {
	var channel Channel
	err := DB.First(&channel, channelID).Error
	if err != nil {
		return nil, tool.ERROR_ARTICLE_NOT_EXIST
	}
	return &channel, tool.SUCCSE
}

func DeleteChannel(id int) int {
	var channel Channel
	err := DB.Where("id = ? ", id).Delete(&channel).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}