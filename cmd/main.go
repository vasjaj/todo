package main

import (
	"os"

	"github.com/vasjaj/todo/internal/db"

	log "github.com/sirupsen/logrus"
	_ "github.com/vasjaj/todo/cmd/docs"
	"github.com/vasjaj/todo/internal/config"
	"github.com/vasjaj/todo/internal/server"
)

func main() {
	path := os.Args[1]
	log.Info("Config file path: ", path)

	conf, err := config.New(path)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	db, err := db.New(conf)
	if err != nil {
		log.Fatal("Failed to init database: ", err)
	}

	srv := server.New(db)
	srv.Run(conf.Server.Listen)
}
