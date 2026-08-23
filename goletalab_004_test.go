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
