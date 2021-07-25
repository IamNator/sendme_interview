package schema

//User is used to store user data in database
type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	HashedPassword  string `json:"hashed_password"  validate:"required"`
	Token           string `json:"token"`
	TokenExpiration string `json:"token_expiration"`
}

func (User) TableName() string {
	return "user"
}
