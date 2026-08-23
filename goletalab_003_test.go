package logfmt

import (
	"bytes"
	"testing"
)

func TestGoletaLogfmt003(t *testing.T) {
	var b bytes.Buffer
	if e := NewEncoder(&b).EncodeKeyval("a", 1); e != nil || b.String() != "a=1" {
		t.Fatalf("got=%q e=%v", b.String(), e)
	}
}

func TestGoletaLogfmt003Boundary(t *testing.T) {
	var b bytes.Buffer
	if e := NewEncoder(&b).EncodeKeyval("key", "value"); e != nil || b.String() != "key=value" {
		t.Fatalf("got=%q e=%v", b.String(), e)
	}
}
