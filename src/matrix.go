/*
	Basic matrix operations.
*/

package linear

import (
	"reflect"
	"fmt"
	"sort"
	"io"
	//"strings"
	//"log"
)

import . "big"


// Matrix is a two-dimensional collection of Rational numbers.
// It can be initialized with a row count and column count.
type Matrix struct {
	data MatrixData
	rows int
	cols int
	sort.Interface
}

// Len satisfies the sort.Interface interface
func (m Matrix) Len() int {
	return m.rows
}

// Less satisfies the sort.Interface interface
func (m Matrix) Less(i, j int) bool {
	return m.data.Less(i, j)
}

// EmptyMatrix creates a 0x0 matrix, which is useful only for comparison to other matricies.
func EmptyMatrix() Matrix {
	return MakeMatrix(0, 0)
}

// MakeMatrix initializes a matrix with a number of rows and columns
func MakeMatrix(rows int, cols int) Matrix {
	return Matrix{data: make(MatrixData, rows), rows: rows, cols: cols}
}

func (m Matrix) nullRowCount() (nullCount int) {
	for _, r := range m.data {
		if len(r) == 0 {
			nullCount = nullCount + 1
		}
	}
	return
}

// IsComplete only if there are no null rows.
func (m Matrix) IsComplete() bool {
	return 0 == m.nullRowCount()
}

// AddRow with the specified integer or string values.
func (m Matrix) AddRow(vals ...int64) bool {
	// TODO: Should use Find() to get first empty row, or ???
	for i := 0; i < len(m.data); i++ {
		if m.data[i] == nil {
			m.data[i] = createRow(m.cols, vals...)
			return true
		}
	}
	return false
}

func createRow(cols int, vals ...int64) MatrixRow {
	row := make(MatrixRow, cols)
	for i, v := range vals {
		row[i] = NewRat(v, 1)
		i += 1
	}
	return row
}

// SetCell to a value.
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

// IsEmpty if number of rows or columns is 0.
func (m Matrix) IsEmpty() bool {
	return m.cols == 0 || m.rows == 0
}

// ZeroMatrix creates an NxN matrix filled in with zero values.
func ZeroMatrix(rows, cols int) Matrix {
	m := MakeMatrix(rows, cols)
	for i := range m.data {
		m.data[i] = make(MatrixRow, cols)
		for j := range m.data[i] {
			m.data[i][j] = NewRat(0, 1)
		}
	}

	return m
}

func (m Matrix) print(w io.Writer, title string) {
	fmt.Fprintf(w, "Printing Matrix %v (%v rows, %v cols):\n", title, m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		fmt.Fprintf(w, "\t[")
		for j := 0; j < m.cols; j++ {
			cellStr := "<nil>"
			cell := m.data[i][j]
			if cell != nil {
				cellStr = fmt.Sprintf("%v, ", cell)
			}
			fmt.Fprintf(w, cellStr)
		}
		fmt.Fprintf(w, "]\n")
	}
}

func (m Matrix) hasSameDimension(m2 Matrix) bool {
	return m.cols == m2.cols && m.rows == m2.rows
}

func (m Matrix) hasComplementaryDimension(m2 Matrix) bool {
	return m.cols == m2.rows && m.rows == m2.cols
}

// IsDegenerate if not all rows or columns are filled in.
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

// Print out the matrix values as pretty as possible.
func (m Matrix) Print(name string) {
	fmt.Printf("%s\n", name)
	for _, r := range m.data {
		fmt.Printf("\t")
		for _, c := range r {
			fmt.Printf("%s,", c.String())
		}
		fmt.Printf("\n")
	}
}

// Equals determines if the given matrix has the same values as another matrix.
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
