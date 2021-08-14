package v1

import (
	"context"
	"errors"

	"github.com/sakshi29/getir-app/internal/model"
	"github.com/sakshi29/getir-app/internal/model/lib"
)

func (ucV1 useCaseV1) AddInMemoryRecord(ctx context.Context, record model.InMemoryRecord) error {

	if _, ok := ucV1.inMemoryStore[record.Key]; ok {
		return errors.New("key already exists")
	}

	ucV1.inMemoryStore[record.Key] = record.Value

	return nil
}

func (ucV1 useCaseV1) GetInMemoryRecord(ctx context.Context, key string) (model.InMemoryRecord, error) {

	record := model.InMemoryRecord{}

	if val, ok := ucV1.inMemoryStore[key]; ok {
		record.Key = key
		record.Value = val
		return record, nil
	}

	return record, lib.NewAPIError(2000)
}

func (ucV1 useCaseV1) FetchDocuments(ctx context.Context, record model.DocumentFilters) ([]model.Record, error) {

	var documents []model.Record
	var err error

	documents, err = ucV1.mongoDB.FetchDocuments(ctx, record)
	if err != nil {
		return nil, err
	}

	return documents, nil
}
