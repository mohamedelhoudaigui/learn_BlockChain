package Miner

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io"
	_ "io"
	"log"
	_ "log"
	"net"
)


func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}


func	GetState(Conn net.Conn) MinerData {

	defer Conn.Close()

	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, Conn)
	if err != nil && err != io.EOF {
		log.Println("Error reading from connection :", err)
	}
	Data := bytes.Trim(buffer.Bytes(), "\x00")

	var	State MinerData

	err = json.Unmarshal(Data, &State)
	if err != nil {
		log.Println("Error marshaling to MinerData (Miner->Sockets)", err)
	}
	return State
}

func	MinerServer(Port string, State *MinerData) {
	ip := GetOutboundIP() + ":" + Port
	fmt.Println("mining server started at :", ip)
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
		*State = GetState(Conn)
	}
}

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