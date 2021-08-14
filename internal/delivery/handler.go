package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	constant "github.com/sakshi29/getir-app/internal/model"
	"github.com/sakshi29/getir-app/internal/model/lib"
)

// Each handler can return the data and error, and ServeHTTP can chose how to convert this
type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	response := Response{}
	var cancelFn func()
	var data interface{}
	var err error
	var buffer []byte

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	ctx := r.Context()

	// Set a timeout of 15 seconds on the responses.
	ctx, cancelFn = context.WithTimeout(ctx, 15*time.Second)
	defer cancelFn()

	r = r.WithContext(ctx)

	errStatus := http.StatusInternalServerError

	data, err = fn(w, r)

	w.Header().Set("Content-Type", "application/json")

	if data != nil && err == nil {
		response.Data = data
		response.Message = constant.ApiOkResponse
		if buffer, err = json.Marshal(response); err == nil {
			_, err := w.Write(buffer)
			if err != nil {
				log.Printf("[API] [ServeHTTP] [ERR] error in writing http response, path: %s, error: %s\n", r.URL.Path, err.Error())
				return
			}
		}
	}

	if err != nil {
		errStatus = processError(err, &response, errStatus)
		log.Printf("[API] [ServeHTTP] [ERR] api: %s failed with status: %d with error: %s\n", r.URL.Path, errStatus, err.Error())
		w.WriteHeader(errStatus)
	} else {
		return
	}

	buffer, _ = json.Marshal(response)
	_, err = w.Write(buffer)
	if err != nil {
		log.Printf("[API] [ServeHTTP] [ERR] error in writing http for api: %s response error: %s\n", r.URL.Path, err.Error())
	}
}

func processError(err error, response *Response, errStatus int) int {

	switch t := err.(type) {

	case *lib.APIError:
		response.Code = t.Code
		response.Message = t.Message

		return t.Status

	case net.Error:
		if t.Timeout() {
			response.Message = errors.New(constant.ApiTimeoutResponse).Error()
		}
	}

	return errStatus
}
