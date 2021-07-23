type schema


import "github.com/jinzhu/gorm"


type Wallet struct {
	gorm.Model
	LastTransactionID uint `json:"last_transaction_id" validate:"required"`
	Balance  float64 `json:"balance" validate:"required"`
}