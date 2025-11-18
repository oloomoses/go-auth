package handler

import (
	"errors"
	"sync"

	"github.com/oloomoses/go-auth/internals/model"
	"golang.org/x/crypto/bcrypt"
)

var db []model.User
var mutex sync.Mutex

func CreateUser(newUser model.User) (model.User, error) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, u := range db {
		if u.Username == newUser.Username {
			return model.User{}, errors.New("username already exits")
		}
	}

	hashedPass, err := hashPassword(newUser.Password)

	if err != nil {
		return model.User{}, err
	}

	hashedUser := model.User{
		Username: newUser.Username,
		Password: hashedPass,
	}

	db = append(db, hashedUser)

	return hashedUser, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// func LoginUser(username string, password string) (mode.User error){

// }
