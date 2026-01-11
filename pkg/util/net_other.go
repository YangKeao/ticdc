//go:build !linux

package util

import (
	"context"
	"net"
	"syscall"
)

// BindToDevice binds a raw connection to a network device.
func BindToDevice(c syscall.RawConn, device string) error {
	return nil
}

// CreateDialer creates a dialer that binds to the specified network interface.
func CreateDialer(interfaceName string) *net.Dialer {
	return &net.Dialer{}
}

// DialContextFunc returns a DialContext function that binds to the specified interface.
func DialContextFunc(interfaceName string) func(ctx context.Context, network, address string) (net.Conn, error) {
	return (&net.Dialer{}).DialContext
}
