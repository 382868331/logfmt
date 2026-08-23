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
