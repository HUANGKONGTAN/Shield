package model

import (
	"Shield/tool"
	"fmt"
)

type Photo struct {
	BaseModel
	Title string `gorm:"index"`
	FileName string
}

// 新增文章
func InsertPhoto(data *Photo) int {
	fmt.Print(data)
	err := DB.Create(&data).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func ListPhotos(pageSize int, pageNum int) ([]*Photo, int, int64) {
	var photos []*Photo
	var total int64

	err := DB.Model(&Photo{}).Select("photo.title, " +
		"photo.id, photo.file_name").
		Limit(pageSize).Offset((pageNum-1)*pageSize).Scan(&photos).Error
	DB.Model(&Photo{}).Count(&total)
	if err != nil {
		return nil, tool.ERROR, 0
	}
	return photos, tool.SUCCSE, total
}

func PhotoByID(photoID interface{})(*Photo, int) {
	var photo *Photo
	err := DB.Model(&Photo{}).Select("photo.title, " +
		"photo.file_name").
		Where("photo.id = ?", photoID).First(&photo).Error
	if err != nil {
		return nil, tool.ERROR_ARTICLE_NOT_EXIST
	}
	return photo, tool.SUCCSE
}

func DeletePhoto(id int) int {
	var photo Photo
	err := DB.Model(&Photo{}).Where("id = ? ", id).Delete(&photo).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func GiftPhoto() (Photo,  int) {
	var gift Photo
	err := DB.Model(&Photo{}).Select("ID, Title").Scan(&gift).Error
	if err != nil {
		return gift, tool.ERROR
	}
	return gift, tool.SUCCSE
}