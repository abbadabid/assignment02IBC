package main

import (
	"fmt"

	a2 "github.com/abbadabid/assignment02IBC"
)

func main() {
	var chainHead *a2.Block
	//This insertion is invalid as Alice is neither miner nor has enough coins for the transaction, pay 50 from Alice to Bob
	aliceToBob := []a2.BlockData{{Title: "ALice2Bob", Sender: "Alice", Receiver: "Bob", Amount: 50}}
	chainHead = a2.InsertBlock(aliceToBob, chainHead)

	//Lets mine some blocks to start the chain and check Satoshi's balance
	chainHead = a2.PremineChain(chainHead, 2)
	//fmt.Printf("Satoshi's balance: %v\n", a2.CalculateBalance("Satoshi", chainHead))

	//Now Satoshi can send some coins to Alice
	SatoshiToAlice := []a2.BlockData{{Title: "SatoshiToAlice", Sender: "Satoshi", Receiver: "Alice", Amount: 50}}
	chainHead = a2.InsertBlock(SatoshiToAlice, chainHead)

	//We can verify this by checking balances once again and listing the chain
	/*
	   fmt.Printf("Satoshi's balance: %v\n", a2.CalculateBalance("Satoshi", chainHead))
	   fmt.Printf("Alice's balance: %v\n", a2.CalculateBalance("Alice", chainHead))
	   a2.ListBlocks(chainHead)
	*/
	//Alice can then make the transactions using her coins, She can make multiple
	//transactions at once, notice that field Data has type []BlockData in Block Struct
	AliceToBobCharlie := []a2.BlockData{{Title: "ALice2Bob", Sender: "Alice", Receiver: "Bob", Amount: 20}, {Title: "ALice2Charlie", Sender: "Alice", Receiver: "Charlie", Amount: 10}}
	chainHead = a2.InsertBlock(AliceToBobCharlie, chainHead)

	//We can verify this by checking balances once again and listing the chain
	a2.ListBlocks(chainHead)

	fmt.Printf("Satoshi's balance: %v ", a2.CalculateBalance("Satoshi", chainHead))
	fmt.Printf("Alice's balance: %v ", a2.CalculateBalance("Alice", chainHead))
	fmt.Printf("Charlie's balance: %v\n", a2.CalculateBalance("Charlie", chainHead))

	//Finally the transaction verification fails if any of the transaction is invalid
	oneInvalidoneValid := []a2.BlockData{{Title: "ALice2EZ", Sender: "Alice", Receiver: "Bob", Amount: 100}, {Title: "Satoshi2EZ", Sender: "Satoshi", Receiver: "EZ", Amount: 200}}
	chainHead = a2.InsertBlock(oneInvalidoneValid, chainHead)

	//Bonus (2 absolutes) - Fix the erroneous behavior below
	//The transactions are valid individually but when applied to chain Alice's balance
	//become negative :(
	bonusTransactions := []a2.BlockData{{Title: "ALice2Bob", Sender: "Alice", Receiver: "Bob", Amount: 15}, {Title: "AliceToEZ", Sender: "Alice", Receiver: "EZ", Amount: 15}}
	chainHead = a2.InsertBlock(bonusTransactions, chainHead)
	fmt.Printf("Alice's balance: %v\n", a2.CalculateBalance("Alice", chainHead))

}
