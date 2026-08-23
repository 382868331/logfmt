package logfmt

import (
	"testing"
)

func TestGoletaLogfmt012(t *testing.T) {
	b, e := MarshalKeyvals("v", "null")
	if e != nil || string(b) != "v=\"null\"" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}

func TestGoletaLogfmt012Boundary(t *testing.T) {
	b, e := MarshalKeyvals("v", "ok")
	if e != nil || string(b) != "v=ok" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}
