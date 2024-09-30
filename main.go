package main

import (
	"log"
	"sync"
	"web3_go/BlockChain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func routine() {
	bc := BlockChain.NewBlockChain()
	bc.LaunchServer()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go routine()
	w1 := BlockChain.NewWallet()
	w2 := BlockChain.NewWallet()
	w1.MakeTransaction(w2.PublicKey, 12)
	wg.Wait()
}