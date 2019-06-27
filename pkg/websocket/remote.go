package websocket

import (
	"io"
	"io/ioutil"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Remote is a struct that hold information about connection
type Remote struct {
	id uint

	conn io.ReadWriteCloser

	io sync.Mutex
}

// Receive reads next message from user's underlying connection.
// It blocks until full message received.
func (r *Remote) Receive() error {
	p, err := r.readRequest()
	if err != nil {
		r.conn.Close()
		return err
	}

	if p == nil {
		// handled some control message
		return nil
	}

	// @todo #1 extract packet and call packetFunc

	return nil
}

// readRequest reads bytes from connection
// It will lock connection before read
func (r *Remote) readRequest() ([]byte, error) {
	r.io.Lock()
	defer r.io.Unlock()

	hdr, rdr, err := wsutil.NextReader(r.conn, ws.StateServerSide)
	if err != nil {
		return nil, err
	}

	if hdr.OpCode.IsControl() {
		wsutil.ControlFrameHandler(r.conn, ws.StateServerSide)(hdr, rdr)
	}

	p, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *Remote) write(p []byte) error {
	wtr := wsutil.NewWriter(r.conn, ws.StateServerSide, ws.OpBinary)

	r.io.Lock()
	defer r.io.Unlock()

	// @todo #2 write packet in properly format with packetWriter

	return wtr.Flush()
}
