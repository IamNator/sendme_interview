package httperror

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/IamNator/sendme_interview/logger"
)

//Detail is a type
type Error = _error

//Detail ...
type _error struct {
	Code    int         `json:"code"`              //e.g 201, 200, 401
	Status  bool        `json:"status"`            //e.g false, true
	Name    string      `json:"name,omitempty"`    //name of the error
	Err     error       `json:"err,omitempty"`     //standard go error
	Message string      `json:"message,omitempty"` // any message for the error
	Detail  interface{} `json:"detail,omitempty"`  //custom error
}

//Default returns a new object
//
// status is false by default
func Default(err interface{}) _error {

	if err, ok := err.(*Error); ok {
		return _error{Status: false, Err: err.Err, Message: err.Err.Error(), Detail: err.Detail}
	}

	if err, ok := err.(error); ok {
		//Check for empty(nil) error
		var msg string
		if err != nil {
			msg = err.Error()
		} else {
			msg = ""
		}

		return _error{Status: false, Err: err, Message: msg, Detail: msg}
	}

	return _error{Status: false, Err: nil, Message: fmt.Sprintln(err), Detail: ""}
}

//New returns a new object
//
// status is false by default, Err can be nil
func New(code int, status bool, name string, Err error, msg string, Detail interface{}) _error {
	return _error{
		Code:    code,
		Status:  false,
		Name:    name,
		Err:     Err,
		Message: msg,
		Detail:  Detail}
}

//New2 creates a pointer to a httperror.Error object
func New2(code int, er error, detail interface{}) *_error {
	if detail == nil && er != nil {
		detail = er.Error()
	}
	return &_error{
		Code:    code,
		Status:  false,
		Name:    er.Error(),
		Err:     er,
		Detail:  detail,
		Message: er.Error(),
	}
}

/************************************SET FIELDS*********************************/

//SetStatus sets status field
func (e *_error) SetStatus(status bool) {
	e.Status = status
}

//SetCode sets code field
func (e *_error) SetCode(code int) {
	e.Code = code
}

//SetName sets name field
func (e *_error) SetName(name string) {
	e.Name = name
}

//SetErr sets err field
func (e *_error) SetErr(er error) {
	e.Err = er
}

//SetMessage sets message field
func (e *_error) SetMessage(message string) {
	e.Message = message
}

//SetError sets error field
func (e *_error) SetError(er interface{}) {
	e.Detail = er
}

/**********************************GET FIELDS***************************************/

//GetStatus gets status field
func (e _error) GetStatus() bool {
	return e.Status
}

//GetCode gets code field
func (e _error) GetCode() int {
	return e.Code
}

//GetName gets name field
func (e _error) GetName() string {
	return e.Name
}

//GetErr gets err field
func (e _error) GetErr() error {
	return e.Err
}

//GetMessage sets message field
func (e _error) GetMessage() string {
	return e.Message
}

//GetError gets error field
func (e _error) GetError() interface{} {
	return e.Detail
}

/******************************** HTTP REPLY***********************************/
func (e _error) ReplyInternalServerError(w http.ResponseWriter) {
	e.Code = http.StatusInternalServerError
	e.reply(w)
}

func (e _error) ReplyBadRequest(w http.ResponseWriter) {
	e.Code = http.StatusBadRequest
	e.reply(w)
}

//ReplyUnkwownResponse replies with status code 0
func (e _error) ReplyUnkwownResponse(w http.ResponseWriter) {
	e.Code = 0
	e.reply(w)
}

//ReplyUnauthorizedResponse replies with status code 0
func (e _error) ReplyUnauthorizedResponse(w http.ResponseWriter) {
	e.Code = http.StatusUnauthorized
	e.reply(w)
}

//ReplyUnprocessableEntity replies with status code 422 -  (http - StatusOK)
///
//this is the standard response for when a request is unsuccessful but it's not
//an error
func (e _error) ReplyUnprocessableEntity(w http.ResponseWriter) {
	e.Code = http.StatusUnprocessableEntity
	e.reply(w)
}

//Reply ...
func (e _error) Reply(w http.ResponseWriter) {
	if e.Code == 0 {
		e.Code = http.StatusUnprocessableEntity
	}
	logger.Warn.Println(e.Message, e.Detail, e.Err)
	if e.Code == 500 {
		e.Message = "Unexpected Internal Server Detail Occured"
	}
	e.reply(w)
}

func (e _error) reply(w http.ResponseWriter) {

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(e.Code)                                    //Write http code of error
	if er := json.NewEncoder(w).Encode(e.web()); er != nil { //encode error message
		http.Error(w, er.Error(), http.StatusInternalServerError)
	}
}

//Removes standard go error field
func (e _error) web() interface{} {

	return struct {
		Code    int         `json:"code"`              //e.g 201, 200, 401
		Status  bool        `json:"status"`            //e.g false, true
		Name    string      `json:"name,omitempty"`    //name of the error
		Message string      `json:"message,omitempty"` // any message for the error
		Detail  interface{} `json:"detail,omitempty"`  //custom error
	}{
		Code:    e.Code,
		Status:  e.Status,
		Name:    e.Name,
		Message: e.Message,
		Detail:  e.Detail,
	}
}

/**************** OTHERS ************************/

//Returns the error in string format
func (e _error) String() string {
	var status string
	if e.Status {
		status = "true"
	} else {
		status = "false"
	}
	return "code: " + strconv.Itoa(e.Code) + "\nstatus: " + status + "\nerror: " + e.Err.Error()
}

//Read reads the next len(p) bytes from the buffer
//or until the buffer is drained.
//The return value n is the number of bytes read. If the buffer has no data to return,
// err is io.EOF (unless len(p) is zero); otherwise it is nil.
func (e _error) Read(p []byte) (n int, err error) {
	b, err := json.Marshal(e)

	if err == nil {
		buf := bytes.NewBuffer(b)
		n, err = buf.Read(p)
	}

	return
}

//WriteToWriter writes to io.writer interface
func (e _error) WriteToWriter(w io.Writer) (n int, err error) {
	b, err := json.Marshal(e)

	if err == nil {
		n, err = w.Write(b)
	}

	return
}
