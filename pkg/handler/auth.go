package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/database"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/util/password"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthHandler struct {
	userStore database.UserStoreInterface // this way it can be easily switched
	secretKey []byte                      // should be in config file, but for demo perposes...
}

func NewAuthHandler(db *gorm.DB, secretKey []byte) *AuthHandler {
	return &AuthHandler{
		userStore: database.NewUserStore(db),
		secretKey: secretKey,
	}
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email  string `json:"email"`
	UserID int    `json:"userId"`
	jwt.StandardClaims
}

// Signin checks if the user exists in the DB, if he is not found
// create new user entry in DB, and return token.
func (a *AuthHandler) Signin(c echo.Context) error {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		return c.String(http.StatusBadRequest, "bad request body provided: "+err.Error())

	}

	// Get the expected password from our in memory map
	user, err := a.getUserOrCreateNewOne(creds)
	if err != nil {
		return c.String(http.StatusBadRequest, "error getting or creating new user: "+err.Error())

	}

	if !password.CheckPasswordHash(creds.Password, user.PasswordHash) {
		return c.String(http.StatusUnauthorized, "bad password provided")

	}

	// Declare the expiration time of the token
	// here, we have kept it as 60 minutes
	expirationTime := time.Now().Add(60 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email:  creds.Email,
		UserID: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			Id: fmt.Sprint(user.ID),
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(a.secretKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return c.String(http.StatusInternalServerError, "can not sign token")
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(c.Response(), &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	lr := loginResp{Token: tokenString}
	return c.JSON(http.StatusOK, lr)
}

type loginResp struct {
	Token string `json:"token"`
}

func (a *AuthHandler) getUserOrCreateNewOne(creds Credentials) (*model.User, error) {
	user, err := a.userStore.GetUserByEmail(creds.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		hashedPass, err := password.HashPassword(creds.Password)
		if err != nil {
			return nil, fmt.Errorf("error hasing: %w", err)
		}
		newUser := model.User{
			Email:        creds.Email,
			PasswordHash: hashedPass,
		}

		return a.userStore.CreateUser(newUser)
	}

	return user, nil
}
