package BlockChain

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	trn_pool	[]string // transaction pool
	chain		[]*Block
}

func (bc *BlockChain) CreateBlock(Nonce int, Prhash [32]byte) *Block {
	block := NewBlock(Nonce, Prhash)
	bc.chain = append(bc.chain, block)
	return block
}

func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

//---------------------------------------

func NewBlockChain() *BlockChain {
	block := &Block{}
	bc := new(BlockChain)
	bc.CreateBlock(0, block.Hash())
	return bc
}