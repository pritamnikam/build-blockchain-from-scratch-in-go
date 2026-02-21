package blockchain

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/big"
)

// Difficulty is the number of leading zeros required in the hash for a valid proof of work.
const Difficulty = 10

// ProofOfWork represents the proof of work for a block in the blockchain.
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork creates a new ProofOfWork instance for a given block.
func NewProofOfWork(block *Block) *ProofOfWork {
	targetVal := big.NewInt(1)
	targetVal.Lsh(targetVal, uint(256-Difficulty))
	return &ProofOfWork{Block: block, Target: targetVal}
}

// ComputeData prepares the data for hashing by combining the block's previous hash, data, nonce, and difficulty.
func (pow *ProofOfWork) ComputeData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash),
			[]byte(pow.Block.Data),
			make([]byte, 8),
			make([]byte, 8),
		},
		[]byte{},
	)

	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))

	return data
}

// MineBlock performs the proof of work algorithm to find a valid nonce and hash for the block.
func (pow *ProofOfWork) MineBlock() (int, []byte) {
	var intHash big.Int
	var computedHash [16]byte

	nonce := 0
	for {
		computedData := pow.ComputeData(nonce)
		computedHash = md5.Sum(computedData)

		fmt.Printf("\r%x", computedHash)

		intHash.SetBytes(computedHash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		}

		nonce++
	}
	fmt.Println()

	pow.Block.Nonce = nonce
	return nonce, computedHash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	computedData := pow.ComputeData(pow.Block.Nonce)

	computedHash := md5.Sum(computedData)
	intHash.SetBytes(computedHash[:])

	if intHash.Cmp(pow.Target) == -1 {
		return true
	} else {
		return false
	}
}
