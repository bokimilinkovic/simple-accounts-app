package handler

import (
	"net/http"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountDTO struct {
	Owner   string  `json:"owner"`
	Balance float32 `json:"balance"`
}

// AccountHandler is a handler that manages all account calls.
type AccountHandler struct {
	db *gorm.DB
}

// NewAccountHandler creates new account handler.
func NewAccountHandler(db *gorm.DB) *AccountHandler {
	return &AccountHandler{db}
}

// GetAccount retrives all accounts from database.
func (a *AccountHandler) GetAccounts(c echo.Context) error {
	var accounts []model.Account
	if err := a.db.WithContext(c.Request().Context()).Find(&accounts).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error getting all accounts: "+err.Error())
	}

	return c.JSON(http.StatusOK, accounts)
}

// CreateAccount creates new account.
func (a *AccountHandler) CreateAccount(c echo.Context) error {
	var account AccountDTO
	if err := c.Bind(&account); err != nil {
		return c.String(http.StatusBadRequest, "error binding body: "+err.Error())
	}

	accountDb := &model.Account{
		Owner:   account.Owner,
		Balance: account.Balance,
	}

	if err := a.db.WithContext(c.Request().Context()).Create(&accountDb).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error creating new account: "+err.Error())
	}

	return c.JSON(http.StatusCreated, accountDb)
}

// GetAccount retrives one account based on ID.
func (a *AccountHandler) GetAccount(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "no Id provided")
	}

	var account model.Account
	if err := a.db.WithContext(c.Request().Context()).Find(&account, "id = ?", id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error getting one account: "+err.Error())
	}

	return c.JSON(http.StatusOK, account)
}

// DeleteAccount soft deletes one account based on ID.
func (a *AccountHandler) DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "no Id provided")
	}

	if err := a.db.WithContext(c.Request().Context()).Delete(&model.Account{}, "id = ?", id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error deleting account: "+err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// UpdateAccount updates account's balance and owner.
func (a *AccountHandler) UpdateAccount(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "no Id provided")
	}

	var accountDTO AccountDTO
	if err := c.Bind(&accountDTO); err != nil {
		return c.String(http.StatusBadRequest, "error binding body: "+err.Error())
	}

	// first fetch account from DB
	var account model.Account
	err := a.db.WithContext(c.Request().Context()).First(&account, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "account with that ID does not exist.")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "unknown error getting account from DB: "+err.Error())
	}

	account.Balance = accountDTO.Balance
	if accountDTO.Owner != "" {
		account.Owner = accountDTO.Owner
	}

	if err := a.db.WithContext(c.Request().Context()).Save(&account).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error updating account in DB: "+err.Error())
	}

	return c.JSON(http.StatusOK, account)
}
