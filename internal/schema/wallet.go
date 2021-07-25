package schema

import "time"

type Wallet struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	UserID    uint       `json:"user_id" validate:"required"`
	Balance   float64    `json:"balance" validate:"required"`
}

func (Wallet) TableName() string {
	return "user"
}
