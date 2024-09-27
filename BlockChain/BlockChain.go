package BlockChain

import (
	"fmt"
	"strings"
	"web3_go/Hash"
)

type BlockChain struct {
	trn_pool	[]string // transaction pool
	chain		[]*Block
	blockN      int
}

func (bc *BlockChain) AddBlock() *Block {
	prHash := Hash.GetHash(bc.blockN - 1,
				int(bc.chain[bc.blockN - 1].time),
				bc.chain[bc.blockN - 1].prHash)

	block := NewBlock(bc.blockN, prHash)
	bc.chain = append(bc.chain, block)
	bc.blockN += 1
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
	blockchain := new(BlockChain)
	iniBlock := NewBlock(0, "0270")
	blockchain.chain = append(blockchain.chain, iniBlock)
	blockchain.blockN = 1
	return blockchain
}