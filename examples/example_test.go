// Package examples is a quite simple list of usage scenarios of the go-assert
package examples

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
