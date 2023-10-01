package option_test

import (
	"testing"

	. "github.com/oneofthezombies/option"
)

func TestSome(t *testing.T) {
	o := Some(1)
	v, ok := o.Value()
	if !ok {
		t.Error("expected option to have value")
	}

	if v != 1 {
		t.Error("expected option value to be 1")
	}
}

func TestNone(t *testing.T) {
	o := None[int]()
	v, ok := o.Value()
	if ok {
		t.Error("expected option to not have value")
	}

	if v != 0 {
		t.Error("expected option value to be 0")
	}
}

func TestNoneString(t *testing.T) {
	o := None[string]()
	v, ok := o.Value()
	if ok {
		t.Error("expected option to not have value")
	}

	if v != "" {
		t.Error("expected option value to be empty string")
	}
}

func TestNoneEmptyStruct(t *testing.T) {
	o := None[struct{}]()
	v, ok := o.Value()
	if ok {
		t.Error("expected option to not have value")
	}

	if v != (struct{}{}) {
		t.Error("expected option value to be empty struct")
	}
}

func TestNoneStruct(t *testing.T) {
	type S struct {
		A int
		B string
	}

	o := None[S]()
	v, ok := o.Value()
	if ok {
		t.Error("expected option to not have value")
	}

	if v != (S{}) {
		t.Error("expected option value to be empty struct")
	}
}
