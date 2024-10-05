package main

import (
	"web3_miner/Miner"
	_ "fmt"
)


func main() {
	FullNodeAddr := "10.12.2.13:2626"
	data := ""
	State := Miner.MinerData{}
	data_b := []byte(data)
	go Miner.MinerServer("2525", &State)
	Miner.Client(&data_b, FullNodeAddr)
	Miner.StartMining(State)
	select {}
}
