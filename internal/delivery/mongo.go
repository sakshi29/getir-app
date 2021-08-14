package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sakshi29/getir-app/internal/model"
	constant "github.com/sakshi29/getir-app/internal/model"
	"github.com/sakshi29/getir-app/internal/model/lib"
)

func (a *API) GetDocuments(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	ctx := r.Context()

	if r.Method != http.MethodPost {
		return nil, lib.NewAPIError(4040)
	}

	var record model.DocumentFilters

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&record)
	if err != nil {
		log.Printf("[Delivery] [GetDocuments] [ERR] Failed to decode the request body: %s", err.Error())
		return nil, lib.NewAPIError(4000)
	}

	record.StartDate, err = time.Parse(constant.TimeFormatYYYYMMDD, record.StartDateStr)
	if err != nil {
		log.Printf("[Delivery] [GetDocuments] [ERR] Filed to parse start date: %s", record.StartDateStr)
		return nil, lib.NewAPIError(4001)
	}

	record.EndDate, err = time.Parse(constant.TimeFormatYYYYMMDD, record.EndDateStr)
	if err != nil {
		log.Printf("[Delivery] [GetDocuments] [ERR] Filed to parse end date: %s", record.EndDateStr)
		return nil, lib.NewAPIError(4002)
	}

	if record.EndDate.Before(record.StartDate) {
		log.Printf("[Delivery] [GetDocuments] [ERR] End date: %s should not be before start date: %s", record.EndDateStr, record.StartDateStr)
		return nil, lib.NewAPIError(4003)
	}

	if record.MaxCount < record.MinCount {
		log.Printf("[Delivery] [GetDocuments] [ERR] Maxcount: %d should not be less than Mincount: %d", record.MaxCount, record.MinCount)
		return nil, lib.NewAPIError(4004)
	}

	filteredDocuments, err := a.useCaseV1.FetchDocuments(ctx, record)
	if err != nil {
		log.Printf("[Delivery] [GetDocuments] [ERR] Failed to get documents with error: %s", err.Error())
		return nil, lib.NewAPIError(5000)
	}

	return filteredDocuments, nil
}
