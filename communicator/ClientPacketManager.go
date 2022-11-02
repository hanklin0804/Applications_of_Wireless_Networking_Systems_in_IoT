package communicator

import (
	"encoding/json"
	"fmt"
	"sync"
)

type ClientPacketManager struct {
	commandMap map[string]*PacketTrack
	mux        sync.Mutex
}

func newClientPacketManager() *ClientPacketManager {
	return &ClientPacketManager{commandMap: make(map[string]*PacketTrack)}
}
func (m *ClientPacketManager) ProccessMessage(jsonString []byte) (*MyPacket, error) {
	cmd := MyPacket{}
	if err := json.Unmarshal(jsonString, &cmd); err != nil {
		fmt.Println("format error")
		return nil, err

	}
	return m.proccessCommand(&cmd)
}

func (m *ClientPacketManager) proccessCommand(packet *MyPacket) (*MyPacket, error) {

	var result *MyPacket = nil
	pack := MyPacket{}
	m.mux.Lock()
	defer m.mux.Unlock()
	if val, isExisted := m.commandMap[packet.ID]; isExisted {
		val.Finish(packet)
		delete(m.commandMap, packet.ID) //對應的Command會被刪掉
		pack = val.Packet               //因此我們先把該指令內容複製到暫時的容器裡
		result = &pack
	}
	return result, nil
}

//發出request的tracker，在對方reply前在呼叫一次會將response忽略
func (m *ClientPacketManager) RequestPacket(packet MyPacket) *PacketTrack {
	m.mux.Lock()
	defer m.mux.Unlock()
	if val, isExisted := m.commandMap[packet.ID]; isExisted {
		val.Finish(&packet)
		delete(m.commandMap, packet.ID)
	}

	if _, isExisted := m.commandMap[packet.ID]; !isExisted {
		tempTrack := new(PacketTrack)
		m.commandMap[packet.ID] = tempTrack
		return tempTrack
	}
	return nil
}