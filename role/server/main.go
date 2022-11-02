package main

import (
	"fmt"
	"network-test/communicator"
)

func main() {
	ch := make(chan bool)
	server := communicator.NewServerAgent()
	server.Run()
	fmt.Println("Server Start")
	<-ch
}
