package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"crypto/sha256"
	"encoding/json"
)
type Block struct {
	nonce int
	previousHash string
	timeStamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timeStamp = time.Now().Unix()
	return b
}

func (b *Block) Print(){
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("timeStamp: %d\n", b.timeStamp)
	fmt.Printf("transactions: %s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init block")
	return bc
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce int
		PreviousHash string
		TimeStamp int64
		Transactions []string
	}{
		Nonce: b.nonce,
		PreviousHash: b.previousHash,
		TimeStamp: b.timeStamp,
		Transactions: b.transactions,
	})
}


func (b *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	block := NewBlock(nonce, previousHash)
	b.chain = append(b.chain, block)
	return block
}


func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n",strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func main() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
	// bc := NewBlockchain()
	// bc.Print()
	// bc.CreateBlock(5, "block 1")
	// bc.Print()
	// bc.CreateBlock(5, "block 2")
	// bc.Print()
}



func init() {
	log.SetPrefix("Blockchain: ")
}

// func main() {
// 	log.Println("Starting blockchain")
// 	b := NewBlock(0, "init block")
// 	b.Print()
// }