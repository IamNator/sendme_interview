package schema

import "github.com/jinzhu/gorm"

//Transaction keeps track of every transaction that occurs
type Transaction struct {
	gorm.Model
	// TransactionID uint    `json:"transaction_id" validate:"required"`
	UserID uint    `json:"user_id" validate:"required"`
	Type   string  `json:"type" validate:"required" gorm:"type:enum('credit','debit')"` //enum("debit", "credit")
	Amount float64 `json:"amount" validate:"required"`
}

func (Transaction) TableName() string {
	return "user"
}
