package communicator

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type ProxyAgent struct {
	targetConn net.Conn
	recvConn   net.PacketConn
}

func rollEvent() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20)
	if num == 0 { // delay 100ms
		return 0
	} else if num == 1 || num == 2 { //drop
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
	m.recvConn, _ = net.ListenPacket("udp", ProxyPort)
	for !m.reconnect() {
		fmt.Println("reconnect")
	}

	go m.read()
	return true
}
func (m *ProxyAgent) reconnect() bool {
	var err error = nil
	if m.targetConn != nil {
		m.targetConn.Close()
		m.targetConn, err = net.Dial("udp", ServerAddr)
	}
	if err != nil {
		time.Sleep(200 * time.Millisecond)
		return false
	} else {
		return true
	}

}

func (m *ProxyAgent) write(data []byte) {

	var err error
	for {
		_, err = m.targetConn.Write(data)

		if err != nil {

			for !m.reconnect() {
				fmt.Println("reconnect")
			}

		} else {
			break
		}
	}

}

func (m *ProxyAgent) read() {
	bs := make([]byte, 1024)
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
	log.Println(string(playload))
	switch rollEvent() {
	case 0:
		time.Sleep(100 * time.Millisecond)
	case 1:
		return
	case 2:
		m.write(playload)
	}

}
