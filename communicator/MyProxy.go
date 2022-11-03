package communicator

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
)

type ProxyAgent struct {
	targetConn net.Conn
	recvConn   net.PacketConn
}

func rollEvent(isEven bool) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20)
	if (num == 0) && !isEven { // delay 100ms
		return 0
	} else if (num == 1 || num == 2) && isEven { //drop
		return 1
	} else {
		return 2
	}
}

func NewProxyAgent() *ProxyAgent {

	agent := ProxyAgent{targetConn: nil, recvConn: nil}

	return &agent
}

func (m *ProxyAgent) Run() bool {
	var err error = nil
	m.recvConn, err = net.ListenPacket("udp", ProxyPort)
	if err != nil {
		return false
	}

	for {
		m.targetConn, err = net.Dial("udp", ServerAddr)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	go m.read()
	return true
}

func (m *ProxyAgent) write(data []byte) error {

	_, err := m.targetConn.Write(data)

	return err

}

func (m *ProxyAgent) read() {
	bs := make([]byte, 256)
	for {

		len, _, err := m.recvConn.ReadFrom(bs)
		if err != nil {
			continue
		}

		go m.onRecived(bs[:len], err)

	}

}

func (m *ProxyAgent) onRecived(playload []byte, err error) {
	if err != nil {
		return
	}
	contentText, _ := jsonparser.GetString(playload, "Content")
	num, _ := strconv.Atoi(strings.Split(contentText, " ")[1])
	switch rollEvent(num%2 == 0) {
	case 0:
		time.Sleep(100 * time.Millisecond)
		m.write(playload)
		// log.Println(string(playload))
	case 1:
		// return
		// m.write(playload)
		// log.Println(string(playload))
	case 2:
		m.write(playload)
		log.Println(string(playload))
	}

}
