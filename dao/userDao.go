package dao

import (
	"douyin-lite/model"
	"gorm.io/gorm"
)

func IsNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	return user.ID != 0
}

func FindByUsername(db *gorm.DB, name string) model.User {
	var user model.User
	db.Where("name = ?", name).First(&user)
	return user
}
