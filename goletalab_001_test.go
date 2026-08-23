package logfmt

import (
	"testing"
)

func TestGoletaLogfmt001(t *testing.T) {
	b, e := MarshalKeyvals("a", 1)
	if e != nil || string(b) != "a=1" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}

func TestGoletaLogfmt001Boundary(t *testing.T) {
	b, e := MarshalKeyvals("msg", "hello")
	if e != nil || string(b) != "msg=hello" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}
