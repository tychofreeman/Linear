package linear

import (
	"bignum"
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

func MakeMatrix(rows int, cols int) Matrix {
	return Matrix{data: nil, rows: rows, cols: cols}
}

func (m Matrix) IsZero() bool {
	//if rows == 0 || cols == 0 { return true }
	//return false
	return false
}

func (m Matrix) IsComplete() bool {
	return m.data != nil
}

func (m Matrix) IsEmpty() bool {
	return m.cols == 0 || m.rows == 0;
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
