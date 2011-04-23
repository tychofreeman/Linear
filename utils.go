package linear

import (
	"fmt"
	"exp/bignum"
	"reflect"
	"testing"
	"exp/iterable"
)

// Create an nXn matrix with '1' on the diagonal, and zeros otherwise.
func unitMatrix(rows int) Matrix {
	cols := rows
	m := ZeroMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		m.SetCell(i, i, 1)
	}
	return m
}

func (m Matrix) getRow(index int) MatrixRow {
	row := make(MatrixRow, m.cols)
	copy(row, m.data[index])
	return row
}

func (m Matrix) getCol(index int) MatrixRow {
	column := make(MatrixRow, m.rows)
	for i, row := range m.data {
		column[i] = row[index]
	}
	return column
}

func valueToRational(v reflect.Value) (rational *bignum.Rational, success bool) {
	rational, success = nil, false
	switch i := v.(type) {
	case *reflect.StringValue:
		rational, _, _ = bignum.RatFromString(i.Get(), 10)
		success = true
	case *reflect.IntValue:
		rational, success = bignum.Rat(int64(i.Get()), 1), true
	case *reflect.InterfaceValue:
		rational, success = valueToRational(i.Elem())
	}
	return
}

func forArgs(fn func(reflect.Value), vals ...interface{}) {

	if len(vals) == 0 {
		return
	}
	vals2 := reflect.NewValue(vals)
	switch i := vals2.(type) {
	case *reflect.SliceValue:
		for j := 0; j < i.Len(); j++ {
			fn(i.Elem(j))
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

func (v MatrixRow) multiply(v2 MatrixRow) (result MatrixRow) {
	result = make(MatrixRow, len(v))
	for i := 0; i < len(v); i++ {
		result[i] = v[i].Mul(v2[i])
	}
	return
}

func (v MatrixRow) sumAll() *bignum.Rational {
	zero := bignum.Rat(0, 1)
	return iterable.Inject(v, zero, sum).(*bignum.Rational)
}
func sum(a interface{}, b interface{}) interface{} {
	return a.(*bignum.Rational).Add(b.(*bignum.Rational))
}
