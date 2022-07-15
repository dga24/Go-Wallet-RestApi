package domain

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type Transaction struct {
	WalletId        string  `json:"account_id"`
	CustomerId      string  `json:"-"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}