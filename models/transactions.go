package models

//DebitUser reduces a user's wallet balance
type DebitUser struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount" validate:"required"`
}

//CreditUser increases a user's wallet balance
type CreditUser struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount" validate:"required"`
}
