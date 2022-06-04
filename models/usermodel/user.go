package usermodel

import (
	"github.com/gin_learn/models/book"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	Buku      []book.Book
}

func (u *User) SetPassword(pass string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	u.Password = hashPassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
