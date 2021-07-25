package models

//LoginCredential is used for logging in
type LoginCredential struct {
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

//RegistrationCredential is used for registration
type RegistrationCredential struct {
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//LoginResponse is the response after a successful login or registration
type LoginResponse struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"  validate:"required"`
	Token    string `json:"token" validate:"token"`
}
