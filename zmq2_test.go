package zmq2_test

import (
	zmq "github.com/pebbe/zmq2"

	"testing"
)

func TestVersion(t *testing.T) {
	major, _, _ := zmq.Version()
	if major != 2 {
		t.Errorf("Expected major version 2, got %d", major)
	}
}
