package utility

import "testing"

func TestBinaryPath(t *testing.T) {
	p, err := BinaryPath()
	if err != nil {
		t.Error(err)
	}
	t.Log(p)
}
