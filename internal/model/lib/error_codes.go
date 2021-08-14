package lib

import (
	"net/http"
)

const (
	errInvalidStartDate        = "invalid startDate format. Allowed date format YYYY-MM-DD"
	errInvalidEndDate          = "invalid endDate format. Allowed date format YYYY-MM-DD"
	errBadRequest              = "invalid request body"
	errDateRange               = "startDate cannot be after endDate"
	errCountRange              = "maxCount value must be greater than or equal to minCount"
	errMongoDBFailure          = "something went wrong. Please try again later"
	errEmptyKey                = "key cannot be empty"
	errKeyNotFound             = "key does not exists"
	errFailedaddInMemoryRecord = "failed to add new record"
	errPageNotFound            = "404 page not found"
)

var ErrorMap map[int]*APIError

func InitErrorMap() {

	ErrorMap = make(map[int]*APIError)

	ErrorMap = map[int]*APIError{

		4000: {
			Message: errBadRequest,
			Status:  http.StatusBadRequest,
			Code:    4000,
		},
		4001: {
			Message: errInvalidStartDate,
			Status:  http.StatusBadRequest,
			Code:    4001,
		},
		4002: {
			Message: errInvalidEndDate,
			Status:  http.StatusBadRequest,
			Code:    4002,
		},
		4003: {
			Message: errDateRange,
			Status:  http.StatusBadRequest,
			Code:    4003,
		},
		4004: {
			Message: errCountRange,
			Status:  http.StatusBadRequest,
			Code:    4004,
		},
		4005: {
			Message: errBadRequest,
			Status:  http.StatusBadRequest,
			Code:    4005,
		},
		4006: {
			Message: errEmptyKey,
			Status:  http.StatusBadRequest,
			Code:    4006,
		},
		2000: {
			Message: errKeyNotFound,
			Status:  http.StatusOK,
			Code:    2000,
		},
		4040: {
			Message: errPageNotFound,
			Status:  http.StatusNotFound,
			Code:    4040,
		},
		4041: {
			Message: errPageNotFound,
			Status:  http.StatusNotFound,
			Code:    4041,
		},
		5000: {
			Message: errMongoDBFailure,
			Status:  http.StatusInternalServerError,
			Code:    5000,
		},
		5001: {
			Message: errFailedaddInMemoryRecord,
			Status:  http.StatusInternalServerError,
			Code:    5001,
		},
	}
}
