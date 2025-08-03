package model

import (
	"github.com/thoseJanes/tinyblog/pkg/auth"
	"gorm.io/gorm"
)

type User struct {
	TimeAndId
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Nickname string `gorm:"column:nickname"`
	Email string `gorm:"column:email"`
	Phone string `gorm:"column:phone"`
}

func (*User) TableName() string {
	return "user_table"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	var err error
	u.Password, err = auth.Encrypt(u.Password)
	
	return err
}