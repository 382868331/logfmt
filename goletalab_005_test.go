package logfmt

import (
	"bytes"
	"testing"
)

func TestGoletaLogfmt005(t *testing.T) {
	var b bytes.Buffer
	e := NewEncoder(&b)
	err := e.EncodeKeyvals([]string{"bad"}, 1, "ok", 2)
	if err != nil || b.String() != "ok=2" {
		t.Fatalf("got=%q err=%v", b.String(), err)
	}
}
