package communicator

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ServerPacketManager struct {
}

func newServerPacketManager() *ServerPacketManager {
	return &ServerPacketManager{}
}
func (m *ServerPacketManager) ProccessMessage(jsonString []byte) (*MyPacket, error) {
	pack := MyPacket{}
	if err := json.Unmarshal(jsonString, &pack); err != nil {
		fmt.Println("format error")
		return nil, err
	}

	return m.proccessCommand(&pack)
}

func (m *ServerPacketManager) proccessCommand(packet *MyPacket) (*MyPacket, error) {

	var result *MyPacket = nil
	s := strings.Split(packet.Content, " ")
	packet.Response = "World" + " " + s[1]
	result = packet
	return result, nil
}
