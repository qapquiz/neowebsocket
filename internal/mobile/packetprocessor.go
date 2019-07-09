package mobile

import (
	"errors"

	"github.com/qapquiz/neowebsocket/pkg/websocket"
	"go.uber.org/zap"
)

const (
	csLogin uint16 = 10001
)

// PacketProcessor will hold every packet and packet function for mobile client
type PacketProcessor struct {
	mapper map[uint16]func(*websocket.Remote, []byte)
}

// NewPacketProcessor will hold packet id and
func NewPacketProcessor() PacketProcessor {
	return PacketProcessor{
		mapper: map[uint16]func(*websocket.Remote, []byte){
			csLogin: receiveLogin,
		},
	}
}

// GetPacketFunc will return packet function that associate with packet id
func (pp PacketProcessor) GetPacketFunc(packetID uint16) (func(*websocket.Remote, []byte), error) {
	packetFunc, ok := pp.mapper[packetID]
	if !ok {
		zap.S().Errorf("there is no packetID: %d", packetID)
		return nil, errors.New("there is no packetID")
	}

	return packetFunc, nil
}

func receiveLogin(remote *websocket.Remote, p []byte) {

}
