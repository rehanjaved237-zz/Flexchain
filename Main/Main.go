package main

import (
  "fmt"
  "flag"
  "log"
  "net/http"
  "html/template"
  n1 "../Network"
  p1 "../PersInfo"
  b1 "../Block"
  c1 "../Blockchain"
  w1 "../Web"
  buff1 "../BlockBuffer"
)

// IMPORTANT PROBLEMS TO SOLVE

// Correct GenerateBlockHash function in ../Block
// Remove blk.PrevHash = ""
// Blockchain printing issue


var ownAddr = flag.String("addr", ":10000", "Own Address")
var defaultPeer = flag.String("dpeer", ":10000", "Default Peer")
var frontEnd = flag.String("fend", ":12000", "Front End")

var (
	templates = template.Must(template.ParseFiles("index.html", "myHtml.html", "Home.html", "Login.html"))
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "Login.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "Home.html", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func myHtmlHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "myHtml.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func runHandlers()  {

  http.HandleFunc("/", DefaultHandler)
  http.HandleFunc("/Home/", HomeHandler)
  http.HandleFunc("/Login/", LoginHandler)
  http.HandleFunc("/myHtml/", myHtmlHandler)

  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  log.Fatal(http.ListenAndServe(w1.FrontEnd , nil))
}

func main() {
  flag.Parse()

  n1.OwnAddress = *ownAddr
  n1.DefaultPeer = *defaultPeer
  w1.FrontEnd = *frontEnd


  fmt.Println("Own Address:", n1.OwnAddress)
  fmt.Println("Default Peer:", n1.DefaultPeer)
  fmt.Println("Front End:", w1.FrontEnd)

  go runHandlers()
  n1.StartServer()

  input := 0
  for {
    fmt.Println("Choose one of the following scenario:")
    fmt.Println("Enter 1 to view add a new node")
    fmt.Println("Enter 2 to view all known nodes")
    fmt.Println("Enter 3 to add a new PersInfo Block")
    fmt.Println("Enter 4 to view PersChain")
    fmt.Println("Enter 5 to view the BlockBuffer")
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
//      blk.PrintBlock()
      n1.BroadCastBlock(blk)
    case 4:
      c1.PrintBlockchain(c1.PersInfoChain)
    case 5:
      buff1.BlkBuffer.PrintBlockBuffer()
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
