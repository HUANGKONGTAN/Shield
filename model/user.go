package model

import (
	"Shield/tool"
)

type User struct {
	BaseModel
	UserName string
	PassWord string
	NickName string
}

// 新增用户
func InsertUser(data *User) int {
	var count int64
	DB.Model(&User{}).Where("user_name = ? ", data.UserName).Count(&count)
	if count != 0 {
		return tool.ERROR_USER_EXIST
	}else {
		err := DB.Create(&data).Error
		if err != nil {
			return tool.ERROR
		}else{
			return tool.SUCCSE
		}
	}
}

func UpdateUser(id int, user *User) int {
	var oldUser User
	var maps = make(map[string]interface{})
	maps["nickname"] = user.NickName
	maps["password"] = user.PassWord

	err := DB.Model(&oldUser).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func UserByID(userID int)(*User, int) {
	var user User
	err := DB.First(&user, userID).Error
	if err != nil {
		return nil, tool.ERROR_USER_NOT_EXIST
	}
	return &user, tool.SUCCSE
}

func DeleteUser(id int) int {
	var user User
	err := DB.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return tool.ERROR
	}
	return tool.SUCCSE
}

func VerifyUser(user *User) int {
	var origin *User
	err := DB.Debug().Model(&User{}).Where("user_name = ? ", user.UserName).First(&origin).Error
	if err != nil || origin == nil {
		return tool.ERROR_USER_NOT_EXIST
	}else {
		if origin.UserName == user.UserName && origin.PassWord == user.PassWord {
			return  tool.SUCCSE
		} else{
			return tool.ERROR_PASSWORD_WRONG
		}
	}
}