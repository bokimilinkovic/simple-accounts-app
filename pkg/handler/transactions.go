package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adjust/rmq"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// TransactionsHandler registers transactions handler
type TransactionsHandler struct {
	db    *gorm.DB
	queue rmq.Queue
}

type TransactionRequestDTO struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float32 `json:"amount"`
}

type TransactionResponseDTO struct {
	TransactionID string `json:"transactionID"`
	Status        string `json:"status"`
}

// NewTransactionsHandler creates new transactions handler
func NewTransactionsHandler(db *gorm.DB, queue rmq.Queue) *TransactionsHandler {
	return &TransactionsHandler{db, queue}
}

// CreateTransaction creates new transaction. Sender, receiver and amount is required.
func (th *TransactionsHandler) CreateTransaction(c echo.Context) error {
	var t TransactionRequestDTO
	if err := c.Bind(&t); err != nil {
		return c.String(http.StatusBadRequest, "error binding request:"+err.Error())
	}

	tx := th.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var sender model.Account
	if err := tx.WithContext(c.Request().Context()).Where("owner = ?", t.Sender).Find(&sender).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error fetching sender's account from DB: "+err.Error())
	}

	var receiver model.Account
	if err := tx.WithContext(c.Request().Context()).Where("owner = ?", t.Receiver).Find(&receiver).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error fetching receiver's account from DB: "+err.Error())
	}

	if sender.ID == 0 || receiver.ID == 0 {
		return c.String(http.StatusBadRequest, "check sender and receiver")
	}

	if sender.Balance < t.Amount {
		return c.String(http.StatusBadRequest, fmt.Sprintf("insuficient founds on sender's account: %s , %0.2f, need to send : %0.2f", sender.Owner, sender.Balance, t.Amount))
	}

	sender.Balance -= t.Amount
	receiver.Balance += t.Amount

	if err := tx.WithContext(c.Request().Context()).Save(&sender).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error updating sender's account: "+err.Error())
	}

	if err := tx.WithContext(c.Request().Context()).Save(&receiver).Error; err != nil {
		return c.String(http.StatusInternalServerError, "error updating receiver's account: "+err.Error())
	}

	// TODO: Send event in message queue so that new transaction can be created in transaction's database
	tr := model.Transaction{
		Sender:   sender.Owner,
		Receiver: receiver.Owner,
		Amount:   t.Amount,
	}
	payloadBytes, err := json.Marshal(&tr)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error marshaling transaction: "+err.Error())
	}

	if !th.queue.Publish(string(payloadBytes)) {
		log.Printf("failed to publish: %s", err)
	}

	if err := tx.Commit().Error; err != nil {
		return c.String(http.StatusInternalServerError, "error commiting error: "+err.Error())
	}

	resp := TransactionResponseDTO{
		Status: "PROCESSING",
	}

	return c.JSON(http.StatusCreated, resp)
}
