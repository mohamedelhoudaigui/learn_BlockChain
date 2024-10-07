package main

import (
	"web3_wallet/Wallet"
)

func main() {
	NodeAdress := Wallet.GetOutboundIP()
	WalletPort := "2727"
	FullNodeAdress := NodeAdress + ":" + WalletPort

	w1 := Wallet.NewWallet()
	w2 := Wallet.NewWallet()
	w1.MakeTransaction(w2.PublicKey, 12, FullNodeAdress)
}
