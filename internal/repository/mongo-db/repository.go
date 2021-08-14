package mangoDB

import (
	"context"

	"github.com/sakshi29/getir-app/internal/model"
)

// Interface to manage mongo DB related functionalities
type MongoDBRepo interface {
	FetchDocuments(ctx context.Context, filters model.DocumentFilters) ([]model.Record, error)
}
