package main

import (
	"web3_miner/Miner"
	_ "fmt"
)


func main() {
	FullNodeAddr := "10.12.9.7:2626"
	data := ""
	bc := Miner.BlockChain{}
	PingData := []byte(data)

	Miner.MinerServer("2525", &bc, &PingData, &FullNodeAddr)
	Miner.StartMining(&bc)
}
