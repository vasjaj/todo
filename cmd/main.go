package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	_ "github.com/vasjaj/todo/docs"
	"github.com/vasjaj/todo/internal/config"
	"github.com/vasjaj/todo/internal/db"
	"github.com/vasjaj/todo/internal/server"
)

func main() {
	log.SetLevel(log.InfoLevel)

	path := os.Args[1]
	log.Info("Config file path: ", path)

	conf, err := config.New(path)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	database, err := db.New(conf)
	if err != nil {
		log.Fatal("Failed to init database: ", err)
	}

	log.Error("Failed to run server: ", server.New(database, conf).Run())

	if err := database.Close(); err != nil {
		log.Error(err)
	}
}
