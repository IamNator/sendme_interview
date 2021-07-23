package schema

import (
	"github.com/jinzhu/gorm"
)

//User is used to store user data in database
type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"type:varchar(100);unique_index"  validate:"required"`
	Gender         string `json:"Gender" validate:"required"`
	HashedPassword string `json:"HashedPassword"  validate:"required"`
}
