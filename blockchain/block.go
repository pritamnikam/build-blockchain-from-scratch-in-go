package blockchain

import (
	"bytes"
	"crypto/md5"
	"math/rand"
	"time"
)

type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

// ComputeHash calculates the hash of the block based on its data and previous hash.
func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	hash := md5.Sum(concatenatedData)
	b.Hash = string(hash[:])
}

// CreateBlock creates a new block with the given data and previous hash, computes its hash, and returns it.
func CreateBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	initialNonce := rand.Intn(10000)

	block := &Block{"", data, prevHash, initialNonce, transactions}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

// GenesisBlock creates the first block in the blockchain, known as the genesis block,
// with predefined data and an empty previous hash.
func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}
	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
