package Miner

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
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


func	GetState(Conn net.Conn) {
	defer Conn.Close()

	buf := make([]byte, 2048)
	_, err := Conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	Data := bytes.Trim(buf, "\x00")

	var	State MinerData

	err = json.Unmarshal(Data, &State)
	if err != nil {
		log.Printf("Error marshaling to MinerData: %v", err)
		return
	}
	fmt.Println(State)
}

func	MinerServer(Port string) {
	ip := GetOutboundIP() + ":" + Port
	fmt.Printf("mining server started at : %s\n", ip)
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
		go GetState(Conn)
	}
}
