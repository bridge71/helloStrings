package models

import (
	"time"
)

type User struct {
	// Nickname string `gorm:"column:nickname;unique;not null" `
	// Email    string `gorm:"column:email;unique;not null" `
	Nickname     string `gorm:"column:nickname;unique;not null" json:"nickname"`
	Email        string `gorm:"column:email;unique;not null" json:"email"`
	PasswordHash string `gorm:"column:passwordHash" json:"password"`

	UserId uint `gorm:"column:userId;primaryKey"`
	Level  int  `gorm:"default:1"`
}

//	type UserAuth struct {
//		Password string `gorm:"column:passwordHash" json:"password"`
//		Nickname string `gorm:"column:nickname;unique;not null" json:"nickname"`
//		UserId   uint   `gorm:"column:userId"`
//	}
type Message struct {
	RetMessage string
	BookSale   []BookSale
	User       User
}

type BookSale struct {
	CreatedAt  time.Time
	BookName   string `json:"bookName" gorm:"column:book_name;"`
	Author     string `json:"author" gorm:"column:author;"`
	Profession string `json:"profession" gorm:"column:profession;"`
	Course     string `json:"course" gorm:"column:course;"`
	Common     bool   `json:"common" gorm:"column:common;"`
	IsSold     bool   `gorm:"default:false"`
	UserId     uint   `json:"userId" gorm:"column:userId;"`
	Value      uint   `json:"value" gorm:"column:value;"`
}
