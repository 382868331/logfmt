package logfmt

import (
	"bytes"
	"testing"
)

func TestGoletaLogfmt002(t *testing.T) {
	var b bytes.Buffer
	e := NewEncoder(&b)
	_ = e.EncodeKeyval("a", 1)
	_ = e.EncodeKeyval("b", 2)
	if b.String() != "a=1 b=2" {
		t.Fatalf("got=%q", b.String())
	}
}
