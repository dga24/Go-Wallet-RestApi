package dto

import (
	"net/http"
	"testing"
)

func Test_type_of_first_amount(t *testing.T) {
	req := NewWalletRequest{
		CustomerId: "200",
		Amount:     -30,
	}
	appError := req.Validate()
	if appError.Message != "To open a new account you need to deposit at least 10 €" {
		t.Error("Invalid message testing trasaction type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing trasaction type")
	}
}