package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type Transaction struct {
	Id         int
	UserId     int
	CallerID   int
	Reference  string
	Withdraw   int
	Deposit    int
	Currency   string
	CreatedAt  time.Time
	RevertedAt *time.Time
}

func (d *database) GetTransaction(playerName, transactionRef string) (Transaction, error) {
	var t Transaction
	err := d.QueryRow(getTransactionQuery, playerName, transactionRef).Scan(
		&t.Id, &t.UserId, &t.CallerID, &t.Reference, &t.Withdraw, &t.Deposit, &t.Currency, &t.CreatedAt, &t.RevertedAt,
	)
	return t, err
}

func (d *database) GetBalance(playerName, currency string) (int, error) {
	var balance int
	err := d.QueryRow(getBalanceQuery, playerName, currency).Scan(&balance)
	return balance, err
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (d *database) RegisterTransaction(
	callerID, withdraw, deposit int, playerName, currency, transactionRef string,
) (Transaction, int, error) {
	tx, err := d.BeginTx(context.TODO(), nil)
	if err != nil {
		return Transaction{}, 0, err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
		}
	}(tx)

	var balance int
	err = tx.QueryRow(getBalanceQuery, playerName, currency).Scan(&balance)
	if err != nil {
		return Transaction{}, 0, err
	}

	if balance < withdraw {
		return Transaction{}, 0, ErrInsufficientFunds
	}

	if _, err = tx.Exec(registerTransactionQuery, playerName, callerID, transactionRef, withdraw, deposit, currency); err != nil {
		return Transaction{}, 0, err
	}

	var t Transaction
	if err = tx.QueryRow(getTransactionQuery, playerName, transactionRef).Scan(
		&t.Id, &t.UserId, &t.CallerID, &t.Reference, &t.Withdraw, &t.Deposit, &t.Currency, &t.CreatedAt, &t.RevertedAt,
	); err != nil {
		return Transaction{}, 0, err
	}

	if err = tx.Commit(); err != nil {
		return Transaction{}, 0, err
	}

	return t, balance - withdraw + deposit, err
}

func (d *database) RollbackTransaction(callerId int, playerName, transactionRef string) error {
	tx, err := d.BeginTx(context.TODO(), nil)
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
		}
	}(tx)

	var exists bool
	err = tx.QueryRow(existsTransactionQuery, playerName, transactionRef).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if exists {
		if _, err = tx.Exec(rollbackTransactionQuery, playerName, transactionRef); err != nil {
			return err
		}
	} else {
		if _, err = tx.Exec(registerTransactionQuery, playerName, callerId, transactionRef, 0, 0); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

const (
	existsTransactionQuery = `
	SELECT EXISTS(SELECT 1 FROM transactions WHERE user_id = (SELECT id FROM users WHERE name = ?) AND reference = ?)`

	getTransactionQuery = `
	SELECT id, user_id, caller_id, reference, withdraw, deposit, currency, created_at, reverted_at
	FROM transactions
	WHERE user_id = (SELECT id FROM users WHERE name = ?) AND reference = ?`

	getBalanceQuery = `
	SELECT SUM(deposit - withdraw) FROM transactions
	WHERE user_id = (SELECT id FROM users WHERE name = ?)
	AND currency = ?
	AND reverted_at IS NULL`

	registerTransactionQuery = `
	INSERT INTO transactions (user_id, caller_id, reference, withdraw, deposit, currency)
	VALUES ((SELECT id FROM users WHERE name = ?), ?, ?, ?, ?, ?)`

	rollbackTransactionQuery = `
	UPDATE transactions
	SET reverted_at = NOW()
	WHERE reference = ?
	AND user_id = (SELECT id FROM users WHERE name = ?)`
)
