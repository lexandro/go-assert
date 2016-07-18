# go-assert - lightweight assertion

[![GoDoc](https://godoc.org/github.com/lexandro/go-assert?status.svg)](https://godoc.org/github.com/lexandro/go-assert)

Simple assertion library for unit testing in Golang without vendoring/dependencies.

Design principles:

 * No additional dependencies
 * Readibility

# Installation

`go get github.com/lexandro/go-assert`

# Update

`go get -u github.com/lexandro/go-assert`

# Usage
```go
// required imports
import (
	"github.com/lexandro/go-assert"
	"testing"
)

// Simple equality test
func Test_Equals(t *testing.T) {
	assert.Equals(t, 1, 1)
}

// Simple non-equality test
func Test_NotEquals(t *testing.T) {
	assert.NotEquals(t, 2, 3)
}

// Simple test for nil
func Test_IsNil(t *testing.T) {
	assert.IsNil(t, nil)
}

// Simple test for not nil
func Test_IsNotNil(t *testing.T) {
	assert.IsNotNil(t, "")
}

// Simple test for true
func Test_IsTrue(t *testing.T) {
	assert.IsTrue(t, true)
}

// Simple test for false
func Test_IsFalse(t *testing.T) {
	assert.IsFalse(t, false)
}
```

# TODO:
assert.IsEmpty

Contributing
============

We would love for you to contribute to go-assert. Fork the repository and send a PR with the fix or report any issues!
