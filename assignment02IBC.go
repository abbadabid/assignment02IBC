package assignment02IBC

import (
	"crypto/sha256"
	"fmt"
)

//a2 "github.com/ehteshamz/assignment02IBC"

const miningReward = 100
const rootUser = "Satoshi"

type BlockData struct {
	Title    string
	Sender   string
	Receiver string
	Amount   int
}

type Block struct {
	Data        []BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {
	sum := 0
	temp := chainHead

	for temp != nil {
		for _, v := range temp.Data {
			if v.Receiver == userName {
				sum = sum + v.Amount
			}
		}
		temp = temp.PrevPointer
	}

	temp = chainHead

	for temp != nil {
		for _, v := range temp.Data {
			if v.Sender == userName {
				sum = sum - v.Amount
			}
		}
		temp = temp.PrevPointer
	}

	return sum
}

func PrintSlice(s BlockData) {
	fmt.Printf("Title=%v Sender=%v Receiver=%v Amount=%d\n", s.Title, s.Sender, s.Receiver, s.Amount)
}

func CalculateHash(inputBlock *Block) string {
	s := fmt.Sprintf("%v", inputBlock)
	return fmt.Sprintf("%x\n", sha256.Sum256([]byte(s)))
}

func VerifyTransaction(transaction []BlockData, chainHead *Block) bool {

	amount := 0
	check := true
	var str[]string

	for index := 0; index < len(transaction); index++ {
		sender_name := transaction[index].Sender
		//sum := 0
		//temp := chainHead
		amount = 0
		balance := CalculateBalance(transaction[index].Sender, chainHead)
		//fmt.Printf("Balance%d:", balance)

		for k := 0; k < len(transaction); k++ {
			if k != index {
				if sender_name == transaction[k].Sender {
					amount = amount - transaction[k].Amount
				}

				if sender_name == transaction[k].Receiver {
					amount = amount + transaction[k].Amount
				}

			}
			//sum = amount
			//amount = amount - transaction[index].Amount
		}
		balance = balance + amount
		balance = balance - transaction[index].Amount
		//fmt.Printf("amount%v", amount)
		if balance >= 0 && check != false {
			check = true
		} else {
			check = false
		}

		c := false

		for i:=0; i<len(str) ; i++{
			if str[i] == sender_name{
				c = true
				break
			}
		}

		if balance < 0 && c == false{
			fmt.Printf("Error %v has %d coins - %d  were needed\n", transaction[index].Sender, CalculateBalance(transaction[index].Sender, chainHead), transaction[index].Amount-amount)
			str = append(str,sender_name)
		}

	}
	return check
}

func InsertBlock(blockData []BlockData, chainHead *Block) *Block {

	if VerifyTransaction(blockData, chainHead) == true {
		if chainHead == nil {
			chainHead = new(Block)
			chainHead.PrevHash = ""
			chainHead.CurrentHash = CalculateHash(chainHead)
			chainHead.PrevPointer = nil
			chainHead.Data = append(chainHead.Data, BlockData{Title: "Coinbase Sender", Sender: "System", Receiver: rootUser, Amount: miningReward})
		} else {
			temp := new(Block)
			temp.Data = blockData
			temp.CurrentHash = CalculateHash(temp)
			temp.PrevHash = temp.CurrentHash
			temp.Data = append(temp.Data, BlockData{Title: "Coinbase Sender", Sender: "System", Receiver: rootUser, Amount: miningReward})
			temp.PrevPointer = chainHead
			chainHead = temp
		}
	}
	return chainHead
}

func ListBlocks(chainHead *Block) {

	var i int = 1
	var head *Block = chainHead
	for head != nil {
		fmt.Printf("Block %d\n", i)
		//fmt.Printf("CurrentHash %v\n", head.CurrentHash)
		//fmt.Printf("PreviousHash %v\n", head.PrevHash)
		for _, i := range head.Data {
			PrintSlice(i)
		}
		head = head.PrevPointer
		i++
		fmt.Printf("\n\n")
	}

}

func VerifyChain(chainHead *Block) {
	t := true
	for chainHead.PrevPointer != nil {
		if chainHead.PrevHash != chainHead.PrevPointer.CurrentHash {
			t = false
			break
		}
		chainHead = chainHead.PrevPointer
	}
	if t == true {
		println("chain verified")
	} else {
		println("chain not smooth")
	}
}

//Sprintf("%v")

//-------------------------Premine chain
func PremineChain(chainHead *Block, numBlocks int) *Block {
	if chainHead == nil {
		for i := 0; i < numBlocks; i++ {
			block := []BlockData{{Title: "Premined", Sender: "nil", Receiver: "nil", Amount: 0}}
			chainHead = InsertBlock(block, chainHead)
		}
	} else {
		fmt.Print("Invalid Premine of block")
	}
	return chainHead
}
