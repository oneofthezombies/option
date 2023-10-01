package option_test

import (
	"testing"

	. "github.com/oneofthezombies/option"
)

func TestSome(t *testing.T) {
	o := Some(1)

	if !o.Has() {
		t.Error("expected option to have value")
	}

	if o.Value() != 1 {
		t.Error("expected option value to be 1")
	}
}

func TestNone(t *testing.T) {
	o := None[int]()

	if o.Has() {
		t.Error("expected option to not have value")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()

	o.Value()
}
