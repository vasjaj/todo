package server

import (
	"net/http"
	"strconv"

	"github.com/vasjaj/todo/internal/database"

	"github.com/go-playground/validator"
)

type SeamlessService struct {
	db database.Database
}

type GetBalanceRequest struct {
	CallerId             int    `json:"callerId" validate:"required,numeric"`
	PlayerName           string `json:"playerName" validate:"required"`
	Currency             string `json:"currency" validate:"required,len=3"`
	GamedId              string `json:"gamedId"`
	SessionId            string `json:"sessionId"`
	SessionAlternativeId string `json:"sessionAlternativeId"`
	BonusId              string `json:"bonusId"`
	Balance              string `json:"balance"`
	FreeroundsLeft       string `json:"freeroundsLeft"`
}

type GetBalanceResponse struct {
	Balance int `json:"balance"`
}

func (s *SeamlessService) GetBalance(_ *http.Request, req *GetBalanceRequest, res *GetBalanceResponse) error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	balance, err := s.db.GetBalance(req.PlayerName, req.Currency)
	if err != nil {
		return err
	}

	res.Balance = balance

	return nil
}

type WithdrawAndDepositRequest struct {
	CallerId             int         `json:"callerId" validate:"required,numeric"`
	PlayerName           string      `json:"playerName" validate:"required"`
	Withdraw             int         `json:"withdraw" validate:"required,numeric"`
	Deposit              int         `json:"deposit" validate:"required,numeric"`
	Currency             string      `json:"currency" validate:"required,len=3"`
	TransactionRef       string      `json:"transactionRef" validate:"required"`
	GameRoundRef         string      `json:"gameRoundRef"`
	GameId               string      `json:"gameId"`
	Source               string      `json:"source"`
	Reason               string      `json:"reason" validate:"oneof=GAME_PLAY GAME_PLAY_FINAL"`
	SessionId            string      `json:"sessionId"`
	SessionAlternativeId string      `json:"sessionAlternativeId"`
	SpinDetails          SpinRequest `json:"spinDetails"`
	BonusId              string      `json:"bonusId"`
	ChargeFreerounds     int         `json:"chargeFreerounds"`
}

type SpinRequest struct {
	BetType string `json:"betType"`
	WinType string `json:"winType"`
}

type WithdrawAndDepositResponse struct {
	NewBalance     int    `json:"newBalance" validate:"required,numeric"`
	TransactionID  string `json:"transactionId" validate:"required"`
	FreeroundsLeft int    `json:"freeroundsLeft"`
}

func (s *SeamlessService) WithdrawAndDeposit(_ *http.Request, req *WithdrawAndDepositRequest, res *WithdrawAndDepositResponse) error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	transaction, amount, err := s.db.RegisterTransaction(req.CallerId, req.Withdraw, req.Deposit, req.PlayerName, req.Currency, req.TransactionRef)
	if err != nil {
		return err
	}

	res.NewBalance = amount
	res.TransactionID = strconv.Itoa(transaction.Id)

	return nil
}

type RollbackTransactionRequest struct {
	CallerId             int    `json:"callerId" validate:"required,numeric"`
	PlayerName           string `json:"playerName" validate:"required"`
	TransactionRef       string `json:"transactionRef" validate:"required"`
	GameID               string `json:"gameId"`
	SessionId            string `json:"sessionId"`
	SessionAlternativeId string `json:"sessionAlternativeId"`
	RoundId              string `json:"roundId"`
}

type RollbackTransactionResponse struct{}

func (s *SeamlessService) RollbackTransaction(_ *http.Request, req *RollbackTransactionRequest, res *RollbackTransactionResponse) error {
	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return s.db.RollbackTransaction(req.PlayerName, req.TransactionRef)
}
