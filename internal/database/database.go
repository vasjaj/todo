package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	Close() error
	RegisterTransaction(callerID, withdraw, deposit int, playerName, currency, transactionRef string) (Transaction, int, error)
	GetBalance(playerName, currency string) (int, error)
	RollbackTransaction(callerId int, playerName, transactionRef string) error
}

type database struct {
	*sql.DB
}

func New(conn string) (Database, error) {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return &database{db}, nil
}

func (d *database) Close() error {
	return d.DB.Close()
}
