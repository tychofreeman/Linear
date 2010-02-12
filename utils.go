package linear

import (
	"fmt"
	"bignum"
	"reflect"
	"testing"
)


func valueToRational(v reflect.Value) (rational *bignum.Rational, success bool) {
	rational, success = nil, false
	switch i := v.(type) {
		case *reflect.StringValue:
			rational, _, _ = bignum.RatFromString(i.Get(), 10)
			success = true
		case *reflect.IntValue:
			rational, success = bignum.Rat(int64(i.Get()), 1), true
	}
	return
}

func forArgs(fn func(reflect.Value), vals ...) {
	

	vals2 := reflect.NewValue(vals)
	switch i := vals2.(type) {
		case *reflect.StructValue:
			for j := 0; j < i.NumField(); j++ {
				fn(i.FieldByIndex([]int{j}))
			}
	}
}

// --- Testing utils

func FailIf(t *testing.T, cond bool) {
	if cond {
		t.Fail()
	}
}

type Tst struct {
	t *testing.T
}

func Fail(t *testing.T) (test *Tst) {
	test = new(Tst)
	test.t = t
	return
}

func (t Tst) If(msg string, pred bool) {
	if pred {
		t.t.Error(msg)
	}
}

func intsAreEqual(expected, actual int) (msg string, pred bool) {
	msg = fmt.Sprintf("Should not have found %d, but did", actual)
	if actual == expected {
		pred = true
	}
	return
}

func intsAreNotEqual(expected, actual int) (msg string, pred bool) {
	msg = fmt.Sprintf("Expected %d; Actual %d", expected, actual)
	if actual != expected {
		pred = true
	}
	return
}

func rationalsAreNotEqual(expected, actual *bignum.Rational) (msg string, pred bool) {
	msg = fmt.Sprintf("Expected %o; Actual %o", expected, actual)
	if expected.Cmp(actual) != 0 {
		pred = true
	}
	return
}
