package communicator

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
)

const (
	ProxyAddr  string = "127.0.0.1:5406"
	ProxyPort  string = ":5406"
	ServerAddr string = "127.0.0.1:5405"
	ServerPort string = ":5405"
	ClientAddr string = "127.0.0.1:5407"
	ClientPort string = ":5407"
)

type ClientAgent struct {
	targetConn    net.Conn
	recvConn      net.PacketConn
	packetManager *ClientPacketManager
}

func NewClientAgent() *ClientAgent {

	agent := ClientAgent{targetConn: nil, recvConn: nil, packetManager: newClientPacketManager()}

	return &agent
}

func (m *ClientAgent) Run() bool {
	var err error = nil
	m.recvConn, err = net.ListenPacket("udp", ClientPort)
	if err != nil {
		fmt.Println(err)
		return false
	}

	for {
		m.targetConn, err = net.Dial("udp", ProxyAddr)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	go m.read()
	return true
}

func (m *ClientAgent) write(data []byte) error {

	var err error
	_, err = m.targetConn.Write(data)
	return err

}

func (m *ClientAgent) read() {
	bs := make([]byte, 256)
	for {

		len, _, err := m.recvConn.ReadFrom(bs)
		if err != nil {
			continue
		}

		go m.onRecived(bs[:len], err)

	}

}

func (m *ClientAgent) onRecived(playload []byte, err error) {
	if err != nil {
		return
	}
	_, err = m.packetManager.ProccessMessage(playload)
	if err != nil {
		return
	}

}

func (m *ClientAgent) Send(packet MyPacket) MyPacket {
	packet.ID = uuid.NewString()
	content, _ := json.Marshal(packet)

	if m.targetConn != nil {
		log.Println(string(content))
		m.write(content)
	}

	track := m.packetManager.RequestPacket(packet)
	result := track.StartTrack()
	if !result {
		m.packetManager.RequestPacket(packet)
	}

	replyPacket := track.Packet
	return replyPacket
}
