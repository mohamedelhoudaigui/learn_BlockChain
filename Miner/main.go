package main

import (
	_ "fmt"
	"web3_miner/Miner"
)

func main() {
	FullNodeAddr := "10.12.13.4:2626"
	data := ""
	bc := Miner.BlockChain{}
	PingData := []byte(data)

	Miner.MinerServer("2525", &bc, &PingData, &FullNodeAddr)
	Miner.StartMining(&bc)
}
