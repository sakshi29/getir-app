package delivery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/sakshi29/getir-app/internal/mocks"
	"github.com/sakshi29/getir-app/internal/model"
	"github.com/sakshi29/getir-app/internal/model/lib"
	"github.com/sakshi29/getir-app/internal/usecase"
)

func TestAPI_HandleRecord(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUseCaseV1 := mocks.NewMockUseCaseV1(mockCtrl)

	lib.InitErrorMap()

	record := model.InMemoryRecord{
		Key:   "active-tabs",
		Value: "getir",
	}

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
			name: "Case 1: Test POST Success",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/in-memory", "POST", []byte(`{
					"key":"active-tabs",
					"value":"getir"
				}`)),
			},
			want:    record,
			wantErr: false,
			mock: func() {
				mockUseCaseV1.EXPECT().AddInMemoryRecord(gomock.Any(), record).Return(nil)
			},
		},
		{
			name: "Case 2: Test POST Failure",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/in-memory", "POST", []byte(`{
					"key":"active-tabs",
					"value":"getir"
				}`)),
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockUseCaseV1.EXPECT().AddInMemoryRecord(gomock.Any(), record).Return(errors.New("failed to add record"))
			},
		},
		{
			name: "Case 3: Invalid POST body",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/in-memory", "POST", []byte(`{
					"key": 7878,
					"value":"getir"
				}`)),
			},
			want:    nil,
			wantErr: true,
			mock:    func() {},
		},
		{
			name: "Case 4: Invalid GET key",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/in-memory?key=", "GET", nil),
			},
			want:    nil,
			wantErr: true,
			mock:    func() {},
		},
		{
			name: "Case 5: Test GET Success",
			fields: fields{
				useCaseV1: mockUseCaseV1,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: generateRequest("/in-memory?key=test-key", "GET", nil),
			},
			want:    record,
			wantErr: false,
			mock: func() {
				mockUseCaseV1.EXPECT().GetInMemoryRecord(gomock.Any(), "test-key").Return(record, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				useCaseV1: tt.fields.useCaseV1,
			}

			tt.mock()

			got, err := a.HandleRecord(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.HandleRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.HandleRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
