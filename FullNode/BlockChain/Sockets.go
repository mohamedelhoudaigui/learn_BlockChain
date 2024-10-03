package BlockChain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func HandleReq(conn net.Conn, pool *[]*Transaction) *Transaction {
	defer conn.Close()

	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, conn)
	if err != nil && err != io.EOF {
		log.Printf("Error reading from connection: %v", err)
		return nil
	}
	data := bytes.Trim(buffer.Bytes(), "\x00")

	var tx Transaction
	err = json.Unmarshal(data, &tx)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return nil
	}

	*pool = append(*pool, &tx)
	tx.Print()
	return &tx
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

func Server(Pool *[]*Transaction, bc *BlockChain) {
	ip := GetOutboundIP() + ":2727"
	fmt.Printf("server started at : %s\n", ip) // change to ip
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
  		HandleReq(Conn, Pool)
		//Validate(tx)
	}
}
