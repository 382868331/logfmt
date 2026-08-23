package logfmt

import (
	"bytes"
	"testing"
)

func TestGoletaLogfmt015(t *testing.T) {
	var b bytes.Buffer
	e := NewEncoder(&b)
	_ = e.EncodeKeyval("a", 1)
	e.Reset()
	b.Reset()
	_ = e.EncodeKeyval("b", 2)
	if b.String() != "b=2" {
		t.Fatalf("got=%q", b.String())
	}
}
