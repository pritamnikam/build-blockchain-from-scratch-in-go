package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	// Initialize the blockchain
	chain := blockchain.InitBlockChain()

	// Add some blocks to the blockchain
	chain.AddBlock("Block #1")
	chain.AddBlock("Block #2")
	chain.AddBlock("Block #3")

	// Print the blockchain
	for i, block := range chain.Blocks {
		fmt.Printf("Block #%d:\n", i)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n\n", block.PrevHash)
	}

}
