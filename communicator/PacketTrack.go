package communicator

import "time"

type PacketTrack struct {
	Packet MyPacket
	track  chan bool
}

func (m *PacketTrack) StartTrack() bool {
	m.track = make(chan bool)
	select {
	case <-time.After(time.Millisecond * 1):
		return false
	case <-m.track:
		return true

	}

}

func (m *PacketTrack) Finish(command *MyPacket) {
	if command != nil {
		m.Packet = *command
	}
	m.track <- true
}
