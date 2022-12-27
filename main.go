package main

import (
	"fmt"
	"indexer/index"
	"indexer/sockets"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// things need to add - debug mode(with an extra prints)
	// print all the keys
	createdIndex := index.BuildIndex([]string{"data"}, 10000)
	createdIndex.Display()
	fmt.Println("Finished building inverted index")

	sockets.StartServer(createdIndex)
}
