package v1

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mocks "github.com/sakshi29/getir-app/internal/mocks"
	"github.com/sakshi29/getir-app/internal/model"
	mongoDB "github.com/sakshi29/getir-app/internal/repository/mongo-db"
)

func Test_useCaseV1_AddInMemoryRecord(t *testing.T) {

	m := make(map[string]string)
	m["test key"] = "test value"

	type fields struct {
		inMemoryStore map[string]string
	}
	type args struct {
		ctx    context.Context
		record model.InMemoryRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Case 1: Test Add Success",
			fields: fields{
				inMemoryStore: make(map[string]string),
			},
			args: args{
				ctx: context.Background(),
				record: model.InMemoryRecord{
					Key:   "test key",
					Value: "test value",
				},
			},
			wantErr: false,
		},
		{
			name: "Case 2: Duplicate record",
			fields: fields{
				inMemoryStore: m,
			},
			args: args{
				ctx: context.Background(),
				record: model.InMemoryRecord{
					Key:   "test key",
					Value: "test value",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucV1 := useCaseV1{
				inMemoryStore: tt.fields.inMemoryStore,
			}
			if err := ucV1.AddInMemoryRecord(tt.args.ctx, tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("useCaseV1.AddInMemoryRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCaseV1_GetInMemoryRecord(t *testing.T) {

	m := make(map[string]string)
	m["test key"] = "test value"

	type fields struct {
		inMemoryStore map[string]string
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.InMemoryRecord
		wantErr bool
	}{
		{
			name: "Case 1: Key exists",
			fields: fields{
				inMemoryStore: m,
			},
			args: args{
				ctx: context.Background(),
				key: "test key",
			},
			want: model.InMemoryRecord{
				Key:   "test key",
				Value: "test value",
			},
			wantErr: false,
		},
		{
			name: "Case 2: Key not found",
			fields: fields{
				inMemoryStore: make(map[string]string),
			},
			args: args{
				ctx: context.Background(),
				key: "test key",
			},
			want:    model.InMemoryRecord{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucV1 := useCaseV1{
				inMemoryStore: tt.fields.inMemoryStore,
			}
			got, err := ucV1.GetInMemoryRecord(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCaseV1.GetInMemoryRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCaseV1.GetInMemoryRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseV1_FetchDocuments(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepoMongoDB := mocks.NewMockMongoDBRepo(mockCtrl)

	date, err := time.Parse(model.TimeFormatYYYYMMDD, "2017-09-01")
	if err != nil {
		t.Fatal("failed tp parse date")
	}

	filteredDocuments := []model.Record{
		{
			Key:        "yetur",
			CreatedAt:  date,
			TotalCount: 2200,
		},
	}

	type fields struct {
		mongoDB mongoDB.MongoDBRepo
	}
	type args struct {
		ctx    context.Context
		record model.DocumentFilters
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		want                   []model.Record
		wantErr                bool
		mockFetchDocumentsFunc func(ctx context.Context, filter model.DocumentFilters)
	}{
		{
			name: "Case 1: MongoDB Error",
			fields: fields{
				mongoDB: mockRepoMongoDB,
			},
			args: args{
				ctx: context.Background(),
				record: model.DocumentFilters{
					MaxCount: 3000,
					MinCount: 2000,
				},
			},
			want:    nil,
			wantErr: true,
			mockFetchDocumentsFunc: func(ctx context.Context, filter model.DocumentFilters) {
				mockRepoMongoDB.EXPECT().FetchDocuments(ctx, filter).Return(nil, errors.New("Invalid query"))
			},
		},
		{
			name: "Case 2: MongoDB Success",
			fields: fields{
				mongoDB: mockRepoMongoDB,
			},
			args: args{
				ctx: context.Background(),
				record: model.DocumentFilters{
					MaxCount: 3000,
					MinCount: 2000,
				},
			},
			want:    filteredDocuments,
			wantErr: false,
			mockFetchDocumentsFunc: func(ctx context.Context, filter model.DocumentFilters) {
				mockRepoMongoDB.EXPECT().FetchDocuments(ctx, filter).Return(filteredDocuments, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucV1 := useCaseV1{
				mongoDB: tt.fields.mongoDB,
			}

			tt.mockFetchDocumentsFunc(context.Background(), tt.args.record)
			got, err := ucV1.FetchDocuments(tt.args.ctx, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCaseV1.FetchDocuments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCaseV1.FetchDocuments() = %v, want %v", got, tt.want)
			}
		})
	}
}
