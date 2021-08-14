package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sakshi29/getir-app/internal/model"
	"github.com/sakshi29/getir-app/internal/model/lib"
)

func (a *API) HandleRecord(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	switch r.Method {
	case http.MethodGet:
		return a.getInMemoryRecord(w, r)
	case http.MethodPost:
		return a.addInMemoryRecord(w, r)
	default:
		return nil, lib.NewAPIError(4041)
	}
}

func (a *API) addInMemoryRecord(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	var record model.InMemoryRecord

	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&record)
	if err != nil {
		log.Printf("[Delivery] [addInMemoryRecord] [ERR] Failed to decode the request body: %s", err.Error())
		return nil, lib.NewAPIError(4005)
	}

	err = a.useCaseV1.AddInMemoryRecord(ctx, record)
	if err != nil {
		log.Printf("[Delivery] [addInMemoryRecord] [ERR] Failed to add record  err: %s", err.Error())
		return nil, lib.NewAPIError(5001)
	}

	return record, nil
}

func (a *API) getInMemoryRecord(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()

	key := strings.TrimSpace(r.URL.Query().Get("key"))

	if key == "" {
		log.Printf("[Delivery] [getInMemoryRecord] [ERR] Empty key found")
		return nil, lib.NewAPIError(4006)
	}

	record, err := a.useCaseV1.GetInMemoryRecord(ctx, key)
	if err != nil {
		log.Printf("[Delivery] [getInMemoryRecord] [ERR] Failed to get record for key: %s, err: %s", key, err.Error())
		return nil, err
	}

	return record, nil
}
