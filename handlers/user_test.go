package handlers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IamNator/sendme_interview/handlers"
	"github.com/IamNator/sendme_interview/internal/schema"
	"github.com/IamNator/sendme_interview/models"
	"github.com/IamNator/sendme_interview/utils/httprequest"
)

var store = make(map[string]schema.User)
var usr = user{DB: store}

func Test_RegisterHandler(t *testing.T) {

	for _, tt := range testLogin {
		t.Run(tt.Name, func(t *testing.T) {

			//forming the body of the request
			body, er := httprequest.NewReader(tt.User.RegistrationRequestBody)
			if er != nil {
				t.Fatalf(er.Error())
			}

			//setting up the reesponse recorder and request object
			rr := httptest.NewRecorder()
			req, er := http.NewRequest(http.MethodPost, "/", body)
			if er != nil {
				t.Fatalf(er.Error())
			}

			handler := http.HandlerFunc(handlers.RegistrationHandler(usr))
			handler.ServeHTTP(rr, req)

			if rr.Code != 201 {
				by, _ := ioutil.ReadAll(rr.Body)
				t.Fatal(string(by))
			}

			var responseToUser struct {
				Status  bool                 `json:"status,omitempty"`
				Code    int                  `json:"code,omitempty"`
				Name    string               `json:"name,omitempty"` //name of the error
				Message string               `json:"message,omitempty"`
				Error   interface{}          `json:"error,omitempty"` //for errors that occur even if request is successful
				Data    models.LoginResponse `json:"data,omitempty"`
			}

			json.NewDecoder(rr.Body).Decode(&responseToUser)

			assertEqual(t, tt.User.Want, responseToUser.Data, tt.Name)

		})
	}

}

func Test_LoginHandler(t *testing.T) {

	for _, tt := range testLogin {
		t.Run(tt.Name, func(t *testing.T) {

			//forming the body of the request
			body, er := httprequest.NewReader(tt.User.LoginRequestBody)
			if er != nil {
				t.Fatalf(er.Error())
			}

			//setting up the reesponse recorder and request object
			rr := httptest.NewRecorder()
			req, er := http.NewRequest(http.MethodPost, "/", body)
			if er != nil {
				t.Fatalf(er.Error())
			}

			handler := http.HandlerFunc(handlers.LoginHandler(usr))
			handler.ServeHTTP(rr, req)

			if rr.Code != 201 {
				by, _ := ioutil.ReadAll(rr.Body)
				t.Fatal(string(by))
			}

			var responseToUser struct {
				Status  bool                 `json:"status,omitempty"`
				Code    int                  `json:"code,omitempty"`
				Name    string               `json:"name,omitempty"` //name of the error
				Message string               `json:"message,omitempty"`
				Error   interface{}          `json:"error,omitempty"` //for errors that occur even if request is successful
				Data    models.LoginResponse `json:"data,omitempty"`
			}

			json.NewDecoder(rr.Body).Decode(&responseToUser)

			assertEqual(t, tt.User.Want, responseToUser.Data, tt.Name)

		})
	}

}
