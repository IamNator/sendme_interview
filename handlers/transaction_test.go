package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/IamNator/sendme_interview/handlers"
	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/models"
	"github.com/IamNator/sendme_interview/utils/httprequest"
)

var wad = map[int]schema.Wallet{
	1: {
		ID:      1,
		UserID:  1,
		Balance: 4000,
	},
	2: {
		ID:      2,
		UserID:  2,
		Balance: 8000,
	},
}

var testTransaction = testTranxDB{
	DB: transaction{
		DB: struct {
			Wallet      map[int]schema.Wallet //userID ->  wallet
			Transaction map[int][]schema.Transaction
		}{
			Wallet:      wad,
			Transaction: make(map[int][]schema.Transaction),
		},
	},

	Tests: []struct {
		Name              string
		UserID            int
		DebitRequestBody  models.DebitUser
		CreditRequestBody models.CreditUser
	}{

		{
			Name:   "first test",
			UserID: 1,
			DebitRequestBody: models.DebitUser{
				UserID: 1,
				Amount: 1000,
			},

			CreditRequestBody: models.CreditUser{
				UserID: 1,
				Amount: 2000,
			},
		},
	},
}

func Test_CreditUserHandler(t *testing.T) {

	for _, tt := range testTransaction.Tests {
		t.Run(tt.Name, func(t *testing.T) {

			//forming the body of the request
			body, er := httprequest.NewReader(tt.CreditRequestBody)
			if er != nil {
				t.Fatalf(er.Error())
			}

			//setting up the reesponse recorder and request object
			rr := httptest.NewRecorder()
			req, er := http.NewRequest(http.MethodPost, "/", body)
			if er != nil {
				t.Fatalf(er.Error())
			}

			handler := http.HandlerFunc(handlers.CreditUserHandler(&testTransaction.DB))
			handler.ServeHTTP(rr, req)

			if rr.Code != 201 {
				by, _ := ioutil.ReadAll(rr.Body)
				t.Fatal(string(by))
			}

			var responseToUser struct {
				Status  bool          `json:"status,omitempty"`
				Code    int           `json:"code,omitempty"`
				Name    string        `json:"name,omitempty"` //name of the error
				Message string        `json:"message,omitempty"`
				Error   interface{}   `json:"error,omitempty"` //for errors that occur even if request is successful
				Data    schema.Wallet `json:"data,omitempty"`
			}

			json.NewDecoder(rr.Body).Decode(&responseToUser)

			assertEqual(t, testTransaction.DB.DB.Wallet[tt.CreditRequestBody.UserID], responseToUser.Data, tt.Name)

		})
	}
}

func Test_DebitUserHandler(t *testing.T) {

	for _, tt := range testTransaction.Tests {
		t.Run(tt.Name, func(t *testing.T) {

			//forming the body of the request
			body, er := httprequest.NewReader(tt.DebitRequestBody)
			if er != nil {
				t.Fatalf(er.Error())
			}

			//setting up the reesponse recorder and request object
			rr := httptest.NewRecorder()
			req, er := http.NewRequest(http.MethodPost, "/", body)
			if er != nil {
				t.Fatalf(er.Error())
			}

			handler := http.HandlerFunc(handlers.DebitUserHandler(&testTransaction.DB))
			handler.ServeHTTP(rr, req)

			if rr.Code != 201 {
				by, _ := ioutil.ReadAll(rr.Body)
				t.Fatal(string(by))
			}

			var responseToUser struct {
				Status  bool          `json:"status,omitempty"`
				Code    int           `json:"code,omitempty"`
				Name    string        `json:"name,omitempty"` //name of the error
				Message string        `json:"message,omitempty"`
				Error   interface{}   `json:"error,omitempty"` //for errors that occur even if request is successful
				Data    schema.Wallet `json:"data,omitempty"`
			}

			json.NewDecoder(rr.Body).Decode(&responseToUser)

			if bal := (testTransaction.DB.DB.Wallet[tt.DebitRequestBody.UserID].Balance - tt.DebitRequestBody.Amount); bal < 0 {
				if rr.Code == http.StatusUnprocessableEntity {
					t.Log(responseToUser.Message)
				}
			}

			assertEqual(t, testTransaction.DB.DB.Wallet[tt.DebitRequestBody.UserID], responseToUser.Data, tt.Name)

		})
	}

}

func Test_WalletHandler(t *testing.T) {

	for _, tt := range testTransaction.Tests {
		t.Run(tt.Name, func(t *testing.T) {

			//setting up the reesponse recorder and request object
			rr := httptest.NewRecorder()
			req, er := http.NewRequest(http.MethodGet, "/", nil)
			if er != nil {
				t.Fatalf(er.Error())
			}

			q := req.URL.Query()

			q.Add("user_id", strconv.Itoa(tt.UserID))
			req.URL.RawQuery = q.Encode()

			handler := http.HandlerFunc(handlers.WalletHandler(&testTransaction.DB))
			handler.ServeHTTP(rr, req)

			if rr.Code != 201 {
				by, _ := ioutil.ReadAll(rr.Body)
				t.Fatal(string(by))
			}

			var responseToUser struct {
				Status  bool          `json:"status,omitempty"`
				Code    int           `json:"code,omitempty"`
				Name    string        `json:"name,omitempty"` //name of the error
				Message string        `json:"message,omitempty"`
				Error   interface{}   `json:"error,omitempty"` //for errors that occur even if request is successful
				Data    schema.Wallet `json:"data,omitempty"`
			}

			json.NewDecoder(rr.Body).Decode(&responseToUser)

			assertEqual(t, testTransaction.DB.DB.Wallet[tt.UserID], responseToUser.Data, tt.Name)

		})
	}

}
