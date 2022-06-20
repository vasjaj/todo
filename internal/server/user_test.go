package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/vasjaj/todo/internal/db"
	mock_db "github.com/vasjaj/todo/internal/db/mock"
)

//var conf = config.Config{
//	Server: config.Server{
//		Listen: ":80",
//		JWT: struct {
//			Secret     string `yaml:"secret" validate:"required"`
//			SecondsTTL int    `yaml:"seconds_ttl" validate:"required"`
//		}{
//			Secret:     "secret",
//			SecondsTTL: 3600,
//		},
//	},
//	Database: config.Database{
//		User:     "root",
//		Password: "root",
//		Host:     "localhost",
//		Port:     3306,
//		Name:     "todo",
//	},
//}

//func TestServer_login(t *testing.T) {
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, "/login",
//		strings.NewReader(`{"login":"admin","password":"admin123"}`))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	h := New(mock_db.NewMockDatabase(ctrl), &conf)
//
//	//if assert.NoError(t, h.login(c)) {
//	//	assert.Equal(t, http.StatusOK, rec.Code)
//	//	assert.Equal(t, `{"token":"ok"}`, rec.Body.String())
//	//}
//}

func TestServer_login(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login":"admin","password":"admin123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, httptest.NewRecorder())

	invalidReq := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login":"admin","password":""}`))
	invalidReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	invalidC := e.NewContext(invalidReq, httptest.NewRecorder())

	incorrectReq := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"login":"admin","password":"admin123"}`))
	incorrectReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	incorrectC := e.NewContext(incorrectReq, httptest.NewRecorder())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_db.NewMockDatabase(ctrl)
	// h := New(mock_db.NewMockDatabase(ctrl), &conf)

	type fields struct {
		Echo      *echo.Echo
		db        db.Database
		jwtConfig jwtConfig
		listen    string
	}
	type args struct {
		c        echo.Context
		mockFunc func()
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"login",
			fields{
				Echo:      e,
				db:        mockDB,
				jwtConfig: jwtConfig{Secret: "some-secret", SecondsTTL: 3600},
				listen:    ":80",
			},
			args{c, func() {
				passwordHash, err := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
				assert.NoError(t, err)

				mockDB.EXPECT().GetUser("admin").Return(&db.User{Login: "admin", PasswordHash: string(passwordHash)}, nil)
			}},
			false,
		},
		{
			"incorrect input",
			fields{
				Echo:      e,
				db:        mockDB,
				jwtConfig: jwtConfig{Secret: "some-secret", SecondsTTL: 3600},
				listen:    ":80",
			},
			args{invalidC, func() {}},
			true,
		},
		{
			"incorrect password",
			fields{
				Echo:      e,
				db:        mockDB,
				jwtConfig: jwtConfig{Secret: "some-secret", SecondsTTL: 3600},
				listen:    ":80",
			},
			args{incorrectC, func() {
				passwordHash, err := bcrypt.GenerateFromPassword([]byte("other_password"), 14)
				assert.NoError(t, err)

				mockDB.EXPECT().GetUser("admin").Return(&db.User{Login: "admin", PasswordHash: string(passwordHash)}, nil)
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.mockFunc()

			s := &Server{
				Echo:      tt.fields.Echo,
				db:        tt.fields.db,
				jwtConfig: tt.fields.jwtConfig,
				listen:    tt.fields.listen,
			}

			if err := s.login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
