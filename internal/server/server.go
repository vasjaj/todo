package server

import (
	"log"
	"net/http"

	"github.com/vasjaj/todo/internal/database"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

type Server interface {
	Run(addr string)
}

type server struct {
	*rpc.Server
}

func New(db database.Database) (Server, error) {
	srv := rpc.NewServer()
	srv.RegisterCodec(json2.NewCodec(), "application/json")

	if err := srv.RegisterService(&SeamlessService{db}, "service"); err != nil {
		return nil, err
	}

	return &server{srv}, nil
}

func (s *server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s))
}
