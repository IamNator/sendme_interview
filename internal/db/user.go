package dao

import "github.com/jinzhu/gorm"

type User struct {
	DB *gorm.DB
}

//CustomQuery ...
func (u User) CustomQuery(query string, values ...interface{}) (interface{}, error) {

	var output interface{}
	result := u.DB.Raw(query, values).Find(&output)
	if er := result.Error; er != nil {
		return nil, er
	}

	return output, nil
}
