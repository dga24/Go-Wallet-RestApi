package dto

type GetWalletResponse struct {
	WalletId    string  `json:"wallett_id"`
	CustomerId  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}
