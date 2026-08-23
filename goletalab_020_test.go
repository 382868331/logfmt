package logfmt

import (
	"strings"
	"testing"
)

func TestGoletaLogfmt020(t *testing.T) {
	d := NewDecoder(strings.NewReader(string([]byte{97, 61, 34, 120})))
	if !d.ScanRecord() {
		t.Fatal("record")
	}
	_ = d.ScanKeyval()
	e, ok := d.Err().(*SyntaxError)
	if !ok || e.Pos != 5 {
		t.Fatalf("err=%v", d.Err())
	}
}

func TestGoletaLogfmt020Boundary(t *testing.T) {
	d := NewDecoder(strings.NewReader(string([]byte{97, 98, 99, 61, 34, 117, 110, 116, 101, 114, 109, 105, 110, 97, 116, 101, 100})))
	if !d.ScanRecord() {
		t.Fatal("record")
	}
	_ = d.ScanKeyval()
	e, ok := d.Err().(*SyntaxError)
	if !ok || e.Pos != 18 {
		t.Fatalf("err=%v", d.Err())
	}
}
