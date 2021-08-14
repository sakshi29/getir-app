package lib

import (
	"bytes"
	"fmt"
	"net/http"
)

type APIError struct {
	Message string `json:"msg"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
}

func (err *APIError) Error() string {

	var errorBuffer bytes.Buffer

	if err.Message != "" {
		errorBuffer.WriteString(fmt.Sprintf("Message: %s", err.Message))
	}

	errorBuffer.WriteString(fmt.Sprintf("\nAPI: %s", err.Message))

	if err.Code != 0 {
		errorBuffer.WriteString(fmt.Sprintf("\nCode: %d", err.Code))
	}

	return errorBuffer.String()
}

func NewAPIError(code int) *APIError {
	// Check error code in map
	if errMod, isExist := ErrorMap[code]; isExist {
		return &APIError{
			Message: errMod.Message,
			Code:    errMod.Code,
			Status:  errMod.Status,
		}
	}

	return &APIError{
		Message: "internal server error",
		Code:    -1,
		Status:  http.StatusInternalServerError,
	}
}
