package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IamNator/sendme_interview/utils/httperror"
	"github.com/IamNator/sendme_interview/utils/httpresp"
)

func TransactionHistoryHandler(transaction Transaction) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		userIDstr := r.FormValue("user_id")
		userID, er := strconv.Atoi(userIDstr)
		if er != nil {
			httperror.Default(fmt.Errorf("invalid user_id")).ReplyBadRequest(w)
			return
		}

		response, er := transaction.TransactionHistory(uint(userID))
		if er != nil {
			httperror.Default(er).ReplyUnprocessableEntity(w)
			return
		}

		httpresp.Default(response).ReplyCreated(w)
	}

}
