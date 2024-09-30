package BlockChain

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net"
)

func Client(data *[]byte) {
	conn, err := net.Dial("tcp", "localhost:2727")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write(*data)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Close()
}

func HandleReq(Conn net.Conn, _Pool []*[]byte) {
	defer Conn.Close()
	buf := make([]byte, 1024)
	_, err := Conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}


func Server(_Pool [] *[]byte) {
    ln, err := net.Listen("tcp", ":2727")
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
        go HandleReq(Conn, _Pool)
    }
}
