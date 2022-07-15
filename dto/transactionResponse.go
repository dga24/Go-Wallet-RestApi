package dto

type TransactionResponse struct {
	WalletId        string  `json:"wallett_id"`
	CustomerId      string  `json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"` // "deposit" || "withdraw"
}
