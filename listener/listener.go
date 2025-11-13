package listener

import (
	"net"
	"os"
)

func NewSocket(socketDir string, pattern string) (net.Listener, error) {
	socket, err := os.CreateTemp(socketDir, pattern)
	if err != nil {
		return nil, err
	}

	path := socket.Name()
	if err := socket.Close(); err != nil {
		return nil, err
	}
	if err := os.Remove(path); err != nil {
		return nil, err
	}

	return net.Listen("unix", path)
}
