package dto

import "bluelabs/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	WalletId        string  `json:"wallett_id"`
	CustomerId      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"` // "deposit" || "withdraw"
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}