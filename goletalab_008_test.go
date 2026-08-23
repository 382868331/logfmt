package logfmt

import (
	"testing"
)

func TestGoletaLogfmt008(t *testing.T) {
	b, e := MarshalKeyvals("a=b", 1)
	if e != nil || string(b) != "ab=1" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}

func TestGoletaLogfmt008Boundary(t *testing.T) {
	b, e := MarshalKeyvals("x=y=z", 2)
	if e != nil || string(b) != "xyz=2" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}
