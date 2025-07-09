package user

import (
	usersDomain "github.com/Wookkie/notes-g2/internal/domain/users"
	"github.com/google/uuid"
)

type Repository interface { // описываем те методы, которые мы ожидаем от хранилища (от реализации БД)
	SaveUser(user usersDomain.User) error //string- это ID
	GetUser(login string) (usersDomain.User, error)
}

type UserService struct { //в ней должна быть структура, реализуящая базу данных
	repo Repository
}

func New(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) RegisterUser(user usersDomain.User) (string, error) {
	user.UID = uuid.New().String() //записываем строковое представление в наш ID

	err := us.repo.SaveUser(user)
	if err != nil {
		return ``, err
	}
	return user.UID, nil
}

func (us *UserService) LoginUser(userCreds usersDomain.UserRequest) (string, error) {
	dbUser, err := us.repo.GetUser(userCreds.Email)
	if err != nil {
		return ``, err
	}

	if dbUser.Password != userCreds.Password {
		return ``, usersDomain.ErrInvalidUserCreads
	}

	return dbUser.UID, nil
}
