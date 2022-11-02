package main

import (
	"fmt"
	"network-test/communicator"
)

func main() {
	ch := make(chan bool)
	proxy := communicator.NewProxyAgent()
	proxy.Run()
	fmt.Println("Proxy Start")
	<-ch
}
