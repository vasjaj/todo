package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vasjaj/todo/internal/config"
	"github.com/vasjaj/todo/internal/database"
	"github.com/vasjaj/todo/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("config path not provided")
	}

	conf, err := config.New(os.Args[1])
	if err != nil {
		log.Fatal("config error: ", err)
	}

	log.Println("config loaded: ", conf)

	db, err := database.New(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name))
	if err != nil {
		log.Fatal("database error: ", err)
	}
	defer func(db database.Database) {
		err := db.Close()
		if err != nil {
			log.Println("database close error: ", err)
		}
	}(db)

	log.Println("database connected")

	srv, err := server.New(db)
	if err != nil {
		log.Fatal("server error: ", err)
	}

	log.Println("starting server")

	srv.Run(conf.Server.Listen)
}
