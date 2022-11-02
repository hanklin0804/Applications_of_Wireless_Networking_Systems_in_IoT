package communicator

import (
	"encoding/json"
	"fmt"
	"net"
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

	for {
		m.targetConn, err = net.Dial("udp", ClientAddr)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	go m.read()
	return true
}

func (m *ServerAgent) write(data []byte) error {

	var err error
	_, err = m.targetConn.Write(data)
	return err

}

func (m *ServerAgent) read() {
	// bs := make([]byte, 1024)
	bs := make([]byte, 256)
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
