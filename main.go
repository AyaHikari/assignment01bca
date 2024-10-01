package main

import (
	"fmt"

	"github.com/MinamFaisal/assignment01bca/assignment01bca" //package se functions use krengy
)

func main() {
	// Genesis Block (1st block) iska previous hash kuch ni hga
	genesisBlock := assignment01bca.NewBlock("genesis to Arsh", 1, "")
	fmt.Println("\033[1;32m---------------------\033[0m")
	fmt.Println("\033[1;32mGenesis Block created!\033[0m")
	fmt.Println("\033[1;32m---------------------\033[0m")

	// second block jo pcihly block k hash ko use kr k bnaengy
	block1 := assignment01bca.NewBlock("Arsh to Mahareeb", 2, genesisBlock.Hash)

	// third block jo 1 wale block k hash ko use kr k aega
	assignment01bca.NewBlock("Mahareeb to Maheen", 3, block1.Hash)

	// all blocks print on terminal
	fmt.Println("\n\033[1;34mListing all blocks in the blockchain:\033[0m")
	assignment01bca.ListBlocks()

	//change krrhe 1 block ko
	fmt.Println("\n\033[1;33mChanging the transaction of Block 1...\033[0m")
	assignment01bca.ChangeBlock(1, "Arsh to Mahareeb Updated")
	assignment01bca.ListBlocks()

	// Verify the blockchain
	fmt.Println("\033[1;35m--------------------------------------------\033[0m")
	fmt.Println("\n\033[1;35mVerifying the blockchain after modification:\033[0m")
	fmt.Println("\033[1;35m---------------------------------------------\033[0m")
	assignment01bca.VerifyChain()
}
