package models

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
	User       User
}
