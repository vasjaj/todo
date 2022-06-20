package db

import (
	"fmt"
	"time"

	"github.com/vasjaj/todo/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	Ping() error
	Close() error

	GetUser(login string) (*User, error)
	CreateUser(login string, password string) (*User, error)

	GetTask(id int) (*Task, error)
	GetTasks(userID int) ([]Task, error)
	GetCompletedTasks(userID int) ([]Task, error)
	GetUncompletedTasks(userID int) ([]Task, error)
	CreateTask(userID int, title string, description string, dueDate *time.Time) error
	UpdateTask(taskID int, title string, description string, dueDate *time.Time) error
	DeleteTask(taskID int) error
	CompleteTask(taskID int) error
	UncompleteTask(taskID int) error
	GetTasksByLabel(labelID int) ([]Task, error)
	AddLabelToTask(labelID, taskID int) error
	RemoveLabelFromTask(labelID, taskID int) error

	GetLabel(id int) (*Label, error)
	GetLabels(userID int) ([]Label, error)
	CreateLabel(userID int, title string) error
	UpdateLabel(labelID int, title string) error
	DeleteLabel(labelID int) error

	GetComment(id int) (*Comment, error)
	GetComments(taskID int) ([]Comment, error)
	CreateComment(userID int, taskID int, title string, description string) error
	UpdateComment(id int, title string, description string) error
	DeleteComment(id int) error
}

type db struct {
	*gorm.DB
}

func New(conf *config.Config) (Database, error) {
	gormDB, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := gormDB.AutoMigrate(
		&Task{},
		&Comment{},
		&Label{},
		&LabelTask{},
		&User{},
	); err != nil {
		return nil, err
	}

	return &db{gormDB}, nil
}

func (d *db) Ping() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func (d *db) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
