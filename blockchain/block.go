package blockchain

import (
	"bytes"
	"crypto/md5"
)

type Block struct {
	Hash     string
	Data     string
	PrevHash string
}

// ComputeHash calculates the hash of the block based on its data and previous hash.
func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	hash := md5.Sum(concatenatedData)
	b.Hash = string(hash[:])
}

// CreateBlock creates a new block with the given data and previous hash, computes its hash, and returns it.
func CreateBlock(data string, prevHash string) *Block {
	block := &Block{
		Data:     data,
		PrevHash: prevHash,
	}

	block.ComputeHash()
	return block
}

// GenesisBlock creates the first block in the blockchain, known as the genesis block, with predefined data and an empty previous hash.
func GenesisBlock() *Block {
	return CreateBlock("Genesis Block", "")
}
