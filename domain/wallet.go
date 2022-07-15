package domain

import (
	"bluelabs/errs"
)

type Wallet struct {
	WalletId    string  `json:"wallett_id"`
	CustomerId  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type WalletRepository interface {
	CreateWallet(wallet Wallet) (string, *errs.AppError)
	Transaction(transaction Transaction) (*Transaction, *errs.AppError)
	GetWallet(walletId string) (*Wallet, *errs.AppError)
}
