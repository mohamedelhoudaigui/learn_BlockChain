package Wallet

import (
	"log"
	"net"
)

func Client(Data *[]byte, NodeAddress string) {
	conn, err := net.Dial("tcp", NodeAddress)
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(*Data)
	if err != nil {
		log.Printf("Error sending data: %v", err)
		return
	}
}

