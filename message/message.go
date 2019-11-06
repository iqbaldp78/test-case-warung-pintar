//Package message used for control and display error or succes message
package message

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var messages = map[int]map[string]interface{}{
	0: {"message": "Page not found", "status": http.StatusNotFound},
	1: {"message": "You are not authorize", "status": http.StatusUnauthorized},
	2: {"message": "Forbidden endpoint", "status": http.StatusForbidden},
	3: {"message": "Data not found", "status": http.StatusNotFound},
	4: {"message": "Wrong parameter", "status": http.StatusBadRequest},
	5: {"message": "Data on field '{f}' not permitted", "status": http.StatusBadRequest},
	6: {"message": "{f} failed", "status": http.StatusBadRequest},
	7: {"message": "{f} success", "status": http.StatusOK},
	8: {"message": "Internal Server Error", "status": http.StatusInternalServerError},
}

//Error used for get message for display
type Error struct {
	Message    string `json:"message,omitempty"`
	Code       string `json:"code,omitempty"`
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"-"`
}

//New used for create 'Error' structure
func New(code int, err interface{}, fields ...string) Error {
	message := make(map[string]interface{})
	for k, v := range messages[code] {
		message[k] = v
	}

	for _, field := range fields {
		message["message"] = strings.Replace(message["message"].(string), "{f}", field, 1)
	}

	if err == nil {
		err = errors.New("")
	}

	return Error{
		StatusCode: message["status"].(int),
		Message:    message["message"].(string),
		Code:       fmt.Sprintf("%03d", code),
		Error:      err.(error).Error(),
	}
}

//IsInitial used for checking object 'Error' is empty
func (obj Error) IsInitial() bool {
	if reflect.DeepEqual(obj, Error{}) {
		return true
	}
	return false
}
