package main

import (
	"web3_go/BlockChain"
)


func main() {

	var Difficulity	uint64 = 15
	var WalletPort	string = "2727"
	var MiningPort	string = "2626"

	bc := BlockChain.NewBlockChain(Difficulity, MiningPort, WalletPort) // need to check Ports
	bc.LaunchServer()
	select {}
}
