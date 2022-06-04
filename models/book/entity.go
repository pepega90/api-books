package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Image       string
	UserId      uint `json:"user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
