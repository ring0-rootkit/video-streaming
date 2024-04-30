package live_server

import (
	"testing"
)

const (
	port uint16 = 42069
)

func TestServerStart(t *testing.T) {
	srv := New(port)
	quitChan := srv.Start()

	if srv.sck == nil {
		t.Error("srv.sck == nil failed, expected false")
	}
	if srv.readWriter == nil {
		t.Error("srv.readWriter == nil failed, expected false")
	}
	if srv.port != port {
		t.Error("srv.port != port failed, expected false")
	}
	if srv.isRunning != true {
		t.Error("srv.isRunning != true failed, expected false")
	}
	if srv.peers == nil {
		t.Error("srv.peers == nil failed, expected false")
	}
	if srv.qc != quitChan {
		t.Error("srv.qc != quitChan failed, expected false")
	}
}
