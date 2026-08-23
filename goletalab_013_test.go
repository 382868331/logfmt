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
