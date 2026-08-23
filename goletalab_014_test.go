package logfmt

import (
	"bytes"
	"testing"
)

func TestGoletaLogfmt014(t *testing.T) {
	var b bytes.Buffer
	e := NewEncoder(&b)
	_ = e.EncodeKeyval("a", 1)
	_ = e.EndRecord()
	if b.String() != string([]byte{97, 61, 49, 10}) {
		t.Fatalf("got=%q", b.String())
	}
}
