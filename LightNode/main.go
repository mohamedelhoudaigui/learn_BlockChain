package main

import (
	"web3_wallet/Wallet"
)


func main() {
	NodeAdress := "10.12.9.7:2727"
	w1 := Wallet.NewWallet()
	w2 := Wallet.NewWallet()
	w1.MakeTransaction(w2.PublicKey, 12, NodeAdress)
}