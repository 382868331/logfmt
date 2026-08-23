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

func TestGoletaLogfmt005Boundary(t *testing.T) {
	b, e := MarshalKeyvals(map[string]int{"x": 1}, 1, "next", "yes")
	if e != nil || string(b) != "next=yes" {
		t.Fatalf("got=%q e=%v", b, e)
	}
}
