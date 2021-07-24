package schema

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User is used to store user data in database
type User struct {
	gorm.Model
	UserName        string
	Email           string    `gorm:"type:varchar(100);unique_index"  validate:"required"`
	Gender          string    `json:"gender" validate:"required"`
	HashedPassword  string    `json:"hashed_password"  validate:"required"`
	Token           string    `json:"token"`
	TokenExpiration time.Time `json:"token_expiration"`
}

func (User) TableName() string {
	return "user"
}
