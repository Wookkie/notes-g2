package inmemory

import (
	usersDomain "github.com/Wookkie/notes-g2/internal/domain/users"
)

var emptyUser = usersDomain.User{}

type InMemory struct {
	userStorage map[string]usersDomain.User //ключ - это UID пользователя, значение - это наша структура user
}

func New() *InMemory { // экземпляр нашего хранилища
	return &InMemory{
		userStorage: make(map[string]usersDomain.User), //мапу мы обязательно дожны проинициализировать
	}
}

func (im InMemory) SaveUser(user usersDomain.User) error {
	for _, us := range im.userStorage {
		if us.Email == user.Email { // если 2 пользователя с одинаковыми логинами (почтами), то возвращаем ошибку
			return usersDomain.ErrUserAlreadyExists
		}
	}

	im.userStorage[user.UID] = user // кладем в него нашего user
	return nil
}

func (im InMemory) GetUser(login string) (usersDomain.User, error) {
	for _, us := range im.userStorage {
		if us.Email == login { // возвращаем пустую структуру, если не нашли такого логина, а если нашли
			return us, nil
		}
	}

	//если блок не сработал (не нашли такого пользователя), то возвращаем:
	return emptyUser, usersDomain.ErrUserNotFound
}
