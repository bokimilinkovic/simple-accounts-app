package handler

import (
	"fmt"
	"net/http"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// MoviesHandler is a handler that manages all movies calls.
type MoviesHandler struct {
	db *gorm.DB
}

// NewMovieHandler creates new movie handler.
func NewMovieHandler(db *gorm.DB) *MoviesHandler {
	return &MoviesHandler{db}
}

// GetMovies retrives all accounts from database.
func (a *MoviesHandler) GetMovies(c echo.Context) error {
	props, _ := c.Get("props").(jwt.MapClaims)
	fmt.Println(props["email"])
	fmt.Println(props["userId"])

	var movies []model.Movie
	if err := a.db.WithContext(c.Request().Context()).Find(&movies).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error getting all movies: "+err.Error())
	}

	return c.JSON(http.StatusOK, movies)
}
