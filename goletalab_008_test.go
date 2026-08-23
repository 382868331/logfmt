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
