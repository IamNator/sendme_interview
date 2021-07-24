package schema

import "github.com/jinzhu/gorm"

type Wallet struct {
	gorm.Model
	UserID  uint    `json:"user_id" validate:"required"`
	Balance float64 `json:"balance" validate:"required"`
}

func (Wallet) TableName() string {
	return "user"
}
