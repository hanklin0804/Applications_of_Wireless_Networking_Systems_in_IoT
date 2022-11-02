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
		packet = client.Send(packet)
		if packet.Response != "" {
			log.Println(packet.Response)
			break
		}
		// for {

		// }

	}
	fmt.Println("End")
	<-ch
}
