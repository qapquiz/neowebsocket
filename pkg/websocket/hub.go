package websocket

import (
	"net"
	"sync"

	"go.uber.org/zap"
)

// Hub contains all remote that connected
type Hub struct {
	mu  sync.RWMutex
	seq uint

	remotes map[uint]*Remote
}

// NewHub create new hub for every new server
func NewHub() *Hub {
	hub := &Hub{
		remotes: make(map[uint]*Remote),
	}

	return hub
}

// Register registers new connection as a Remote.
func (h *Hub) Register(conn net.Conn) *Remote {
	remote := &Remote{
		conn: conn,
	}

	h.mu.Lock()
	{
		remote.id = h.seq
		h.remotes[remote.id] = remote

		h.seq++
	}
	h.mu.Unlock()

	return remote
}

// Unregister will remove user from chat.
func (h *Hub) Unregister(remote *Remote) {
	h.mu.Lock()
	removed := h.remove(remote)
	h.mu.Unlock()

	if !removed {
		zap.L().Error("cannot unregister remote")
	}
}

func (h *Hub) remove(remote *Remote) bool {
	if _, has := h.remotes[remote.id]; !has {
		return false
	}

	delete(h.remotes, remote.id)

	return true
}
