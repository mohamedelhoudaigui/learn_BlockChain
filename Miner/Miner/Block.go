package Miner

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"
)

//----------------impl-----------------

type Block struct {
	Nonce	uint64
	Time	time.Time
	Diffic	uint64
	PrHash	[32]byte
	BlHash	[32]byte
	Trs		[]*Transaction
}

type BlockChain struct {
	Diffic			uint64
	Chain			[]*Block
	TransactionPool	[]*Transaction
	Nblock			uint64
}

func (b *Block) Print(){
	fmt.Println("Block {")
	fmt.Printf("	prev hash: %.15x\n", b.PrHash)
	fmt.Printf("	curr hash: %.15x\n", b.BlHash)
	fmt.Println("	time:", b.Time)
	fmt.Println("	nonce:", b.Nonce)
	fmt.Println("	number of trs:", len(b.Trs))
	fmt.Println("	difficulity:", b.Diffic)
	fmt.Println("}")
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


//---------------Nonce-------------------

func	(b *Block) CalcNonce() {
	b.Nonce = 0
	for {
		hash := b.CalcHash()
		if b.IsValidHash(hash) {
			b.BlHash = hash
			return
		}
		b.Nonce += 1
	}
}

func	(b *Block) CalcHash() [32]byte {
	return sha256.Sum256([]byte(fmt.Sprintf("%d%v%x%v", b.Nonce, b.Time, b.PrHash, b.Trs)))
}

func	(b *Block) IsValidHash(hash [32]byte) bool {
	target := big.NewInt(1)
	target.Lsh(target, 256 - uint(b.Diffic))
	tmp := new(big.Int).SetBytes(hash[:])
	return tmp.Cmp(target) == -1
}


//-----------------------init-----------------------

func NewBlock(prHash [32]byte, ChainDiffic uint64) *Block {
	b := new(Block)
	b.PrHash = prHash
	b.Time = time.Now()
	b.Diffic = ChainDiffic
	b.CalcNonce()
	return (b)
}