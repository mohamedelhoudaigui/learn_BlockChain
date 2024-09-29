package BlockChain

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

type Block struct {
	Nonce	uint64
	Time	time.Time
	Diffic	uint64
	PrHash	[32]byte
	BlHash	[32]byte
	Trs		[]*Transaction
}

func (b *Block) Print(){
	fmt.Println("Block {")
	fmt.Printf("	prev hash: %x\n", b.PrHash)
	fmt.Printf("	curr hash: %x\n", b.BlHash)
	fmt.Println("	time:", b.Time)
	fmt.Println("	nonce:", b.Nonce)
	fmt.Println("	number of trs:", len(b.Trs))
	fmt.Println("	difficulity:", b.Diffic)
	fmt.Println("}")
}

//------------------------------------

// CalcNonce -> CalcHash -> IsValidHash

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
	target.Lsh(target, 256-uint(b.Diffic))
	tmp := new(big.Int).SetBytes(hash[:])
	return tmp.Cmp(target) == -1
}

//----------------------------------------------------

func NewBlock(prHash [32]byte) *Block {
	b := new(Block)
	b.PrHash = prHash
	b.Time = time.Now()
	b.Diffic = uint64(rand.Intn(3) + 1)
	b.CalcNonce()
	return (b)
}