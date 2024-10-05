package Miner

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type MinerData struct {
	Chain	[]*Block		`json:"chain_of_blocks"`
	Pool	[]*Transaction	`json:"transaction_pool"`
	Diff	uint64			`json:"difficulity"`
}


func VerifyTransaction(publicKey *rsa.PublicKey, ID, signature []byte) error {
    hashed := sha256.Sum256(ID)
    return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}

func	CreateBlock(Blocks *[]*Block, Diff *uint64, Nblock *uint64) *Block {
	l := len(*Blocks)
	var PrHash [32]byte
	if l != 0 {
		PrHash = (*Blocks)[l - 1].BlHash
	}
	block := NewBlock(PrHash, *Diff)
	*Blocks = append(*Blocks, block)
	*Nblock += 1
	return block
}

func	LastBlock(Blocks *[]*Block) *Block {
	if len(*Blocks) - 1 == -1 {
		return (*Blocks)[0]
	} else {
		return (*Blocks)[len(*Blocks) - 1]
	}
}

func	StartMining(State MinerData) {
	Block := NewBlock(LastBlock(&State.Chain).BlHash, State.Diff)
	Block.Trs = append(Block.Trs, State.Pool...)
	fmt.Println("Block mined !")
}
