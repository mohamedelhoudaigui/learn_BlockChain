package main

import (
	"web3_go/BlockChain"
)


func main() {
	bc := BlockChain.NewBlockChain()
	bc.LaunchServer()
	select {}
}
