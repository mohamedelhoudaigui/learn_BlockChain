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
	MiningPort		string
	MinerPort		string
	WalletPort		string
	Nblock			uint64
}

func NewBlockChain(Difficulity uint64, MiningPort string, WalletPort string, MinerPort string) *BlockChain {
	bc := new(BlockChain)
	bc.Diffic = Difficulity
	bc.Nblock = 0
	bc.WalletPort = WalletPort
	bc.MiningPort = MiningPort
	bc.MinerPort = MinerPort
	bc.CreateBlock()
	return bc
}

//-------------server--------------


func	(bc *BlockChain) LaunchServer() {
	go Server(bc, bc.MiningPort, "Mining")
	go Server(bc, bc.WalletPort, "Wallet")
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
	fmt.Printf("%s\n", strings.Repeat("$", 60))
	fmt.Println("Transaction pool :")
	for _, T := range bc.TransactionPool {
		T.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}
