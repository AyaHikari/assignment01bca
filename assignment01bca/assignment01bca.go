package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv" // used for string-to-integer conversions
)

// Block structure representing a single block in the blockchain
type Block struct {
	Transaction        string
	NonceX             int
	Previous_BlockHash string
	Hash               string
}

// Blockchain structure holding an array of blocks
type Blockchain struct {
	Blocks []*Block
}

// Global blockchain variable
var chain = Blockchain{}

// CalculateHash computes the hash of the given string data using SHA-256
func CalculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// NewBlock creates a new block and appends it to the blockchain
func NewBlock(transaction string, noncex int, previous_BlockHash string) *Block {
	newBlock := &Block{
		Transaction:        transaction,
		NonceX:             noncex,
		Previous_BlockHash: previous_BlockHash,
	}

	// Combine transaction, nonce, and previous block hash for new block's data
	blockData := transaction + strconv.Itoa(noncex) + previous_BlockHash
	// Calculate and assign the block's hash
	newBlock.Hash = CalculateHash(blockData)

	// Add the new block to the blockchain
	chain.Blocks = append(chain.Blocks, newBlock)

	return newBlock
}

// ListBlocks prints the details of all blocks in the blockchain
func ListBlocks() {
	for i, block := range chain.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("\tTransaction: %s\n", block.Transaction)
		fmt.Printf("\tNonce: %d\n", block.NonceX)
		fmt.Printf("\tPrevious Block Hash: %s\n", block.Previous_BlockHash)
		fmt.Printf("\tHash: %s\n\n", block.Hash)
	}
}

// ChangeBlock modifies the transaction of a specific block and recalculates its hash
func ChangeBlock(blockNo int, newTransaction string) {
	if blockNo < len(chain.Blocks) {
		block := chain.Blocks[blockNo]
		block.Transaction = newTransaction

		// Recalculate the block's hash after transaction update
		blockData := newTransaction + strconv.Itoa(block.NonceX) + block.Previous_BlockHash
		block.Hash = CalculateHash(blockData)
	}
}

// VerifyChain verifies the integrity of the blockchain by checking each block's hash
func VerifyChain() bool {
	for i := 1; i < len(chain.Blocks); i++ {
		previousBlock := chain.Blocks[i-1]
		currentBlock := chain.Blocks[i]

		// Recalculate the hash for the current block
		blockData := currentBlock.Transaction + strconv.Itoa(currentBlock.NonceX) + currentBlock.Previous_BlockHash
		calculatedHash := CalculateHash(blockData)

		// Check if the current block's hash matches the recalculated hash
		// and if the previous block's hash matches the current block's previous hash
		if currentBlock.Hash != calculatedHash || currentBlock.Previous_BlockHash != previousBlock.Hash {
			fmt.Printf("Block %d has been tampered!\n", i)
			return false
		}
	}

	// Blockchain is valid if no issues are found
	fmt.Println("\n-****************************************************************************************************\n")
	fmt.Println("                                     Blockchain is Valid!")
	fmt.Println("\n-****************************************************************************************************\n")

	return true
}
