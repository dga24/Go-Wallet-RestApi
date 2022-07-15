package service

import (
	"bluelabs/domain"
	"bluelabs/dto"
	"bluelabs/errs"
	"time"

	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -destination=../mocks/service/mockWalletService.go -package=service bluelabs/service WalletService
type WalletService interface{
	NewWallet(walet dto.NewWalletRequest) (*dto.NewWalletResponse, *errs.AppError)
	DoTransaction(transactionRequest dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	FindWallet(WalletId string) (*dto.GetWalletResponse, *errs.AppError)
}

type DefaultWalletService struct{
	repo domain.WalletRepository
}


func (s DefaultWalletService) NewWallet(req dto.NewWalletRequest) (*dto.NewWalletResponse, *errs.AppError){
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	w := domain.Wallet{
		CustomerId: req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Amount: req.Amount,    
		Status: "active",
	}
	id,err := s.repo.CreateWallet(w)
	if err != nil {
		return nil, err
	}
	response := dto.NewWalletResponse{
		WalletId: id,
	}
	return &response, nil

}


func (s DefaultWalletService) DoTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError){
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	wallet, err := s.FindWallet(req.WalletId)
	if err != nil {
		return nil, errs.NewNotFoundError("The data doesnt match with any wallet")
	}
	if wallet.CustomerId != req.CustomerId{
		return nil, errs.NewValidationError("Customer is not owner of this wallet")
	}
	t := domain.Transaction{
		WalletId: req.WalletId, 
		CustomerId: req.CustomerId,  
		Amount: req.Amount,         
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	resp, err := s.repo.Transaction(t)
	if err != nil {
		return nil, err
	}
	dtoResp := dto.TransactionResponse{
		WalletId: resp.WalletId,
		CustomerId: resp.CustomerId,     
		TransactionType: resp.TransactionType,
		Amount: resp.Amount,         
	}
	return &dtoResp, nil

}


func (s DefaultWalletService) FindWallet(walletId string) (*dto.GetWalletResponse, *errs.AppError){
	w, err := s.repo.GetWallet(walletId)
	if err != nil {
		return nil, err
	}
	resp := dto.GetWalletResponse{
		WalletId: w.WalletId,
		CustomerId: w.CustomerId,
		OpeningDate: w.OpeningDate,
		Amount: w.Amount,
		Status: w.Status,
	}
	return &resp, nil

}

func NewWalletService(repository domain.WalletRepository) DefaultWalletService{
	return DefaultWalletService{repository}
}

