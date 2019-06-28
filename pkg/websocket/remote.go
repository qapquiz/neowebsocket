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

	receivePacketChan chan []byte

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

	r.safeSendToChannel(r.receivePacketChan, p)

	return nil
}

// ReadWorker will read packet from channel
func (r *Remote) ReadWorker(processFunc func(p []byte)) {
	for p := range r.receivePacketChan {
		processFunc(p)
	}
}

// Write will send data to underlying connection
func (r *Remote) Write(p []byte) error {
	wtr := wsutil.NewWriter(r.conn, ws.StateServerSide, ws.OpBinary)

	r.io.Lock()
	defer r.io.Unlock()

	if _, err := wtr.Write(p); err != nil {
		return err
	}

	return wtr.Flush()
}

// Close this connection
func (r *Remote) Close() {
	r.conn.Close()
	r.SafeCloseChannel()
}

// SafeCloseChannel will close channel if error mean channel already close
func (r *Remote) SafeCloseChannel() (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = true
		}
	}()

	close(r.receivePacketChan)

	return true
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

func (r *Remote) safeSendToChannel(ch chan<- []byte, value []byte) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value

	return false
}
