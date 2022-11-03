package main

import (
	"fmt"
	"log"
	"network-test/communicator"
)

func main() {
	ch := make(chan bool)
	const datafmt string = "Hello %d"
	client := communicator.NewClientAgent()
	client.Run()
	fmt.Println("Client Start")
	for i := 1; i <= 10000; i++ {
		packet := communicator.MyPacket{Content: fmt.Sprintf(datafmt, i)}

		for {
			pack := client.Send(packet)
			if pack.Response != "" {
				log.Println(pack.Response)
				break
			}
		}

	}
	fmt.Println("End")
	<-ch

}
