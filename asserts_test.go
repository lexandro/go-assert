package assert

import (
	"testing"
)

var testInt1 int = 11
var testInt2 int = 12

var EqualsTestParams = []struct {
	exp interface{}
	act interface{}
}{
	{true, true},
	{uint(1), uint(1)},
	{uint8(8), uint8(8)},
	{uint16(16), uint16(16)},
	{uint32(32), uint32(32)},
	{uint64(64), uint64(64)},
	{int(1), int(1)},
	{int8(8), int8(8)},
	{int16(16), int16(16)},
	{int32(32), int32(32)},
	{int64(64), int64(64)},
	{float32(32.5), float32(32.5)},
	{float64(64.5), float64(64.5)},
	{complex64(32 + 32i), complex64(32 + 32i)},
	{complex128(64 + 64i), complex128(64 + 64i)},
	{"test", "test"},
	{byte(15), byte(15)},
	{rune(16384), rune(16384)},
	{[4]int8{1, 2, 3, 4}, [4]int8{1, 2, 3, 4}}, // array
	{[]int8{5, 6, 7, 8}, []int8{5, 6, 7, 8}},   // slice
	{&testInt1, &testInt1},                     // ptr

}

var NotEqualsTestParams = []struct {
	exp interface{}
	act interface{}
}{
	{true, false},
	{uint(1), uint(2)},
	{uint8(8), uint8(9)},
	{uint16(16), uint16(17)},
	{uint32(32), uint32(33)},
	{uint64(64), uint64(65)},
	{int(1), int(2)},
	{int8(8), int8(9)},
	{int16(16), int16(17)},
	{int32(32), int32(33)},
	{int64(64), int64(65)},
	{float32(32.5), float32(33.5)},
	{float64(64.5), float64(65.5)},
	{complex64(32 + 32i), complex64(33 + 33i)},
	{complex128(64 + 64i), complex128(65 + 65i)},
	{"test", "anotherTest"},
	{byte(15), byte(14)},
	{rune(16384), rune(16383)},
	{[4]int8{1, 2, 3, 4}, [4]int8{1, 2, 3, 5}}, // array
	{[]int8{5, 6, 7, 8}, []int8{5, 6, 7, 9}},   // slice
	{&testInt1, &testInt2},                     // ptr

}

func Test_Equals(t *testing.T) {
	for _, etp := range EqualsTestParams {
		if !Equals(t, etp.exp, etp.act) {
			t.Fail()
		}
	}
}

func Test_NotEquals(t *testing.T) {
	//
	for _, etp := range NotEqualsTestParams {
		//for _, etp := range NotEqualsTestParams {
		if !NotEquals(t, etp.exp, etp.act) {
			t.Fail()
		}
	}
}

func Test_IsNil(t *testing.T) {
	if !IsNil(t, nil) {
		t.Fail()
	}
}

func Test_IsNotNil(t *testing.T) {
	if !IsNotNil(t, 1) {
		t.Fail()
	}
}

func Test_IsTrue(t *testing.T) {
	if !IsTrue(t, true) {
		t.Fail()
	}

}

func Test_IsFalse(t *testing.T) {
	if !IsFalse(t, false) {
		t.Fail()
	}
}

func Test_Comparator(t *testing.T) {
	for _, etp := range NotEqualsTestParams {
		if isEqual(etp.exp, etp.act) {
			t.Errorf("Values should not be equal. Expected same <%v> was not: <%v>", etp.exp, etp.act)
		}
	}
	//
	for _, etp := range EqualsTestParams {
		if !isEqual(etp.exp, etp.act) {
			t.Errorf("Values should be equal. Expected same <%v> was not: <%v>", etp.exp, etp.act)
		}
	}
}
