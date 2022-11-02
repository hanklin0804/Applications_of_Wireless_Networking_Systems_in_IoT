package communicator

import "time"

type PacketTrack struct {
	Packet MyPacket
	track  chan bool
}

func (m *PacketTrack) StartTrack() bool {
	m.track = make(chan bool)
	select {
	case <-m.track:
		return true
	case <-time.After(time.Millisecond * 10):
		return false
	}

}

func (m *PacketTrack) Finish(command *MyPacket) {
	if command != nil {
		m.Packet = *command
	}
	m.track <- true
}
