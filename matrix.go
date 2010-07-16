package linear

import (
	"reflect"
	"exp/iterable"
	"bignum"
	"container/vector"
	"fmt"
	"sort"
	//"strings"
	//"log"
)


type Matrix struct {
	data MatrixData
	rows int
	cols int
	sort.Interface
}

// --- Satisfying the sort.Interface interface
func (m Matrix) Len() int {
	return m.rows
}

func (m Matrix) Less(i, j int) bool {
	return m.data.Less(i, j);
}

func EmptyMatrix() Matrix {
	return MakeMatrix(0,0)
}

func MakeMatrix(rows int, cols int) Matrix {
	return Matrix{data: make(MatrixData, rows), rows: rows, cols: cols}
}

func (m Matrix) nullRowCount() int {
	return len(
		iterable.Data(
			iterable.Filter(
				m.data,
				func(mr interface{}) bool {
					switch i := mr.(type) {
						case MatrixRow:
							return len(i) == 0
					}
					return false
				})))
}

func (m Matrix) IsComplete() bool {
	return 0 == m.nullRowCount()
}

func (m Matrix) AddRow(vals ...) bool {
	// TODO: Should use Find() to get first empty row, or ???
	for i := 0; i < len(m.data); i++ {
		if m.data[i] == nil {
			m.data[i] = createRow(m.cols, vals)
			break
		}
		if i + 1 == len(m.data) {
			return false
		}
	}
	return true
}

func createRow(cols int, vals ...) MatrixRow {
	row := make(MatrixRow, cols)
	i := 0
	forArgs(
		func(v reflect.Value) {
			rational, _ := valueToRational(v)
			row[i] = rational
			i += 1
		}, vals)
	return row
}

func (m Matrix) SetCell(row, col int, i interface{}) bool {
	if 0 > row || row >= m.rows || 0 > col || col >= m.cols {
		return false
	}
	if len(m.data[row]) == 0 {
		m.data[row] = make(MatrixRow, m.cols)
	}
	var success bool
	m.data[row][col], success = valueToRational(reflect.NewValue(i))
	return success
}

func (m Matrix) IsEmpty() bool {
	return m.cols == 0 || m.rows == 0;
}

func ZeroMatrix(rows, cols int) Matrix {
	m := MakeMatrix(rows, cols)
	for i := range m.data {
		m.data[i] = make(MatrixRow, cols)
		for j := range m.data[i] {
			m.data[i][j] = bignum.Rat(0, 1)
		}
	}
	return m
}

func (m Matrix) hasSameDimension(m2 Matrix) bool {
	return m.cols == m2.cols && m.rows == m2.rows
}

func (m Matrix) hasComplementaryDimension(m2 Matrix) bool {
	return m.cols == m2.rows && m.rows == m2.cols
}

func (m Matrix) IsDegenerate() bool {
	if len(m.data) != m.rows {
		return true
	}
	for _, c := range m.data {
		if len(c) != m.cols {
			return true
		}
	}
	return false 
}

func toStringArray(i iterable.Iterable) []string {
	v := new(vector.StringVector)
	for j := range i.Iter() {
		switch t := j.(type) {
			case *bignum.Rational:
				v.Push(t.String())
		}
	}
	return *v
}

func (m Matrix) Print(name string) {
	fmt.Printf("%s\n", name)
	for _, r := range m.data {
		fmt.Printf("\t")
		for _, c := range r {
			fmt.Printf("%s,",  c.String())
		}
		fmt.Printf("\n")
	}
}

func (m Matrix) Equals(m2 Matrix) bool {

	if !m.hasSameDimension(m2) {
		return false
	}
	if m.IsDegenerate() || m2.IsDegenerate() {
		return false
	}

	for i := range m.data {

		for j := range m.data[i] {
			if m.data[i][j] == nil {
				return false
			}
			
			if m.data[i][j].Cmp(m2.data[i][j]) != 0 {
				return false
			}
		}
	}
	return true
}

