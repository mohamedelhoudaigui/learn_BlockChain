package BlockChain

import (
	"time"
	"fmt"
)

type Block struct {
	nonce int
	prHash string
	time int64
	trs []string
}


func (p *Block) Print(){
	fmt.Println("Block", p.nonce, "{")
	fmt.Println("	prev hash:", p.prHash)
	fmt.Println("	time:", p.time)
	fmt.Println("	transactions:", p.trs)
	fmt.Println("}")
}

//----------------------------------------------------

func NewBlock(nonce int, prHash string) *Block {
	b := new(Block)
	b.time = time.Now().UnixNano()
	b.nonce = nonce
	b.prHash = prHash
	return (b)
}