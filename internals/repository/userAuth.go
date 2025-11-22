package repository

import (
	"errors"
	"sync"

	"github.com/oloomoses/go-auth/internals/auth"
	"github.com/oloomoses/go-auth/internals/model"
	"golang.org/x/crypto/bcrypt"
)

//

type UserRepo struct {
	db    []model.User
	mutex sync.Mutex
}

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (r *UserRepo) CreateUser(newUser model.User) (model.User, string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, u := range r.db {
		if u.Username == newUser.Username {
			return model.User{}, "", errors.New("username already exits")
		}
	}

	hashedPass, err := hashPassword(newUser.Password)

	if err != nil {
		return model.User{}, "", err
	}

	hashedUser := model.User{
		Username: newUser.Username,
		Password: hashedPass,
	}

	token, err := auth.GenerateToken(hashedUser.Username)

	if err != nil {
		return model.User{}, "", err
	}

	r.db = append(r.db, hashedUser)

	return hashedUser, token, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (r *UserRepo) VerifyPassword(username string, password string) bool {
	for _, u := range r.db {
		if u.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
			return err == nil
		}
	}
	return false
}
