package mobile

import (
	"github.com/qapquiz/neowebsocket/pkg/websocket"
	"github.com/qapquiz/packet"
	"go.uber.org/zap"
)

const (
	csLogin uint16 = 10001
)

// PacketProcessor will hold every packet and packet function for mobile client
type PacketProcessor struct {
	mapper map[uint16]func(*websocket.Remote, *packet.Reader)
}

// NewPacketProcessor will hold packet id and
func NewPacketProcessor() PacketProcessor {
	return PacketProcessor{
		mapper: map[uint16]func(*websocket.Remote, *packet.Reader){
			csLogin: receiveLogin,
		},
	}
}

// GetPacketFunc will return packet function that associate with packet id
func (pp PacketProcessor) GetPacketFunc(packetID uint16) func(*websocket.Remote, *packet.Reader) {
	packetFunc, ok := pp.mapper[packetID]
	if !ok {
		zap.S().Error("there is no packetID: %d", packetID)
	}

	return packetFunc
}

func receiveLogin(remote *websocket.Remote, pr *packet.Reader) {

}
