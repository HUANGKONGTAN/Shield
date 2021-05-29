package model

import (
	"Shield/tool"
	"fmt"
)

type Music struct {
	BaseModel
	Name string `gorm:"index"`
	Lyric string 
	Singer string
}

func InsertMusic(data *Music) int {
	fmt.Print(data)
	err := DB.Create(&data).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func ListMusics(pageSize int, pageNum int) ([]*Music, int, int64) {
	var musics []*Music
	var total int64

	err := DB.Model(&Music{}).Select("music.name, " +
		"music.lyric, music.singer").
		Limit(pageSize).Offset((pageNum-1)*pageSize).Scan(&musics).Error
	DB.Model(&Music{}).Count(&total)
	if err != nil {
		return nil, tool.ERROR, 0
	}
	return musics, tool.SUCCSE, total
}

func MusicByID(musicID interface{})(*Music, int) {
	var music *Music
	err := DB.Model(&Music{}).Select("music.title, " +
		"music.file_name").
		Where("music.id = ?", musicID).First(&music).Error
	if err != nil {
		return nil, tool.ERROR_ARTICLE_NOT_EXIST
	}
	return music, tool.SUCCSE
}

func DeleteMusic(id int) int {
	var music Music
	err := DB.Model(&Music{}).Where("id = ? ", id).Delete(&music).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func GiftMusic() (Music,  int) {
	var gift Music
	err := DB.Last(&Music{}).Select("ID, Lyric").Scan(&gift).Error
	if err != nil {
		return gift, tool.ERROR
	}
	return gift, tool.SUCCSE
}