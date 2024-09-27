package BlockChain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

)

type Block struct {
	Nonce int
	PrHash [32]byte
	Time int64
	Trs []string
}


func (p *Block) Print(){
	fmt.Println("Block", p.Nonce, "{")
	fmt.Println("	prev hash:", p.PrHash)
	fmt.Println("	time:", p.Time)
	fmt.Println("	transactions:", p.Trs)
	fmt.Println("}")
}

func (p *Block) Hash() [32]byte {
    hash, _ := json.Marshal(p)
    return sha256.Sum256(hash)
}

//----------------------------------------------------

func NewBlock(nonce int, prHash [32]byte) *Block {
	b := new(Block)
	b.Time = time.Now().UnixNano()
	b.Nonce = nonce
	b.PrHash = prHash
	return (b)
}