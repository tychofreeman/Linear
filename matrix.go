package linear

import (
	"bignum"
	"reflect"
	//"fmt"
	//"log"
)

type MatrixData []MatrixRow
type MatrixRow []* bignum.Rational
type Matrix struct {
	data MatrixData
	rows int
	cols int
}

var emptyMatrix = MakeMatrix(0,0)

// This will be made more efficient by using vectors instead of arrays...
func (md MatrixData) Reduce(pred func(MatrixRow) bool) (out MatrixData) {
	tmpOut := make(MatrixData, len(md))
	count := 0
	for i := 0; i < len(md); i++ {
		if pred(md[i]) {
			tmpOut[i] = md[i]
			count += 1
		}
	}
	out = make(MatrixData, count)
	for i:= 0; i < len(out); i++ {
		out[i] = tmpOut[i]
	}
	return
}

func MakeMatrix(rows int, cols int) Matrix {
	return Matrix{data: make([]MatrixRow, rows), rows: rows, cols: cols}
}

func (m Matrix) IsZero() bool {
	//if rows == 0 || cols == 0 { return true }
	//return false
	return false
}

func (m Matrix) IsComplete() bool {
	return m.data != nil
}

func (m Matrix) AddRow(vals ...) bool {
	for i := 0; i < len(m.data); i++ {
		if m.data[i] == nil {
			m.data[i] = make(MatrixRow, m.cols)
			j := 0
			forArgs(
				func(v reflect.Value) {
					rational, _ := valueToRational(v)
					m.data[i][j] = rational
					j = j + 1
				}, vals)
			break
		}
	}
	return true
}

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

func (m Matrix) IsEmpty() bool {
	return m.cols == 0 || m.rows == 0;
}

func (m Matrix) nullRowCount() int {
	return len(m.data.Reduce(func(mr MatrixRow) bool {
		return mr == nil
	}))
}

func (m Matrix) Add(addend Matrix) (Matrix, bool) {
	return emptyMatrix, false
	/*
	if m.rows != addend.rows || m.cols != addend.cols {
		//log.Exitf("Could not add matrix1(%d,%d) to matrix2(%d,%d)", m.rows, m.cols, addend.rows, addend.cols);
		return emptyMatrix, false
	}
	return Matrix{}, true
	*/
}
