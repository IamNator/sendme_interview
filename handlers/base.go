package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/logger"
	"github.com/IamNator/sendme_interview/models"
	"github.com/IamNator/sendme_interview/utils/httperror"
	"github.com/go-playground/validator"
)

type User interface {
	RegisterNewUser(credentials models.RegistrationCredential) (*models.LoginResponse, error)
	LoginUser(credentials models.LoginCredential) (*models.LoginResponse, error)
	//LogOut()
}

type Transaction interface {
	DebitUser(debit models.DebitUser) (*schema.Wallet, error)
	CreditUser(credit models.CreditUser) (*schema.Wallet, error)
	TransactionHistory(userID uint) ([]*schema.Transaction, error)
	WalletBalance(userID uint) (*schema.Wallet, error)
}

const (
	//InternalServerError is error for internal server error
	InternalServerError = "Unexpected Internal Server Error occured"
)

var v *validator.Validate

func init() {

	v = validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	logger.Logger.Println("validator object created")

}

//Validate validates user input
func Validate(input interface{}, w http.ResponseWriter) (map[string]interface{}, error) {

	details := make(map[string]interface{})

	er := v.Struct(input)
	if er == nil {
		return nil, nil
	}

	//type assertion for validation errors
	validationErrors, ok := er.(validator.ValidationErrors)
	if !ok {
		return nil, er
	}

	for _, err := range validationErrors {
		details[err.Field()] = fmt.Sprintf("%v is %v", err.Field(), err.Tag())
	}

	if w == nil {
		return details, er
	}

	httperror.New(400, false, "validation error", nil, "validation error", details).
		ReplyBadRequest(w)

	return details, er

}
