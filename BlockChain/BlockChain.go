package BlockChain

import (
	"fmt"
	"strings"
)

//------------impl----------------

type BlockChain struct {
	Diffic			uint64
	Chain			[]*Block
	TransactionPool	[]*Transaction
	Nblock			uint64
}

func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	bc.Diffic = 0
	bc.Nblock = 1
	bc.CreateBlock()
	return bc
}

//-------------server--------------



func	(bc *BlockChain) LaunchServer() {
	Server(&bc.TransactionPool)
}

//----------------------------

func	(bc *BlockChain) CreateBlock() *Block {
	l := len(bc.Chain)
	var PrHash [32]byte
	if l != 0 {
		PrHash = bc.Chain[len(bc.Chain) - 1].BlHash
	}

	block := NewBlock(PrHash, bc.Diffic)
	bc.Chain = append(bc.Chain, block)
	bc.Nblock += 1
	return block
}

func	(bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain) - 1]
}

func (bc *BlockChain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s block number %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}
