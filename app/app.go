package app

import (
	"bluelabs/domain"
	"bluelabs/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {

	initCheck()
	router := mux.NewRouter()

	walletRepositorystub := domain.NewWalletRepositoryStub()
	h := Handler{service.NewWalletService(walletRepositorystub)}

	router.HandleFunc("/wallet/{wallet_id:[0-9]+}", h.FindWallet).Methods(http.MethodGet)
	router.HandleFunc("/wallet/{customer_id:[0-9]+}/wallet", h.NewWallet).Methods(http.MethodPost)
	router.HandleFunc("/wallet/{customer_id:[0-9]+}/wallet/{wallet_id:[0-9]+}", h.DoTransaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func initCheck(){
	if os.Getenv("SERVER_ADDRESS")=="" ||
	os.Getenv("SERVER_PORT")== ""{
		log.Fatal("Environment variables not defined, define SERVER_ADDRESS and SERVER_PORT")
	}
}