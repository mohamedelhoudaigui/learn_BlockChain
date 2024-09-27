package main

import (
	"log"
	"web3_go/BlockChain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}


func main() {
	bc := BlockChain.NewBlockChain()
	bc.AddBlock()
	bc.AddBlock()
	bc.AddBlock()
	bc.AddBlock()
	bc.AddBlock()
	bc.Print()
} 