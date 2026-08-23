package logfmt

import (
	"strings"
	"testing"
)

func TestGoletaLogfmt017(t *testing.T) {
	d := NewDecoder(strings.NewReader(string([]byte{97, 61, 34, 120})))
	if !d.ScanRecord() {
		t.Fatal("record")
	}
	_ = d.ScanKeyval()
	e, ok := d.Err().(*SyntaxError)
	if !ok || e.Line != 1 {
		t.Fatalf("err=%v", d.Err())
	}
}
