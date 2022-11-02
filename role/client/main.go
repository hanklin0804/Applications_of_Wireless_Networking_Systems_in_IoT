package main

import (
	"fmt"
	"log"
	"network-test/communicator"

	"github.com/google/uuid"
)

func main() {
	ch := make(chan bool)
	const datafmt string = "Hello %d"
	client := communicator.NewClientAgent()
	client.Run()
	fmt.Println("Client Start")
	for i := 1; i <= 10000; i++ {
		packet := communicator.MyPacket{Content: fmt.Sprintf(datafmt, i), ID: uuid.NewString()}
		pack := client.Send(packet)
		if pack.Response != "" {
			log.Println(pack.Response)
			// break
		}
		// for {

		// }

	}
	fmt.Println("End")
	<-ch

}
