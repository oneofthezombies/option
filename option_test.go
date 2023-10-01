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

func TestAll(t *testing.T) {
	o := Some("foo")
	v, ok := o.Value()
	if !ok {
		t.Error("expected option to have value")
	}

	if v != "foo" {
		t.Error("expected option value to be foo")
	}

	o = None[string]()
	v, ok = o.Value()
	if ok {
		t.Error("expected option to not have value")
	}

	// If None is asigned then the value is the "zero value" of the type. In this case, the zero value of string is an empty string.
	// But you don't have to worry about value because we deal with cases where ok is false.
	// If value type is a struct, then the zero value is an empty struct.
	// And if value type is a pointer, then the zero value is nil.
	if v != "" {
		t.Error("expected option value to be empty string")
	}
}
