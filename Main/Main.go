package main

import (
  "fmt"
  "flag"
  n1 "../Network"
  p1 "../PersInfo"
  b1 "../Block"
  c1 "../Blockchain"
)

// IMPORTANT PROBLEMS TO SOLVE

// Correct GenerateBlockHash function in ../Block
// Remove blk.PrevHash = ""


var ownAddr = flag.String("addr", ":10000", "Own Address")
var defaultPeer = flag.String("dpeer", ":10000", "Default Peer")

func main() {
  flag.Parse()

  n1.OwnAddress = *ownAddr
  n1.DefaultPeer = *defaultPeer
  fmt.Println("Own Address:", n1.OwnAddress)
  fmt.Println("Default Peer:", n1.DefaultPeer)

  n1.StartServer()

  input := 0
  for {
    fmt.Println("Choose one of the following scenario:")
    fmt.Println("Enter 1 to view add a new node")
    fmt.Println("Enter 2 to view all known nodes")
    fmt.Println("Enter 3 to add a new PersInfo Block")
    fmt.Println("Enter 4 to view PersChain")
    fmt.Scanln(&input)

    switch input {
    case 1:
      var ip string
      fmt.Println("Enter the new ip")
      fmt.Scanln(&ip)
      n1.AddToKnownNode(ip)
    case 2:
      n1.PrintKnownNodes()
    case 3:
      var input p1.PersInfo
      input.PersInfoInput()
      blk := b1.GenerateBlock("PersInfo", input)
      n1.SendBlock(blk)
    case 4:
      c1.PrintBlockchain(c1.PersInfoChain)
    default:
      fmt.Println("Invalid Command Entered")
    }
  }

}

/*type Fun1 struct {
  A int
}

type Fun2 struct {
  A string
  B float64
}

func main() {
  blk1 := b1.GenerateBlock("Rehan", Fun1{A: 5})
  blk1.PrintBlock()
  blk2 := b1.GenerateBlock("Rehan", Fun1{A: 6})
  blk2.PrintBlock()
  blk3 := b1.GenerateBlock("Rehan", Fun2{A: "Alhamdulillah", B: 4.7})
  blk3.PrintBlock()
}*/
