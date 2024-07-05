package models

import "time"

type User struct {
	// Nickname string `gorm:"column:nickname;unique;not null" `
	// Email    string `gorm:"column:email;unique;not null" `
	Nickname     string `gorm:"column:nickname;unique;not null" json:"nickname"`
	Email        string `gorm:"column:email;unique;not null" json:"email"`
	PasswordHash string `gorm:"column:passwordHash" json:"password"`
	UserId       uint   `gorm:"column:userId;primaryKey" json:"userId"`
	Level        int    `gorm:"default:1" json:"level"`
}

type IP struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	IP        string    `gorm:"column:IP" json:"ip"`
	Province  string    `gorm:"column:province" json:"province"`
	City      string    `gorm:"column:city" json:"city"`
	District  string    `gorm:"column:district" json:"district"`
	Lat       float32   `gorm:"column:lat" json:"lat"`
	Lng       float32   `gorm:"column:lng" json:"lng"`
	UserId    uint      `gorm:"column:userId" json:"userId"`
}
type Message struct {
	RetMessage  string
	BookSale    []BookSale
	Post        []Post
	Comment     []Comment
	CommentMark []CommentMark
	IP          []IP
	Likes       []Likes
	PostContent PostContent
	User        User
}

type BookSale struct {
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Title      string    `json:"title" gorm:"column:title;"`
	Author     string    `json:"author" gorm:"column:author;"`
	Profession string    `json:"profession" gorm:"column:profession;default:公共课"`
	Course     string    `json:"course" gorm:"column:course;"`
	Common     bool      `json:"common" gorm:"column:common;"`
	IsSold     bool      `gorm:"default:false"`
	UserId     uint      `json:"userId" gorm:"column:userId;"`
	Value      uint      `json:"value" gorm:"column:value;"`
}

type BookBy struct {
	Key string `json:"key"`
	By  string `json:"by"`
}

type Post struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	Title     string    `json:"title" gorm:"column:title;"`
	Nickname  string    `gorm:"column:nickname;not null" json:"nickname"`
	PostId    uint      `gorm:"column:postId;primaryKey" json:"postId"`
	Likes     int       `gorm:"column:likes;" json:"likes"`
	IsShown   bool      `gorm:"default:true" json:"isShown"`
	UserId    uint      `json:"userId" gorm:"column:userId;"`
}

type PostContent struct {
	Content string `gorm:"column:content" json:"content"`
	PostId  uint   `gorm:"column:postId"`
}

type Comment struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	Nickname  string    `gorm:"column:nickname;not null" json:"nickname"`
	Content   string    `gorm:"column:content" json:"content"`
	PostId    uint      `gorm:"column:postId"`
	UserId    uint      `json:"userId" gorm:"column:userId;"`
	IsShown   bool      `gorm:"default:true" json:"isShown"`
}

type Likes struct {
	PostId uint `gorm:"column:postId"`
	UserId uint `json:"userId" gorm:"column:userId;"`
}

type CommentMark struct {
	PostId uint `gorm:"column:postId"`
	UserId uint `json:"userId" gorm:"column:userId;"`
}
