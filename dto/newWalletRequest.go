package dto

import "bluelabs/errs"

type NewWalletRequest struct {
	CustomerId string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
}

func (r NewWalletRequest) Validate() *errs.AppError {
	if r.Amount < 10 {
		return errs.NewValidationError("To open a new account you need to deposit at least 10 €")
	}
	return nil
}