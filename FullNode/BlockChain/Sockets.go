package BlockChain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func	ReqFromWallet(Conn net.Conn, Data []byte, bc *BlockChain) {
	defer Conn.Close()
	var tx Transaction

	err := json.Unmarshal(Data, &tx)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return
	}
	bc.TransactionPool = append(bc.TransactionPool, &tx)
}

func	ReqFromMiner(Conn net.Conn, bc *BlockChain) {
	defer Conn.Close()
	var	State MinerData

	State.Chain = bc.Chain
	State.Diff = bc.Diffic
	State.Pool = bc.TransactionPool
	
	MState, _ := json.Marshal(State)
	BytesWrited, err := Conn.Write(MState)
	fmt.Println(BytesWrited, "-", err)
}

func HandleReq(Conn net.Conn, bc *BlockChain, Port string) {

	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, Conn)
	if err != nil && err != io.EOF {
		log.Printf("Error reading from connection: %v", err)
		return
	}
	Data := bytes.Trim(buffer.Bytes(), "\x00")

	if Port == "2727"  { // remove static ports
		ReqFromWallet(Conn, Data, bc)
	} else if Port == "2626" {
		ReqFromMiner(Conn, bc)
	}
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func	Server(bc *BlockChain, Port string) {
	ip := GetOutboundIP() + ":" + Port
	fmt.Printf("server started at : %s\n", ip)
	ln, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		Conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
  		HandleReq(Conn, bc, Port)
	}
}
