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

func (d *database) GetBalance(playerName, currency string) (int, error) {
	var balance sql.NullInt64
	err := d.QueryRow(getBalanceQuery, playerName, currency).Scan(&balance)
	if err != nil {
		return 0, err
	}
	if !balance.Valid {
		return 0, nil
	}
	return int(balance.Int64), nil
}

func getBalance(db *sql.Tx, playerName, currency string) (int, error) {
	var balance sql.NullInt64
	err := db.QueryRow(getBalanceQuery, playerName, currency).Scan(&balance)
	if err != nil {
		return 0, err
	}
	if !balance.Valid {
		return 0, nil
	}
	return int(balance.Int64), nil
}

func getTransaction(db *sql.Tx, playerName, transactionRef string) (Transaction, error) {
	var t Transaction
	return t, db.QueryRow(getTransactionQuery, playerName, transactionRef).Scan(
		&t.Id, &t.UserId, &t.CallerID, &t.Reference, &t.Withdraw, &t.Deposit, &t.Currency, &t.CreatedAt, &t.RevertedAt,
	)
}

var (
	ErrInsufficientFunds   = errors.New("insufficient funds")
	ErrTransactionReverted = errors.New("transaction reverted")
)

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

	// проверяем существует ли транзакция
	t, err := getTransaction(tx, playerName, transactionRef)
	if err != nil && err != sql.ErrNoRows {
		return Transaction{}, 0, err
	}

	// если транзакция существует, то проверяем не откачена ли она
	if err == nil {
		if t.RevertedAt != nil {
			// если транзакция откачена, то возвращаем ошибку
			return Transaction{}, 0, ErrTransactionReverted
		} else {
			// если транзакция не откачена, то про принципипу идемпотентности возвращаем её и текущий баланс
			balance, err := getBalance(tx, playerName, currency)
			if err != nil {
				return Transaction{}, 0, err
			}

			if err = tx.Commit(); err != nil {
				return Transaction{}, 0, err
			}

			return Transaction{}, balance, nil
		}
	}

	// если транзакции не существует, то проверяем можно ли её провести
	balance, err := getBalance(tx, playerName, currency)
	if err != nil {
		return Transaction{}, 0, err
	}

	// недостаточно средств - не можем провести транзакцию
	if balance < withdraw {
		return Transaction{}, 0, ErrInsufficientFunds
	}

	// регистрируем транзакцию
	if _, err = tx.Exec(registerTransactionQuery, playerName, callerID, transactionRef, withdraw, deposit, currency); err != nil {
		return Transaction{}, 0, err
	}

	// получаем id транзакции
	t, err = getTransaction(tx, playerName, transactionRef)
	if err != nil {
		return Transaction{}, 0, err
	}

	// получаем текущий баланс
	resultBalance := balance - withdraw + deposit

	if err = tx.Commit(); err != nil {
		return Transaction{}, 0, err
	}

	// возвращаем баланс после транзакции и саму транзакцию для "transactionId"
	return t, resultBalance, nil
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

	// проверяем существует ли транзакция
	var t Transaction
	t, err = getTransaction(tx, playerName, transactionRef)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		// если транзакции не существует, то регистрируем её как откаченную
		if _, err = tx.Exec(registerRevertedTransactionQuery, playerName, callerId, transactionRef); err != nil {
			return err
		}
	} else {
		// если транзакция существует, то откачиваем ее в том слуае если она не откачена
		// если уже откачена, то ничего не делаем
		if t.RevertedAt == nil {
			if _, err = tx.Exec(rollbackTransactionQuery, playerName, transactionRef); err != nil {
				return err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

const (
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

	registerRevertedTransactionQuery = `
	INSERT INTO transactions (user_id, caller_id, reference, withdraw, deposit, currency, reverted_at)
	VALUES ((SELECT id FROM users WHERE name = ?), ?, ?, 0, 0, 'USD', NOW())`

	rollbackTransactionQuery = `
	UPDATE transactions
	SET reverted_at = NOW()
	WHERE reference = ?
	AND user_id = (SELECT id FROM users WHERE name = ?)`
)
