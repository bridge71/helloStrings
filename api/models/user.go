package models

type User struct {
	Nickname string `gorm:"column:nickname;unique;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	UserId   uint   `gorm:"column:userId;primaryKey"`
	Level    int    `gorm:"default:1"`
}
type UserAuth struct {
	PasswordHash string `gorm:"column:passwordHash"`
	UserId       uint   `gorm:"column:userId"`
}
type Message struct {
	RetMessage string
}
