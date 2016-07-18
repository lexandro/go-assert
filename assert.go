package assert

import (
	"testing"
	"reflect"
	"strings"
)


/*
assert.Equals
assert.IsTrue
assert.IsFalse
assert.IsEmpty


*/

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

func NotEquals(t *testing.T, exp interface{}, act interface{}) (result bool) {
	result = !isEqual(exp, act)
	if !result {
		Error(t, "Values should be different. Expected: <%e>, but was <%a>", exp, act)

	}
	return
}

func Equals(t *testing.T, exp interface{}, act interface{}) (result bool) {
	result = isEqual(exp, act)

	if !result {
		Error(t, "Values should be equal. Expected same <%e> was not: <%a>", exp, act)
	}
	return
}

func IsNil(t *testing.T, v interface{}) bool {
	if v != nil {
		t.Errorf("Expected nil, but was: <%v>", v)
	}
	return true
}
func IsNotNil(t *testing.T, v interface{}) bool {
	if v == nil {
		t.Errorf("Expected not null, but was: <%v>", v)
	}
	return true
}

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

func Error(t *testing.T, msg string, exp interface{}, act interface{}) {
	expT := reflect.TypeOf(exp)
	actT := reflect.TypeOf(act)
	expF := getTypeFormat(expT)
	actF := getTypeFormat(actT)
	msg = strings.Replace(msg, "%e", expF, 1)
	msg = strings.Replace(msg, "%a", actF, 1)
	t.Errorf(msg, exp, act)
}

func getTypeFormat(t reflect.Type) (result string) {
	result, ok := kindFormatters[t.Kind()]
	if !ok {
		result = "%s"
	}
	return result
}
