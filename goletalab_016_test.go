package logfmt

import (
	"strings"
	"testing"
)

func TestGoletaLogfmt016(t *testing.T) {
	d := NewDecoderSize(strings.NewReader("key=123456789"), 8)
	if d.ScanRecord() || d.Err() == nil {
		t.Fatalf("scan=%v err=%v", d.ScanRecord(), d.Err())
	}
}

func TestGoletaLogfmt016Boundary(t *testing.T) {
	d := NewDecoderSize(strings.NewReader("123456789"), 4)
	if d.ScanRecord() || d.Err() == nil {
		t.Fatalf("err=%v", d.Err())
	}
}
