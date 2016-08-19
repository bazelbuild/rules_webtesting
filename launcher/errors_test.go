package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestNewFromString(t *testing.T) {
	fromString := New("fromString", "error")
	if Component(fromString) != "fromString" {
		t.Errorf(`expected Component(fromString) == "fromString", got %s`, Component(fromString))
	}
	if !strings.Contains(fromString.Error(), "error") {
		t.Errorf(`expected fromString.Error() contains "error", got %s`, fromString.Error())
	}

	err := errors.New("error")
	fromError := New("fromError", err)
	if Component(fromError) != "fromError" {
		t.Errorf(`expected Component(fromError) == "fromError", got %s`, Component(fromError))
	}
	if !strings.Contains(fromError.Error(), "error") {
		t.Errorf(`got %v, expected fromError.Error() contains "error"`, fromError.Error())
	}

	if sameComp := New("fromString", fromString); sameComp != fromString {
		t.Errorf("expected sameComp == %+v, got %+v", fromString, sameComp)
	}

	if newComp := New("newComp", fromString); newComp != fromString {
		t.Errorf("expected newComp == %+v, got %+v", fromString, newComp)
	}
}
