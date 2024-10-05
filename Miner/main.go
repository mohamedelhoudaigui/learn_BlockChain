package main

import (
	"web3_miner/Miner"
	_ "fmt"
)


func main() {
	//FullNodeAddr := "10.12.13.2:2727"
	go Miner.MinerServer("2525")
	select {}
	//Miner.StartMining(State)
}
