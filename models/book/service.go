package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	DeleteBook(id int)
	Update(id int, book BookRequest) Book
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) DeleteBook(id int) {
	s.repository.DeleteBook(id)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	buku := Book{
		Title:       bookRequest.Title,
		Price:       int(bookRequest.Price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		UserId:      bookRequest.UserId,
		Image:       bookRequest.Image,
	}

	newbook, err := s.repository.Create(buku)
	return newbook, err
}

func (s *service) Update(id int, book BookRequest) Book {
	buku := Book{
		Title:       book.Title,
		Price:       int(book.Price),
		Description: book.Description,
		Rating:      book.Rating,
		Image:       book.Image,
	}
	b := s.repository.Update(id, buku)
	return b
}
