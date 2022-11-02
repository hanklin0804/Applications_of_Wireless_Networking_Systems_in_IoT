package communicator

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

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
	m.recvConn, _ = net.ListenPacket("udp", ClientPort)
	for !m.reconnect() {
		fmt.Println("reconnect")
	}
	go m.read()
	return true
}
func (m *ClientAgent) reconnect() bool {
	var err error = nil
	if m.targetConn != nil {
		m.targetConn.Close()
		m.targetConn, err = net.Dial("udp", ProxyAddr)
	}
	if err != nil {
		time.Sleep(200 * time.Millisecond)
		return false
	} else {
		return true
	}

}
func (m *ClientAgent) write(data []byte) {

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

func (m *ClientAgent) read() {
	bs := make([]byte, 1024)
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
		m.write(content)
		// fmt.Println("aaa")
	}

	track := m.packetManager.RequestPacket(packet)
	var replied bool
	if replied = track.StartTrack(); replied {
		m.packetManager.RequestPacket(packet)
	}
	replyPacket := track.Packet
	return replyPacket
}
