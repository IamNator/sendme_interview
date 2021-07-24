package router

import (
	"net/http"

	"github.com/IamNator/sendme_interview/handlers"
	"github.com/IamNator/sendme_interview/internal/services"
	"github.com/IamNator/sendme_interview/middleware"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Routes routes all http requests within the app
func Routes(db *gorm.DB) *mux.Router {

	mx := mux.NewRouter()
	userService := services.NewUser(db)
	transactionService := services.NewTransaction(db)

	//user services
	mx.HandleFunc("/register", handlers.RegistrationHandler(userService)).Methods(http.MethodPost) //get grade configuration
	mx.HandleFunc("/login", handlers.LoginHandler(userService)).Methods(http.MethodPost)           //get grade configuration

	tranx := mx.NewRoute().Subrouter()
	tranx.Use(middleware.ValidateToken(db))
	tranx.HandleFunc("/debit", handlers.DebitUserHandler(transactionService)).Methods(http.MethodPost)   //get grade configuration
	tranx.HandleFunc("/credit", handlers.CreditUserHandler(transactionService)).Methods(http.MethodPost) //get grade configuration

	tranx.HandleFunc("/transaction/history", handlers.TransactionHistoryHandler(transactionService)).Methods(http.MethodGet) //get grade configuration
	tranx.HandleFunc("/wallet", handlers.WalletHandler(transactionService)).Methods(http.MethodGet)                          //get grade configuration

	return mx
}
