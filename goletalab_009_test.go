package logfmt

import (
	"errors"
	"testing"
)

func TestGoletaLogfmt009(t *testing.T) {
	_, e := MarshalKeyvals(" = ", 1)
	if !errors.Is(e, ErrInvalidKey) {
		t.Fatalf("err=%v", e)
	}
}

func TestGoletaLogfmt009Boundary(t *testing.T) {
	_, e := MarshalKeyvals("	", 1)
	if !errors.Is(e, ErrInvalidKey) {
		t.Fatalf("err=%v", e)
	}
}
