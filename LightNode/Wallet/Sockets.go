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

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}


