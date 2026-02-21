package blockchain

// Transaction represents a transfer of value between a sender and a receiver,
// with an amount and a flag indicating if it's a coinbase transaction.
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

// Blockchain represents the collection of blocks in the blockchain.
type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block with the given data to the blockchain.
// It retrieves the hash of the last block to set as the previous hash for the new block.
func (bc *Blockchain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseRcpt,
		Amount:   10.0,
		Coinbase: true,
	}
	newBlock := CreateBlock(
		data,
		prevBlock.Hash,
		append([]*Transaction{coinbaseTransaction}, transactions...),
	)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// InitBlockChain creates a new blockchain with a single genesis block and returns it.
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
