package logfmt

import (
	"bytes"
	"testing"
)

func TestGoletaLogfmt004(t *testing.T) {
	var b bytes.Buffer
	e := NewEncoder(&b)
	if err := e.EncodeKeyvals("a"); err != nil || b.String() != "a=null" {
		t.Fatalf("got=%q err=%v", b.String(), err)
	}
}

func TestGoletaLogfmt004Boundary(t *testing.T) {
	b, e := MarshalKeyvals("a", 1, "b")
	if e != nil || string(b) != "a=1 b=null" {
		t.Fatalf("got=%q e=%v", b, e)
	}
}
