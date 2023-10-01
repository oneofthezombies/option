# Option for Go

## Overview

The Go Option package provides an implementation of the `Option` type similar to Rust's `Option`.  
It allows you to represent optional values and handle them in a safe and expressive way, avoiding the need for explicit null checks.  
And This is written as `go-way`.

## Usage

### Importing the Package

Import the `option` package in your Go code to use it:

```go
import "github.com/oneofthezombies/option"
```

### Creating Options

You can create Option instances using the Some and None functions. Some is used to wrap a value, while None represents the absence of a value:

```go
value := 42
someOption := option.Some(value)
noneOption := option.None()
```

### Checking if a Value is Present and Extracting the Value

You can use the Value method to check if a value is present in an Option. It returns a boolean indicating whether a value is present or not:

```go
// option_test.go
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
```

## Examples

Here are some basic examples of using the option package:

```go
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
```

## Contributing

If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure they pass the existing tests.
4. Add new tests if necessary.
5. Commit your changes with clear commit messages.
6. Push your branch to your fork.
7. Create a pull request to the original repository.

## License

This package is licensed under the MIT License. See the LICENSE file for details.
