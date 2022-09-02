package database

import (
	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	"gorm.io/gorm"
)

// UserStoreInterface iterface that holds all methods that needs to be implemented
type UserStoreInterface interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUsers() ([]model.User, error)
	CreateUser(user model.User) (*model.User, error)
}

// UserStore wrapper around database layer, responsible for user managment in database
type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStoreInterface {
	return &UserStore{db}
}

// GetUserByEmail returns user by email, if exists
func (u *UserStore) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}

// GetUsers gets all users
func (u *UserStore) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser creates new user.
func (u *UserStore) CreateUser(user model.User) (*model.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
