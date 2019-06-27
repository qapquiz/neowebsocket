package main

import (
	"net"
	"time"

	"go.uber.org/zap"
)

func main() {
	var (
		// mobileHub  = websocket.NewHub()
		exit = make(chan struct{})
	)

	<-exit
}

func serveMobile(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		zap.L().Error("error on listen tcp", zap.String("address", addr))
	}

	zap.S().Infof("websocket is listening on %s", ln.Addr().String())

	// @todo #3 add epoll instance to manage this listener
}

func mobileHandler(conn net.Conn) {
	// @todo #4 add upgrader for this handler when mobile connect to this server
	//  we will upgrade from HTTP to Websocket for them. Do not forget to wrap
	//  conn with deadliner!
}

func nameConn(conn net.Conn) string {
	return conn.LocalAddr().String() + " > " + conn.RemoteAddr().String()
}

// deadliner is a wrapper around net.Conn that sets read/write deadlines before
// every Read() or Write() call.
type deadliner struct {
	net.Conn
	t time.Duration
}

func (d deadliner) Write(p []byte) (int, error) {
	if err := d.Conn.SetWriteDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Write(p)
}

func (d deadliner) Read(p []byte) (int, error) {
	if err := d.Conn.SetReadDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Read(p)
}
