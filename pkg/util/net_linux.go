//go:build linux

package util

import (
	"context"
	"net"
	"syscall"

	"github.com/pingcap/log"
	"go.uber.org/zap"
)

// BindToDevice binds a raw connection to a network device.
func BindToDevice(c syscall.RawConn, device string) error {
	var err error
	var operr error
	err = c.Control(func(fd uintptr) {
		operr = syscall.SetsockoptString(int(fd), syscall.SOL_SOCKET, syscall.SO_BINDTODEVICE, device)
	})
	if err != nil {
		return err
	}
	return operr
}

// CreateDialer creates a dialer that binds to the specified network interface.
func CreateDialer(interfaceName string) *net.Dialer {
	if interfaceName == "" {
		return nil
	}

	dialer := &net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			log.Info("Binding connection to interface", zap.String("interface", interfaceName), zap.String("address", address))
			return BindToDevice(c, interfaceName)
		},
	}
	return dialer
}

// DialContextFunc returns a DialContext function that binds to the specified interface.
func DialContextFunc(interfaceName string) func(ctx context.Context, network, address string) (net.Conn, error) {
	if interfaceName == "" {
		// Use default dialer behavior
		return (&net.Dialer{}).DialContext
	}
	return CreateDialer(interfaceName).DialContext
}
