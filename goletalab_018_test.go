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
