package Miner

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)


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

func	LastBlock(bc *BlockChain) *Block {
	if bc.Nblock == 0 {
		return nil
	} else {
		return bc.Chain[bc.Nblock - 1]
	}
}

func	StartMining(bc *BlockChain) {
	LBlock := LastBlock(bc)
	if LBlock != nil {
		Block := NewBlock(LBlock.BlHash, bc.Diffic)
		Block.Trs = append(Block.Trs, bc.TransactionPool...)
		fmt.Println("Block mined !")
	} else {
		fmt.Println("No block in the chain !!")
	}
}
