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

func TestGoletaLogfmt017Boundary(t *testing.T) {
	d := NewDecoder(strings.NewReader(string([]byte{111, 107, 61, 49, 10, 98, 61, 34, 120})))
	if !d.ScanRecord() {
		t.Fatal("first")
	}
	for d.ScanKeyval() {
	}
	if !d.ScanRecord() {
		t.Fatal("second")
	}
	_ = d.ScanKeyval()
	e, ok := d.Err().(*SyntaxError)
	if !ok || e.Line != 2 {
		t.Fatalf("err=%v", d.Err())
	}
}
