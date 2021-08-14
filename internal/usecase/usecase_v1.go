package usecase

import (
	"context"

	"github.com/sakshi29/getir-app/internal/model"
)

//UseCaseV1 defined API Version 1 bussiness logic
type UseCaseV1 interface {
	AddInMemoryRecord(ctx context.Context, record model.InMemoryRecord) error
	GetInMemoryRecord(ctx context.Context, key string) (model.InMemoryRecord, error)
	FetchDocuments(ctx context.Context, record model.DocumentFilters) ([]model.Record, error)
}
