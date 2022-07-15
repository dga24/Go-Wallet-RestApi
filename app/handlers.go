package app

import (
	"bluelabs/dto"
	"bluelabs/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.WalletService
}

func (h Handler) NewWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewWalletRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w,http.StatusBadRequest, err.Error())
	}else{
		request.CustomerId = customerId
		wallet, appError := h.service.NewWallet(request)
		if appError != nil{
			writeResponse(w,appError.Code, appError.Message)
		}else{
			writeResponse(w,http.StatusCreated,wallet)
		}
	}

}

func (h Handler) DoTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	walletId := vars["wallet_id"]
	var req dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w,http.StatusBadRequest, err.Error())
	}
	req.CustomerId = customerId
	req.WalletId = walletId
	resp, appError := h.service.DoTransaction(req)
	if appError!=nil{
		writeResponse(w,appError.Code,appError.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,resp)
	}
}

func (h Handler) FindWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["wallet_id"]
	resp, appError := h.service.FindWallet(id)
	if appError != nil {
		writeResponse(w, appError.Code,appError.AsMessage()) 
	}
	writeResponse(w,http.StatusOK,resp)

}

func writeResponse(w http.ResponseWriter, code int, data interface{}){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err!=nil{
		panic(err)
	}
}