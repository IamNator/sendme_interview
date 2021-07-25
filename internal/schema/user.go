package schema

import (
	"time"
)

//User is used to store user data in database
type User struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	UserName        string     `json:"username"`
	Email           string     `json:"email"`
	HashedPassword  string     `json:"hashed_password"  validate:"required"`
	Token           string     `json:"token"`
	TokenExpiration time.Time  `json:"token_expiration"`
}

func (User) TableName() string {
	return "user"
}
