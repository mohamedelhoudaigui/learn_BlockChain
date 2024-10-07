package main

import (
	"web3_miner/Miner"
)

func main() {

	NodeAddr := Miner.GetOutboundIP()
	MiningFullNodePort := "2626"
	FullNodeAddr := NodeAddr + ":" + MiningFullNodePort
	data := ""
	bc := Miner.BlockChain{}
	PingData := []byte(data)

	Miner.MinerServer("2525", &bc, &PingData, &FullNodeAddr)
	Miner.StartMining(&bc)
}
