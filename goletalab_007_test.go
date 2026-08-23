package logfmt

import (
	"errors"
	"testing"
)

func TestGoletaLogfmt007(t *testing.T) {
	_, e := MarshalKeyvals(nil, 1)
	if !errors.Is(e, ErrNilKey) {
		t.Fatalf("err=%v", e)
	}
}
