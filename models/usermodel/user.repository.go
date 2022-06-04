package usermodel

import (
	"gorm.io/gorm"
)

type Repository interface {
	Register(user User) (User, error)
	Login(user User) User
	GetUser(id int) User
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Login(user User) User {
	var u User
	r.db.Where("email = ?", user.Email).First(&u)

	return u
}

func (r *repository) GetUser(id int) User {
	u := User{
		Id: uint(id),
	}
	r.db.Preload("Buku").Find(&u)
	return u
}
