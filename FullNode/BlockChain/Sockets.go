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

	MinerAddr := "10.12.9.7:2525" // how to automate this ??
	Conn.Close()
	State := NewMinerData(bc)
	JsonState, err := json.Marshal(*State)
	if err != nil {
		log.Printf("Error Marshaling MinerData (FullNode->Sockets): %v", err) //-------------------------------------------------------
	}
	State.Print()
	Client(JsonState, MinerAddr)
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

func Client(Data []byte, Address string) {

	Conn, err := net.Dial("tcp", Address)
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer Conn.Close()
	_, err = Conn.Write(Data)
	if err != nil {
		log.Printf("Error sending data: %v", err)
		return
	}
}
