package logfmt

import (
	"strings"
	"testing"
)

func TestGoletaLogfmt011(t *testing.T) {
	b, e := MarshalKeyvals("msg", "a b")
	if e != nil || string(b) != "msg=\"a b\"" {
		t.Fatalf("b=%q e=%v", b, e)
	}
}

func TestGoletaLogfmt011Boundary(t *testing.T) {
	b, e := MarshalKeyvals("msg", "a	b")
	if e != nil || !strings.Contains(string(b), string([]byte{92, 116})) {
		t.Fatalf("b=%q e=%v", b, e)
	}
}
