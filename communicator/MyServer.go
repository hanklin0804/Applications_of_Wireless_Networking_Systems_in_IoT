package communicator

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type ServerAgent struct {
	targetConn    net.Conn
	recvConn      net.PacketConn
	packetManager *ServerPacketManager
}

func NewServerAgent() *ServerAgent {

	agent := ServerAgent{targetConn: nil, recvConn: nil, packetManager: newServerPacketManager()}

	return &agent
}

func (m *ServerAgent) Run() bool {
	var err error
	m.recvConn, _ = net.ListenPacket("udp", ServerPort)
	if err != nil {
		return false
	}
	for !m.reconnect() {
		fmt.Println("reconnect")
	}

	go m.read()
	return true
}

func (m *ServerAgent) reconnect() bool {
	var err error = nil
	if m.targetConn != nil {
		m.targetConn.Close()
		m.targetConn, err = net.Dial("udp", ClientAddr)
	} else {
		time.Sleep(200 * time.Millisecond)
	}
	return err == nil

}

func (m *ServerAgent) write(data []byte) {

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

func (m *ServerAgent) read() {
	bs := make([]byte, 1024)
	for {
		len, _, err := m.recvConn.ReadFrom(bs)
		if err != nil {
			continue
		}

		go m.onRecived(bs[:len], err)

	}

}

func (m *ServerAgent) onRecived(playload []byte, err error) {
	if err != nil {
		return
	}
	packet, err := m.packetManager.ProccessMessage(playload)
	if err != nil {
		return
	}

	go m.response(*packet)
}

func (m *ServerAgent) response(packet MyPacket) {

	content, _ := json.Marshal(packet)

	if m.targetConn != nil {
		m.write(content)
	}

}
