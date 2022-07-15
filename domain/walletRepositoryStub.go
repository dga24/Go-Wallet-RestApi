package domain

import (
	"bluelabs/errs"
	"fmt"
)

type WalletRepositoryStub struct {
	wallets map[string]Wallet
}

func (s WalletRepositoryStub) GetWallet(id string) (*Wallet, *errs.AppError) {
	for _, w := range s.wallets {
		if w.WalletId == id {
			return &w, nil
		}
	}
	return nil, errs.NewNotFoundError("The wallet doesnt exist")
}

func (s WalletRepositoryStub) CreateWallet(w Wallet) (string, *errs.AppError) {
	newId := fmt.Sprint(len(s.wallets))
	w.WalletId = newId
	s.wallets[newId] = w
	return newId, nil
}

func (s WalletRepositoryStub) Transaction(t Transaction) (*Transaction, *errs.AppError){
	walletId := t.WalletId
	w, err := s.GetWallet(walletId)
	if err != nil {
		return nil, err
	}
	if w.Status =="inactive"{
		return nil, errs.NewValidationError("This wallet is inactive")
	}
	if t.IsWithdrawal(){
		if t.Amount > w.Amount{
			return nil, errs.NewValidationError("Insuficient founds")
		}else{
			w.Amount = w.Amount - t.Amount
			s.wallets[walletId] = *w
			t.Amount = w.Amount
			return &t, nil
		}

	}
	w.Amount = w.Amount + t.Amount
	s.wallets[walletId] = *w
	t.Amount=w.Amount
	return &t, nil
}

func NewWalletRepositoryStub() WalletRepositoryStub {
	wallets := map[string]Wallet{
		"0": Wallet{
			WalletId:    "0",
			CustomerId:  "100",
			OpeningDate: "2022-07-09T11:00:10",
			Amount:      3000,
			Status:      "active",
		},
		"1": Wallet{
			WalletId:    "1",
			CustomerId:  "101",
			OpeningDate: "2022-07-09T11:00:10",
			Amount:      5000,
			Status:      "active",
		},
		"2": Wallet{
			WalletId:    "2",
			CustomerId:  "102",
			OpeningDate: "2022-06-06T11:00:10",
			Amount:      7000,
			Status:      "inactive",
		},
	}
	
	return WalletRepositoryStub{wallets}
}
