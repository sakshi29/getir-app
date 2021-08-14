package delivery

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mocks "github.com/sakshi29/getir-app/internal/mocks"
	"github.com/sakshi29/getir-app/internal/model"
	"github.com/sakshi29/getir-app/internal/model/lib"
	"github.com/sakshi29/getir-app/internal/usecase"
)

func generateRequest(url, method string, body []byte) *http.Request {
	var req *http.Request

	if method == "POST" {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(body))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	return req
}

func TestAPI_getInMemoryRecord(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCaseV1 := mocks.NewMockUseCaseV1(mockCtrl)

	startDate, err := time.Parse(model.TimeFormatYYYYMMDD, "2015-05-10")
	if err != nil {
		t.Fatal("failed to parse date")
	}

	endDate, err := time.Parse(model.TimeFormatYYYYMMDD, "2015-05-11")
	if err != nil {
		t.Fatal("failed to parse date")
	}

	lib.InitErrorMap()

	type fields struct {
		useCaseV1 usecase.UseCaseV1
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "Case 1: Test Success",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/documents", "POST", []byte(`{"startDate": "2015-05-10", "endDate": "2015-05-11", "minCount": 3400, "maxCount": 3500}`)),
			},
			want:    make([]model.Record, 0),
			wantErr: false,
			mock: func() {
				mockUseCaseV1.EXPECT().FetchDocuments(gomock.Any(), model.DocumentFilters{
					StartDateStr: "2015-05-10",
					EndDateStr:   "2015-05-11",
					StartDate:    startDate,
					EndDate:      endDate,
					MinCount:     3400,
					MaxCount:     3500,
				}).Return(make([]model.Record, 0), nil)
			},
		},
		{
			name: "Case 2: Test Failure",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/documents", "POST", []byte(`{"startDate": "2015-05-10", "endDate": "2015-05-11", "minCount": 3400, "maxCount": 3500}`)),
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockUseCaseV1.EXPECT().FetchDocuments(gomock.Any(), model.DocumentFilters{
					StartDateStr: "2015-05-10",
					EndDateStr:   "2015-05-11",
					StartDate:    startDate,
					EndDate:      endDate,
					MinCount:     3400,
					MaxCount:     3500,
				}).Return(nil, errors.New("failed to get records"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				useCaseV1: tt.fields.useCaseV1,
			}

			tt.mock()

			got, err := a.GetDocuments(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.GetDocuments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.GetDocuments() = %v, want %v", got, tt.want)
			}
		})
	}
}
