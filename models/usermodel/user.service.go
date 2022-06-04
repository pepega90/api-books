package usermodel

type Service interface {
	Register(user UserInput) (User, error)
	Login(user UserInput) User
	GetUser(id int) User
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(user UserInput) (User, error) {
	u := User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	u.SetPassword(user.Password)

	createdUser, err := s.repository.Register(u)
	return createdUser, err
}

func (s *service) Login(user UserInput) User {
	u := User{
		Email:    user.Email,
		Password: []byte(user.Password),
	}

	loginUser := s.repository.Login(u)
	return loginUser
}

func (s *service) GetUser(id int) User {
	var user User
	user = s.repository.GetUser(id)
	return user
}
