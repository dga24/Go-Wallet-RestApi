# bluelabs
 Challenge


I made an ApiRest app to implement the exercise
To run de program have to define SERVER_ADDRESS and SERVER_PORT


Query the current state of a wallet.
Get Wallet Request
GET  "/wallet/{wallet_id:[0-9]+}"
Example:
GET localhost:8003/wallet/1



Query to create a wallet.
New Wallet Request
POST "/wallet/{customer_id:[0-9]+}/wallet"
`json:"amount"`
**Customer_id is a number, you can invent the customerId, or use someon already created. One customer can have more thar one wallet

Example: 
POST localhost:8003/wallet/108/wallet
{
    "amount": 3000
}



Query to add or remove funds from a wallet. 
Transaction Request
POST" /wallet/{customer_id:[0-9]+}/wallet/{wallet_id:[0-9]+}
`json:"amount"`
`json:"transaction_type"` // "deposit" || "withdrawal"
**To do a transaction, the customer have to be the owner of the wallet (see wallet list in next lines)
**Minim amount: 10

Example:
POST localhost:8003/wallet/100/wallet/0
{
    "amount": 400,
    "transaction_type": "withdrawal"
}




How this is only a test, this project dont use any DB, insted of this, in the file WalletRepositoryStub you will find a map with 3 example of wallets:
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
  
  
  
  
