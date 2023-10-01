# Go Option Package

## Overview

The Go Option package provides an implementation of the `Option` type similar to Rust's `Option`. It allows you to represent optional values and handle them in a safe and expressive way, avoiding the need for explicit null checks.

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

### Checking if a Value is Present

You can use the Has method to check if an Option contains a value:

```go
if someOption.Has() {
    fmt.Println("SomeOption has a value:", someOption.Value())
} else {
    fmt.Println("SomeOption is None")
}

if noneOption.Has() {
    fmt.Println("NoneOption has a value:", noneOption.Value())
} else {
    fmt.Println("NoneOption is None")
}
```

### Retrieving the Value

To extract the value from an Option, you can use the Value method. However, make sure to check if a value is present using Has before calling Value to avoid panics:

```go
if someOption.Has() {
    value := someOption.Value()
    fmt.Println("SomeOption value:", value)
} else {
    fmt.Println("SomeOption is None")
}
```

### Panics

If you attempt to access the value of a None option without checking first, it will panic:

```go
// This will panic if noneOption is None
value := noneOption.Value()
```

To avoid panics, always use Has to check the presence of a value before calling Value.

## Examples

Here are some basic examples of using the option package:

```go
package main

import (
    "fmt"
    "github.com/oneofthezombies/option"
)

func main() {
    // Creating Some and None options
    someOption := option.Some(42)
    noneOption := option.None()

    // Checking if a value is present
    if someOption.Has() {
        fmt.Println("SomeOption has a value:", someOption.Value())
    } else {
        fmt.Println("SomeOption is None")
    }

    if noneOption.Has() {
        fmt.Println("NoneOption has a value:", noneOption.Value())
    } else {
        fmt.Println("NoneOption is None")
    }

    // Accessing values safely
    if noneOption.Has() {
        value := noneOption.Value()
        fmt.Println("NoneOption value:", value)
    } else {
        fmt.Println("NoneOption is None")
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
