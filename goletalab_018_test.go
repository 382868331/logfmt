package logfmt

import (
	"strings"
	"testing"
)

func TestGoletaLogfmt018(t *testing.T) {
	d := NewDecoder(strings.NewReader("   a=1"))
	if !d.ScanRecord() || !d.ScanKeyval() || string(d.Key()) != "a" {
		t.Fatalf("key=%q err=%v", d.Key(), d.Err())
	}
}

func TestGoletaLogfmt018Boundary(t *testing.T) {
	d := NewDecoder(strings.NewReader("		b=2"))
	if !d.ScanRecord() || !d.ScanKeyval() || string(d.Key()) != "b" {
		t.Fatalf("key=%q err=%v", d.Key(), d.Err())
	}
}
