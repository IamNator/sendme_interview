package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IamNator/sendme_interview/models"
	"github.com/IamNator/sendme_interview/utils/httperror"
	"github.com/IamNator/sendme_interview/utils/httpresp"
)

func LoginHandler(user User) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var userRequestData models.LoginCredential
		if er := json.NewDecoder(r.Body).Decode(&userRequestData); er != nil {
			httperror.Default(fmt.Errorf("unable to read request body, \n %v", er.Error())).ReplyBadRequest(w)
			return
		}

		if _, er := Validate(userRequestData, w); er != nil {
			return
		}

		response, er := user.LoginUser(userRequestData)
		if er != nil {
			httperror.Default(er).ReplyUnprocessableEntity(w)
			return
		}

		httpresp.Default(response).ReplyCreated(w)
	}

}
