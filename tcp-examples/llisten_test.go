package tcpexamples

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listenr, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = listenr.Close()
	}()

	t.Logf("bound to %q", listenr.Addr())
}
