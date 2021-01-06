package service

import (
	"errors"
	"holycode-task/model"
	"holycode-task/repository/postgres"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store *postgres.Store
}

func NewUserService(store *postgres.Store) *UserService {
	return &UserService{store: store}
}

func (u *UserService) FindByID(id uint) (*model.User, error) {
	return u.store.FindUserByID(id)
}

func (u *UserService) CreateUser(user *model.User) error {
	existingUser, err := u.FindByID(uint(user.ID))
	if existingUser == nil && gorm.IsRecordNotFoundError(err) {
		return u.store.CreateUser(user)
	}

	return err
}

func (u *UserService) FindAll() ([]model.User, error) {
	return u.store.FindAll()
}

func (u *UserService) Authentificate(username, password string) (*model.User, error) {
	existing, err := u.store.FindUserByUsername(username)
	if existing == nil || gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user does not exists :" + err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(password)); err != nil {
		return nil, errors.New("Bad password provided")
	}

	return existing, nil
}
