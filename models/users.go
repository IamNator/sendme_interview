package models

//LoginCredential is used for logging in
type LoginCredential struct {
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

//RegistrationCredential is used for registration
type RegistrationCredential struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

//LoginResponse is the response after a successful login or registration
type LoginResponse struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email"  validate:"required"`
	Token     string `json:"token" validate:"token"`
}
