package models

import "github.com/IamNator/sendme_interview/internal/schema"

//DebitUser reduces a user's wallet balance
type DebitUser struct {
	UserID int `json:"user_id"`
	// TransactionID uint    `json:"transaction_id" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

//CreditUser increases a user's wallet balance
type CreditUser struct {
	UserID int `json:"user_id"`
	// TransactionID uint    `json:"transaction_id" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type TransactionStatement struct {
	schema.Wallet
}
