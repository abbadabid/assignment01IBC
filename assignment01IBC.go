package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}
type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	s := fmt.Sprintf("%v", inputBlock)
	return fmt.Sprintf("%x\n", sha256.Sum256([]byte(s)))
}

func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	if chainHead == nil {
		chainHead = new(Block)
		chainHead.PrevPointer = nil
		chainHead.Data = dataToInsert
		chainHead.PrevHash = ""
		chainHead.CurrentHash = CalculateHash(chainHead)
	} else {
		var temp *Block = new(Block)
		temp.PrevPointer = chainHead
		temp.Data = dataToInsert
		temp.PrevHash = chainHead.CurrentHash
		temp.CurrentHash = CalculateHash(temp)
		chainHead = temp
	}
	return chainHead
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	var temp *Block = chainHead
	for temp != nil {
		for j := 0; j < len(temp.Data.Transactions); j++ {
			if temp.Data.Transactions[j] == oldTrans {
				temp.Data.Transactions[j] = newTrans
			}
		}
		temp = temp.PrevPointer
	}

}
func ListBlocks(chainHead *Block) {
	var i int = 1
	var head *Block = chainHead
	for head != nil {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("CurrentHash %v\n", head.CurrentHash)
		fmt.Printf("PreviousHash %v\n", head.PrevHash)
		fmt.Printf("Data %v\n\n", head.Data)
		head = head.PrevPointer
		i++
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
