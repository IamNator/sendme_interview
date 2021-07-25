package schema

import "time"

//Transaction keeps track of every transaction that occurs
type Transaction struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	UserID    int        `json:"user_id" validate:"required"`
	Type      string     `json:"type" validate:"required"` //enum("debit", "credit")
	Amount    float64    `json:"amount" validate:"required"`
}

func (Transaction) TableName() string {
	return "transaction"
}
