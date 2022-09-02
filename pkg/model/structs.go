package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email            string
	PasswordHash     string
	Subscribed       bool
	SubscribedUntil  sql.NullTime
	StripeCustomerID string
	UserMovies       []Movie `gorm:"many2many:user_movies;"`
}

type Movie struct {
	gorm.Model
	Title       string
	Description string
	CoverURL    string
	Price       float32
}

type StripePayment struct {
	gorm.Model
	UserID  int    `gorm:"foreignKey"`
	MovieID int    `gorm:"foreignKey"`
	Status  string // one of [succeeded, processing, failed, requires_action]
}
