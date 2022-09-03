package server

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "github.com/vasjaj/todo/internal/database/mock"

	"github.com/vasjaj/todo/internal/database"
)

var errDatabase = errors.New("database error")

func TestSeamlessService_GetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDatabase(ctrl)

	type fields struct {
		db database.Database
	}
	type args struct {
		in0      *http.Request
		req      *GetBalanceRequest
		res      *GetBalanceResponse
		mockFunc func()
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"invalid request", fields{db: mockDB}, args{in0: nil, req: &GetBalanceRequest{
			Currency: "DUBLON",
		}, res: &GetBalanceResponse{}, mockFunc: func() {
			mockDB.EXPECT().GetBalance(gomock.Any(), gomock.Any()).Times(0)
		}}, true},
		{"database error", fields{db: mockDB}, args{in0: nil, req: &GetBalanceRequest{
			CallerId:   1,
			PlayerName: "John",
			Currency:   "USD",
		}, res: &GetBalanceResponse{}, mockFunc: func() {
			mockDB.EXPECT().GetBalance("John", "USD").Return(0, errDatabase).Times(1)
		}}, true},
		{"ok", fields{db: mockDB}, args{in0: nil, req: &GetBalanceRequest{
			CallerId:   1,
			PlayerName: "John",
			Currency:   "USD",
		}, res: &GetBalanceResponse{}, mockFunc: func() {
			mockDB.EXPECT().GetBalance("John", "USD").Return(1, nil).Times(1)
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.mockFunc()

			s := &SeamlessService{
				db: tt.fields.db,
			}
			if err := s.GetBalance(tt.args.in0, tt.args.req, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
