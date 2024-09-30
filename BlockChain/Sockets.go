package BlockChain

import (
	"bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func Client(tx *[]byte) {

	conn, err := net.Dial("tcp", "localhost:2727")
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(*tx)
	if err != nil {
		log.Printf("Error sending data: %v", err)
		return
	}
	fmt.Println("Transaction sent successfully")
}

func HandleReq(conn net.Conn, _pool *[]*Transaction) {

	defer conn.Close()

	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, conn)
	if err != nil && err != io.EOF {
		log.Printf("Error reading from connection: %v", err)
		return
	}

	data := bytes.Trim(buffer.Bytes(), "\x00")

	var transaction Transaction
	err = json.Unmarshal(data, &transaction)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	//*pool = append(*pool, &transaction)

	transaction.Print()
}


func Server(_Pool *[]*Transaction) {
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
