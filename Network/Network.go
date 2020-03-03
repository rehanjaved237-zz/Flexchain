package Network

import (
	"fmt"
	"log"
	"net"
	"os"
//	"sync"
  "bytes"
  "strconv"
  "encoding/gob"
  b1 "../Block"
  c1 "../Blockchain"
	buff1 "../BlockBuffer"
)

const (
	CommandLength = 12
	Network       = "tcp"
)

var (
	OwnAddress  string
	DefaultPeer string
	KnownNodes  = map[string]string{}
)

type Addr struct {
	AddrList    []string
	DefaultPeer string
}

type BlockSender struct {
	BlockList []b1.Block
}

func StartServer() {
  c1.RegisterAllGobInterfaces()

	conn, err := net.Listen(Network, OwnAddress)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	go turnOnServer(conn)

	fmt.Println("Server has Started ...")
	AddToKnownNode(OwnAddress)
	AddToKnownNode(DefaultPeer)

	AskNodes()
}

// FUNCTIONS UNDER CONSTRUCTION BEGINS...

func HandleBlock(conn net.Conn, data []byte) {

	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var blkList BlockSender
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&blkList)
	if err != nil {
		log.Println(err)
	}

	for i, blk := range blkList.BlockList {
		hash := buff1.GenerateHash(blk)
		found, _ := buff1.BlkBuffer.FindBlock(hash)
		fmt.Println(found, i)
		if !found {
			fmt.Println(hash)
			buff1.BlkBuffer.InsertBlock(blk)
			BroadCastBlock(blk)
		}
	}
}

func BroadCastBlock(block b1.Block) {
	for _, address := range KnownNodes {
		go SendBlock(address, block)
	}
}

func SendBlock(address string, block b1.Block) {
	blk := BlockSender{BlockList: []b1.Block{block}}
	byteBlk := GobEncode(blk)
	data := append(CmdToBytes("block"), byteBlk...)

	SendData(address, data)
}

// FUNCTIONS UNDER CONSTRUCTION ENDS...











func turnOnServer(conn net.Listener) {
	for {
		ln, err := conn.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConnection(ln)
	}
}

func handleConnection(conn net.Conn) {
	data := make([]byte, 8192)
	n, err := conn.Read(data)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	fmt.Printf("Successfully read %d bytes from client\n", n)
	cmd := BytesToCmd(data[:CommandLength])

	switch cmd {
	case "addr":
		HandleAddr(conn, data[CommandLength:])
  case "block":
    HandleBlock(conn, data[CommandLength:])
  case "askaddr":
		HandleAskAddress(conn, data[CommandLength:])
	default:
		fmt.Println("Unknown Command Found")
	}
}

func AskNodes() {
	keys := []string{OwnAddress}
	addr := Addr{AddrList: keys}

	byteAddr := GobEncode(addr)

	data := append(CmdToBytes("askaddr"), byteAddr...)
	for _, v := range KnownNodes {
		go SendData(v, data)
	}
}

func HandleAskAddress(conn net.Conn, data []byte) {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var address Addr
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&address)
	if err != nil {
		log.Println(err)
	}

	AddToKnownNode(address.AddrList[0])
	BroadCastNodes()
}

func HandleAddr(conn net.Conn, data []byte) {
	var buff bytes.Buffer
	_, err := buff.Write(data)
	if err != nil {
		log.Println(err)
	}

	var address Addr
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&address)
	if err != nil {
		log.Println(err)
	}

	broadcast := false
	for _, t := range address.AddrList {
		_, found := KnownNodes[t]
		if !found {
			KnownNodes[t] = t
			broadcast = true
		}
	}

	if broadcast == true {
		BroadCastNodes()
	}
}

func SendNodes(address string) {
	keys := []string{}
	for node, _ := range KnownNodes {
		keys = append(keys, node)
	}

	addr := Addr{AddrList: keys}
	addr.AddrList = append(addr.AddrList, OwnAddress)
	byteAddr := GobEncode(addr)
	newAddr := append(CmdToBytes("addr"), byteAddr...)

	SendData(address, newAddr)
}

func SendData(address string, data []byte) {
	if address == OwnAddress {
		return
	}

	conn, err := net.Dial(Network, address)
	if err != nil {
		log.Printf("Error: %s\n", err)

		delete(KnownNodes, address)
		return
	}
	//  n, err := conn.Write(data)
	n, err := conn.Write(data)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	fmt.Printf("Written %d bytes successfully on %s\n", n, address)
}

func BroadCastNodes() {
	for address, _ := range KnownNodes {
		go SendNodes(address)
	}
}

func PrintKnownNodes() {
	k := 0
	for _, node := range KnownNodes {
		fmt.Println(strconv.Itoa(k)+".", node)
		k += 1
	}
}

func GobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Printf("Error5: %f\n", err)
	}

	return buff.Bytes()
}

func CmdToBytes(cmdString string) []byte {
	var bytes [CommandLength]byte

	for i, c := range cmdString {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func BytesToCmd(byteCmd []byte) string {
	var cmd []byte

	for _, c := range byteCmd {
		if c != 0x0 {
			cmd = append(cmd, c)
		}
	}

	return fmt.Sprintf("%s", cmd)
}

func AddToKnownNode(addr string) {
	KnownNodes[addr] = addr
}
