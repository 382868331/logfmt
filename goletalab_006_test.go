package logfmt

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestGoletaLogfmt006(t *testing.T) {
	e := &MarshalerError{Type: reflect.TypeOf(1), Err: errors.New("boom")}
	if got := e.Error(); !strings.Contains(got, "boom") {
		t.Fatalf("got=%q", got)
	}
}
