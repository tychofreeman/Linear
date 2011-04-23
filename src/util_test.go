package linear

import (
	"testing"
	"reflect"
)

import . "big"


func TestFailIfEqualToWithUnequalInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreEqual(6, 5))
	if t2.Failed() {
		t.Error()
	}
}

func TestFailIfEqualToWithEqualInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreEqual(5, 5))
	if !t2.Failed() {
		t.Error()
	}
}

func TestFailIfNotEqualToWithUnequalInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreNotEqual(6, 5))
	if !t2.Failed() {
		t.Error()
	}
}

func TestFailIfNotEqualToWithEqualInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreNotEqual(5, 5))
	if t2.Failed() {
		t.Error()
	}
}

func TestValueToRationalWithEmptyString(t *testing.T) {
	rational, pred := valueToRational(reflect.NewValue(""))
	if !pred {
		t.Error("Converting an empty string to a *bignum.Rational should always be allowed.")
	}
	if rational == nil {
		t.Error("Converting an empty string to a *bignum.Rational should never result in a nil Rational pointer.")
	} else {
		Fail(t).If(rationalsAreNotEqual(NewRat(0, 1), rational))
	}
}

func TestValueToRationalWithOne(t *testing.T) {
	rational, pred := valueToRational(reflect.NewValue(1))
	if !pred {
		t.Error("Converting '1' to a *bignum.Rational should always return true.")
	}
	if rational == nil {
		t.Error("Converting '1' to a *bignum.Rational should never result in a nil Rational pointer.")
	} else {
		Fail(t).If(rationalsAreNotEqual(NewRat(1, 1), rational))
	}
}

func TestValueToRationalWithFractionString(t *testing.T) {
	oneFifth := "1/5"
	sv := reflect.NewValue(oneFifth)
	expected, _ := new(*Rat).SetString(oneFifth)
	rational, success := valueToRational(sv)
	if !success {
		t.Errorf("Converting the string %s to a *bignum.Rational should always return true.", oneFifth)
	}
	if rational != nil {
		Fail(t).If(rationalsAreNotEqual(rational, expected))
	} else {
		t.Errorf("Converting the string %s to a *bignum.Rational should never result in a nil Rational pointer.", oneFifth)
	}
}

func TestGetRowReturnsCorrectRow(t *testing.T) {
	m := unitMatrix(4)
	vector := m.getRow(2)
	v0 := vector[0]
	v1 := vector[1]
	v2 := vector[2]
	v3 := vector[3]

	actual := MakeMatrix(1, 4)
	actual.AddRow(v0, v1, v2, v3)

	expected := MakeMatrix(1, 4)
	expected.AddRow(0, 0, 1, 0)

	if !actual.Equals(expected) {
		expected.Print("Expected:")
		actual.Print("Actual:")
		t.Fail()
	}
}
