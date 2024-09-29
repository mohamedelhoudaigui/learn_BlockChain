package BlockChain

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	Chain		[]*Block
}

func (bc *BlockChain) CreateBlock() *Block {
	l := len(bc.Chain)
	var PrHash [32]byte
	if l != 0 {
		PrHash = bc.Chain[len(bc.Chain) - 1].BlHash
	}
	block := NewBlock(PrHash)
	bc.Chain = append(bc.Chain, block)
	return block
}

func (bc *BlockChain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s block number %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

//---------------------------------------

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.CreateBlock()
	return bc
}