package sockets

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartClient() {
	c, err := net.Dial("tcp4", "localhost:8000")
	if err != nil {
			fmt.Println(err)
			return
	}

	for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(c, text+"\n")

			message, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message)
			if strings.TrimSpace(string(text)) == "STOP" {
					fmt.Println("TCP client exiting...")
					return
			}
	}
}