package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/mailru/easygo/netpoll"
	"github.com/qapquiz/neowebsocket/internal/mobile"
	"github.com/qapquiz/neowebsocket/pkg/websocket"
	"github.com/qapquiz/packet"
	"go.uber.org/zap"
)

var (
	addr      = flag.String("listen", ":8080", "address to bind to")
	ioTimeout = flag.Duration("io_timeout", time.Millisecond*500, "i/o operations timeout")

	poller netpoll.Poller
)

func main() {
	flag.Parse()

	logger := setupZapLogger()
	defer logger.Sync()

	var err error
	poller, err = netpoll.New(nil)
	if err != nil {
		zap.L().Error("cannot create poller", zap.Error(err))
	}

	var (
		mobileHub = websocket.NewHub()
		exit      = make(chan struct{})
	)

	serveMobile(*addr, mobileHub, mobileHandler)

	<-exit
}

func serveMobile(addr string, hub *websocket.Hub, handle func(net.Conn, *websocket.Hub)) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		zap.L().Error("error on listen tcp", zap.String("address", addr))
	}

	zap.S().Infof("websocket is listening on %s", ln.Addr().String())

	acceptDesc := netpoll.Must(netpoll.HandleListener(ln, netpoll.EventRead|netpoll.EventOneShot))

	accept := make(chan error, 1)

	poller.Start(acceptDesc, func(e netpoll.Event) {
		go func() {
			conn, err := ln.Accept()
			if err != nil {
				accept <- err
				return
			}

			accept <- nil
			handle(conn, hub)
		}()

		err = <-accept

		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				goto cooldown
			}

			zap.S().Fatalf("accept error: %v", err)

		cooldown:
			delay := 5 * time.Millisecond
			zap.S().Infof("accept error: %v; retrying in %s", err, delay)
			time.Sleep(delay)
		}

		poller.Resume(acceptDesc)
	})
}

func mobileHandler(conn net.Conn, hub *websocket.Hub) {
	safeConn := deadliner{conn, *ioTimeout}

	hs, err := ws.Upgrade(safeConn)
	if err != nil {
		zap.S().Infof("%s: upgrade error: %v", nameConn(conn), err)
	}

	zap.S().Infof("%s: established websocket connection: %+v", nameConn(conn), hs)

	remote := hub.Register(safeConn)
	packetProcessor := mobile.NewPacketProcessor()

	go remote.ReadWorker(func(p []byte) {
		rdr := packet.NewReader(p)
		packetID := rdr.ReadUInt16()

		fn, err := packetProcessor.GetPacketFunc(packetID)
		if err != nil {
			return
		}

		fn(remote, rdr)
	})

	desc := netpoll.Must(netpoll.HandleRead(conn))

	poller.Start(desc, func(ev netpoll.Event) {
		if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
			poller.Stop(desc)
			remote.SafeCloseChannel()
			hub.Unregister(remote)
			return
		}

		go func() {
			if err := remote.Receive(); err != nil {
				poller.Stop(desc)
				remote.SafeCloseChannel()
				hub.Unregister(remote)
			}
		}()
	})

}

func setupZapLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("cannot get development logger with err: ", err)
	}

	zap.ReplaceGlobals(logger)

	return logger
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
