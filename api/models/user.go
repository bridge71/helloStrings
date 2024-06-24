package models

type User struct {
	// ID       uint   `gorm:"column:userId;primaryKey"`
	Nickname string `gorm:"column:nickname;unique;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Level    int    `gorm:"column:level"`
}
