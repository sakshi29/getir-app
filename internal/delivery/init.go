package delivery

import (
	"net/http"

	"github.com/sakshi29/getir-app/common/config"
	"github.com/sakshi29/getir-app/internal/usecase"
)

func Init(cfg *config.Config, useCaseV1 usecase.UseCaseV1) {

	api := &API{
		cfg:       cfg,
		useCaseV1: useCaseV1,
	}

	api.InitHandlers()
}

func (delivery *API) InitHandlers() {

	// Healthcheck api
	http.Handle("/healthcheck", HandlerFunc(delivery.Health))

	// Fetch documents from Mongo DB collection
	http.Handle("/documents", HandlerFunc(delivery.GetDocuments))

	// Add or fetch a record from in-memory DB
	http.Handle("/in-memory", HandlerFunc(delivery.HandleRecord))
}
