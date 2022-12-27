package sockets

import (
	"bufio"
	"fmt"
	"indexer/index"
	"net"
	"strings"
)

func handleRequest(conn net.Conn, idx index.Index){
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		docs := idx.GetItem(strings.TrimSpace(string(netData)))
		fmt.Print("-> ", string(docs))
		conn.Write([]byte(docs + "\n"))
	}
}

func StartServer(idx index.Index) {
	// can be limited easily with channels
	PORT := ":" + "8000"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go handleRequest(conn, idx)
	}
}
