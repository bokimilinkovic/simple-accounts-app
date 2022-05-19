package model

import (
	"time"
)

type TransactionStatus int64

const (
	Processing TransactionStatus = iota
	Successfull
	Failed
)

// Transaction represents one transaction in DB, new entity is created whenever some balances have been exchanged.
type Transaction struct {
	ID        string `gorm:"primaryKey"`
	Sender    string
	Receiver  string
	Amount    float32
	Status    TransactionStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
