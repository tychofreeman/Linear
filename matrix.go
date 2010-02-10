package linear

import (
	"bignum"
	"reflect"
	"fmt"
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
	for i := 0; i < len(md); i++ {
		if pred(md[i]) {
			tmpOut[i] = md[i]
		}
	}
	out = make(MatrixData, len(tmpOut))
	for i:= 0; i < len(tmpOut); i++ {
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

func (m Matrix) AddRow(vals ...) {
}

func forArgs(fn func(t reflect.Type), vals ...) {
	t := reflect.Typeof(vals)
	switch i := t.(type) {
		case *reflect.StructType:
			fmt.Printf("StructType: %d fields\n", i.NumField())
			for j := 0; j < i.NumField(); j++ {
				fmt.Printf("\tIs type int? %s %b\n", i.FieldByIndex([]int{j}).Type.(reflect.IntType))
				//need type switch to determine type here...
				fn(i.FieldByIndex([]int{j}).Type)
			}
	}
}

func (m Matrix) IsEmpty() bool {
	return m.cols == 0 || m.rows == 0;
}

func (m Matrix) AddRow(vals ... int) {
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
