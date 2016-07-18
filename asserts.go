// Package assert is lightweight framework of common assertions to make unit testing easier and readable
package assert

import (
	"testing"
	"reflect"
	"strings"
)

// kindFormatters is a map of assignments of data types and their formatters.
var kindFormatters = map[reflect.Kind]string{
	reflect.Uint:"%d",
	reflect.Uint8:"%d",
	reflect.Uint16:"%d",
	reflect.Uint32:"%d",
	reflect.Uint64:"%d",
	reflect.Int:"%d",
	reflect.Int8:"%d",
	reflect.Int16:"%d",
	reflect.Int32:"%d",
	reflect.Int64:"%d",
	reflect.Float32:"%f",
	reflect.Float64:"%f",
	reflect.Complex64:"%g",
	reflect.Complex128:"%g",
	reflect.Bool:"%t",
	reflect.String:"%s",
	reflect.Slice:"%#v",
	reflect.Array:"%#v",
	reflect.Ptr:"%#v",
}

// Equals asserts the equality of the given arguments
//
//   assert.Equals(t, 111, 111)
//
// Returns true if the arguments are equal and false if different.
func Equals(t *testing.T, exp interface{}, act interface{}) (result bool) {
	result = isEqual(exp, act)
	if !result {
		error(t, "Values should be equal. Expected same <%e> was not: <%a>", exp, act)
	}
	return
}

// NotEquals asserts the non-equality of the given arguments
//
//   assert.NotEquals(t, 111, 222)
//
// Returns true if the arguments are not equal and false if equal.
func NotEquals(t *testing.T, exp interface{}, act interface{}) (result bool) {
	result = !isEqual(exp, act)
	if !result {
		error(t, "Values should be different. Expected: <%e>, but was <%a>", exp, act)
	}
	return
}

// IsNil asserts the given argument is nil
//
//   assert.IsNil(t, nil)
//
// Returns true if the argument is nil and false if initialized to something.
func IsNil(t *testing.T, v interface{}) bool {
	if v != nil {
		t.Errorf("Expected nil, but was: <%v>", v)
	}
	return true
}

// IsNotNil asserts the given argument is initialized
//
//   assert.IsNotNil(t, 1)
//
// Returns true if the argument is initialized and false if nil.
func IsNotNil(t *testing.T, v interface{}) bool {
	if v == nil {
		t.Errorf("Expected not null, but was: <%v>", v)
	}
	return true
}

// IsTrue asserts the given bool argument is true
//
//   assert.IsTrue(t, true)
//
// Returns the argument's value
func IsTrue(t *testing.T, v bool) bool {
	if !v {
		t.Errorf("Expected true, but was: <%v>", v)
	}
	return v
}

// IsFalse asserts the given bool argument is false
//
//   assert.IsFalse(t, false)
//
// Returns the negated value of the argument
func IsFalse(t *testing.T, v bool) bool {
	if v {
		t.Errorf("Expected false, but was: <%v>", v)
	}
	return ! v
}

// Error uses the formatter map to display the given error message with readable formatting.
// The %e is placeholder for the first argument and %a is the second.
//
//   assert.error(t, "Error comparing %e with %a", "test", true)
//
func error(t *testing.T, msg string, exp interface{}, act interface{}) {
	expT := reflect.TypeOf(exp)
	actT := reflect.TypeOf(act)
	expF := getTypeFormat(expT)
	actF := getTypeFormat(actT)
	msg = strings.Replace(msg, "%e", expF, 1)
	msg = strings.Replace(msg, "%a", actF, 1)
	t.Errorf(msg, exp, act)
}

// isEqual is an internal function to check the equality of the given arguments. Only uses DeepEqual for arrays and slices.
//
//   isEqual(t, 111, 111)
//
// Returns true if the arguments are equal and false if different.
func isEqual(exp interface{}, act interface{}) (result bool) {
	expT := reflect.TypeOf(exp)
	actT := reflect.TypeOf(act)
	if expT.Kind() != reflect.Slice && actT.Kind() != reflect.Slice {
		result = exp == act
	} else {
		result = reflect.DeepEqual(act, exp)
	}
	return
}
// getTypeFormat is an internal function to provide the formatter string for the given data type.
//
//   getTypeFormat(reflect.Uint)
//
// Returns the formatter string assigned through the formatter map
func getTypeFormat(t reflect.Type) (result string) {
	result, ok := kindFormatters[t.Kind()]
	if !ok {
		result = "%s"
	}
	return result
}
