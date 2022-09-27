package models

import (
	"github.com/sakria9/web-print-server/db"
)

type User struct {
	Email    string `json:"email" gorm:"primary_key" binding:"required,email"`
	MaxPage  int    `json:"max_page"`
	Admin    bool   `json:"admin"`
	Password string `json:"password"`
}

func (u *User) Create() error {
	return db.GetDB().Create(u).Error
}

func (u *User) Update() error {
	return db.GetDB().Save(u).Error
}

func (u *User) Get() error {
	return db.GetDB().First(u).Error
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := db.GetDB().Find(&users).Error
	return users, err
}

func (u *User) GetByEmail() error {
	return db.GetDB().Where("email = ?", u.Email).First(u).Error
}
