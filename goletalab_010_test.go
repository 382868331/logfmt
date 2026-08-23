package logfmt

import (
	"testing"
)

func TestGoletaLogfmt010(t *testing.T) {
	b, e := MarshalKeyvals("a", nil)
	if e != nil || string(b) != "a=null" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}

func TestGoletaLogfmt010Boundary(t *testing.T) {
	var p *int
	b, e := MarshalKeyvals("p", p)
	if e != nil || string(b) != "p=null" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}
