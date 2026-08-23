package logfmt

import (
	"testing"
)

func TestGoletaLogfmt013(t *testing.T) {
	b, e := MarshalKeyvals("v", []byte("a b"))
	if e != nil || string(b) != "v=\"a b\"" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}

func TestGoletaLogfmt013Boundary(t *testing.T) {
	b, e := MarshalKeyvals("v", []byte("plain"))
	if e != nil || string(b) != "v=plain" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}
