package main

import (
	"fmt"
	"github.com/AyaHikari/assignment01bca/assignment01bca" 
)

func main() {
	// Create Genesis Block (first block) with no previous hash
	genesisBlock := assignment01bca.NewBlock("Genesis to Luffy", 1, "")
	fmt.Println("\n****************************************************************************************************\n")
	fmt.Println("                                       Genesis Block Created\n")
	fmt.Println("****************************************************************************************************\n")

	// Create second block using the hash of the Genesis block
	block1 := assignment01bca.NewBlock("Luffy to Zoro", 2, genesisBlock.Hash)

	// Create third block using the hash of the second block
	block2 := assignment01bca.NewBlock("Zoro to Sanji", 3, block1.Hash)
	
	// Create fourth block using the hash of the second block
	assignment01bca.NewBlock("Sanji to Law", 4, block2.Hash)

	// Print the entire blockchain
	fmt.Println("\n****************************************************************************************************\n")
	fmt.Println("                            Displaying all blocks in the Blockchain\n")
	fmt.Println("****************************************************************************************************\n")
	assignment01bca.ListBlocks()

	// Modify the transaction in block 1
	fmt.Println("\n****************************************************************************************************\n")
	fmt.Println("                                  Modifying Block 1 Transaction...")
	fmt.Println("\n****************************************************************************************************\n")
	assignment01bca.ChangeBlock(1, "Luffy to Zoro Updated")
	assignment01bca.ListBlocks()

	// Verify the blockchain integrity
	fmt.Println("\n-****************************************************************************************************\n")
	fmt.Println("                        Verifying Blockchain integrity after Modification")
	fmt.Println("\n****************************************************************************************************\n")
	assignment01bca.VerifyChain()
}
