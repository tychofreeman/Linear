package linear

import (
	"reflect"
	//"fmt"
	//"log"
)


type Matrix struct {
	data MatrixData
	rows int
	cols int
}

func EmptyMatrix() Matrix {
	return MakeMatrix(0,0)
}

func MakeMatrix(rows int, cols int) Matrix {
	return Matrix{data: make([]MatrixRow, rows), rows: rows, cols: cols}
}

func (m Matrix) nullRowCount() int {
	return len(m.data.Reduce(func(mr MatrixRow) bool {
		return len(mr) == 0
	}))
}

func (m Matrix) IsComplete() bool {
	return 0 == m.nullRowCount()
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
		if i + 1 == len(m.data) {
			return false
		}
	}
	return true
}

func (m Matrix) SetCell(row, col int, i interface{}) bool {
	return false
}

func (m Matrix) IsEmpty() bool {
	return m.cols == 0 || m.rows == 0;
}
