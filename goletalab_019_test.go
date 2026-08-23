package logfmt

import (
	"strings"
	"testing"
)

func TestGoletaLogfmt019(t *testing.T) {
	d := NewDecoder(strings.NewReader("a="))
	if !d.ScanRecord() || !d.ScanKeyval() || string(d.Key()) != "a" || d.Value() != nil {
		t.Fatalf("key=%q value=%q err=%v", d.Key(), d.Value(), d.Err())
	}
}

func TestGoletaLogfmt019Boundary(t *testing.T) {
	d := NewDecoder(strings.NewReader("a= b=2"))
	if !d.ScanRecord() || !d.ScanKeyval() || string(d.Key()) != "a" {
		t.Fatalf("key=%q err=%v", d.Key(), d.Err())
	}
}
