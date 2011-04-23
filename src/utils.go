package linear

import (
	"fmt"
	"reflect"
	"testing"
	"os"
)

import . "big"

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

func valueToRational(v reflect.Value) (rational *Rat, success bool) {
	rational, success = nil, false
	switch i := v; i.Kind() {
	case reflect.String:
		str := i.String()
		if len(str) == 0 {
			str = "0"
		}
		rational, _ = new(Rat).SetString(str)
		success = true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rational, success = NewRat(int64(i.Int()), 1), true
	case reflect.Interface:
		rational, success = valueToRational(i.Elem())
	}
	return
}

func forArgs(fn func(reflect.Value), vals ...interface{}) {
	if len(vals) == 0 {
		return
	}
	vals2 := reflect.NewValue(vals)
	switch i := vals2; i.Kind() {
	case reflect.Slice:
		for j := 0; j < i.Len(); j++ {
			fn(i.Index(j))
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

func rationalsAreNotEqual(expected, actual *Rat) (msg string, pred bool) {
	msg = fmt.Sprintf("Expected %o; Actual %o", expected, actual)
	if expected.Cmp(actual) != 0 {
		pred = true
	}
	return
}

func (v MatrixRow) multiply(v2 MatrixRow) (result MatrixRow) {
	if v2 == nil {
		fmt.Fprintf(os.Stderr, "MatrixRow cannot be null here\n")
	}
	if len(v2) < len(v) {
		fmt.Fprintf(os.Stderr, "Length of v2 (%v) < length of v (%v)\n", len(v2), len(v))
	}
	result = make(MatrixRow, len(v))
	for i := 0; i < len(v); i++ {
		if v[i] == nil {
			v[i] = NewRat(0, 1)
		}
		if v2[i] == nil {
			v2[i] = NewRat(0, 1)
		}
		result[i] = new(Rat).Mul(v[i], v2[i])
	}
	return
}

func (v MatrixRow) sumAll() *Rat{
	sum := NewRat(0, 1)
	for _, r := range v {
		sum = sum.Add(sum, r)
	}
	return sum
}

func sum(a interface{}, b interface{}) interface{} {
	return new(Rat).Add(a.(*Rat), b.(*Rat))
}
