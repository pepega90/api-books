package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin_learn/models/book"
	"github.com/gin_learn/models/usermodel"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) AllBooks(c *gin.Context) {
	var books []book.Book
	books, _ = h.bookService.FindAll()
	c.JSON(http.StatusOK, books)
}

func (h *bookHandler) FindBookById(c *gin.Context) {
	var b book.Book
	bookId, _ := strconv.Atoi(c.Param("id"))
	b, _ = h.bookService.FindById(bookId)

	c.JSON(http.StatusOK, b)
}

func (h *bookHandler) HapusBuku(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	h.bookService.DeleteBook(id)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func (h *bookHandler) CreateBookHandler(c *gin.Context) {
	// title,  harga
	var bookReq book.BookRequest
	// ShouldBind(), menerima pointer untuk bind post form ke sebuah variabel struct
	if err := c.ShouldBind(&bookReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	current_user := c.MustGet("current_user")
	user := current_user.(usermodel.User)
	bookReq.UserId = user.Id

	book, err := h.bookService.Create(bookReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "kamus harus login"})
		return
	}

	user.Buku = append(user.Buku, book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) UpdateBuku(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateBook book.BookRequest

	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	b := h.bookService.Update(id, updateBook)
	c.JSON(http.StatusOK, b)
}
