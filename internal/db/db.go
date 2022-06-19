package db

import (
	"fmt"

	"github.com/vasjaj/todo/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func New(conf *config.Config) (*Database, error) {
	gormDB, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := gormDB.AutoMigrate(&Task{}, &User{}); err != nil {
		return nil, err
	}

	return &Database{gormDB}, nil
}
