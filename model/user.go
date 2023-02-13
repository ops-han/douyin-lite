package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	FollowCount int32 `gorm:"type:int;default:0"`
	FollowerCount int32 `gorm:"type:int;default:0"`
}
