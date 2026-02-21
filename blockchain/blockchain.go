package blockchain

type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block with the given data to the blockchain.
// It retrieves the hash of the last block to set as the previous hash for the new block.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// InitBlockChain creates a new blockchain with a single genesis block and returns it.
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
