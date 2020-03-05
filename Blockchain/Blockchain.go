package Blockchain

import (
  "fmt"
  "encoding/gob"
  "time"
  b "../Block"
  c1 "../Course"
  p1 "../PersInfo"
  s1 "../Student"
)

var (
  Chain1 Blockchain
)

func RegisterAllGobInterfaces() {
  gob.Register(c1.Course{})
  gob.Register(p1.PersInfo{})
  gob.Register(s1.Student{})
}

type Blockchain struct {
  Head *b.Block
  Tail *b.Block
  NoOfBlocks int
}

func (a *Blockchain) AddBlock(block b.Block) {
	newBlock := block
	newBlock.No = a.NoOfBlocks
  a.NoOfBlocks += 1
	newBlock.Time = time.Now().String()
  fmt.Println(newBlock)

	if a.Head == nil && a.Tail == nil {
		newBlock.PrevHash = "0000000000000000000000000000000000000000000000000000000000000000"
		newBlock.Hash = newBlock.GenerateBlockHash()
		newBlock.Next = nil
		newBlock.Prev = nil
		a.Head = &newBlock
		a.Tail = &newBlock
	} else {
		newBlock.PrevHash = a.Tail.Hash
		newBlock.Hash = newBlock.GenerateBlockHash()
		newBlock.Next = nil
		newBlock.Prev = a.Tail
		a.Tail.Next = &newBlock
		a.Tail = &newBlock
	}

	a.NoOfBlocks += 1
}

func (a *Blockchain) SliceBlockchain() []b.Block {
	nodePtr := a.Tail
	var ls1 []b.Block
	for nodePtr != nil {
		ls1 = append(ls1, *nodePtr)
		nodePtr = nodePtr.Prev
	}
	return ls1
}

func PrintBlockchain(a Blockchain) {
	tempBlock := a.Head
	fmt.Printf("\t\t<=== Blockchain ===>\n")
	for tempBlock != nil {
    fmt.Println(tempBlock)
		tempBlock = tempBlock.Next
	}
}

func (a Blockchain) FindBlock(hash string) bool {
	tempBlock := a.Head
	for tempBlock != nil {
		if tempBlock.Hash == hash {
			return true
		}
		tempBlock = tempBlock.Next
	}
	return false
}

/*

func (a Blockchain) ReversePrintBlockchain() {
	tempBlock := a.Tail
	fmt.Println("<=== Reverse Blockchain ===>")
	for tempBlock != nil {
		tempBlock.PrintBlock()
		tempBlock = tempBlock.Prev
	}
}

func VerifyBlockchain(a Blockchain) bool {
	if a.Head != nil {
		tempBlock := a.Head.Next
		for tempBlock != nil {
			if tempBlock.PrevHash != tempBlock.Prev.GenerateBlockHash() {
				log.Printf("Blockchain was tempered. Security Compromised ☠")
				return false
			}
			tempBlock = tempBlock.Next
		}
	}
	fmt.Println("Blockchain Verified Successfully. No bugs found.")
	return true
}

func (a *Blockchain) GetLastBlock() b1.Block {
	return *a.Tail
}

func (a Blockchain) FindBlock(hash string) bool {
	tempBlock := a.Head
	for tempBlock != nil {
		if tempBlock.Hash == hash {
			return true
		}
		tempBlock = tempBlock.Next
	}
	return false
}
*/
