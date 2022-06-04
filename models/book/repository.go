package book

import "gorm.io/gorm"

// repository layer
type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
	DeleteBook(id int)
	Update(id int, book Book) Book
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindById(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) DeleteBook(id int) {
	r.db.Delete(&Book{ID: id})
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(id int, book Book) Book {
	b, _ := r.FindById(id)
	b.Title = book.Title
	b.Description = book.Description
	b.Rating = book.Rating
	b.Price = book.Price
	b.Image = book.Image
	r.db.Save(&b)
	return b
}
