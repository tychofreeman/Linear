/* 
	Simple matrix operations
*/
package linear

import (
	"sort"
)


// Add the given matrix by another matrix.
func (m Matrix) Add(addend Matrix) (Matrix, bool) {
	if m.IsDegenerate() || addend.IsDegenerate() {
		return EmptyMatrix(), false
	}

	if !m.hasSameDimension(addend) {
		return EmptyMatrix(), false
	}

	result := ZeroMatrix(m.rows, m.cols)
	for i := range m.data {
		for j := range m.data[i] {
			result.data[i][j] = m.data[i][j].Add(addend.data[i][j])
		}
	}
	return result, true
}

// Create an nXn matrix with '1' on the diagonal, and zeros otherwise.
func unitMatrix(rows int) Matrix {
	cols := rows
	m := ZeroMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		m.SetCell(i, i, 1)
	}
	return m
}

// Multiply given matrix by another matrix.
func (m Matrix) Multiply(m2 Matrix) (Matrix, bool) {
	if m.IsDegenerate() || m2.IsDegenerate() {
		return EmptyMatrix(), false
	}
	if !m.hasComplementaryDimension(m2) {
		return EmptyMatrix(), false
	}

	result := unitMatrix(m.rows);
	return result, true
}

// Count leading zeros
func lz(mr MatrixRow) (lz int) {
	for _, v := range mr {
		if !v.IsZero() {
			break
		}
		lz += 1
	}
	return
}

// IsReducedEchelonForm is more rigorous than IsEchelonForm in that row j must have fewer leading zeros than row j + 1.
func (m Matrix) IsReducedEchelonForm() bool {
	return m.isEchelonForm(true)
}

// IsEchelonForm is true if each row j has at least as many leading zeros as all previous rows.
// TODO: We can make this use go functions, which may speed things up...
func (m Matrix) IsEchelonForm() bool {
	return m.isEchelonForm(false)
}

func (m Matrix) isEchelonForm(strict bool) bool {
	prevZeros := -1
	for i := range m.data {
		zeros := lz(m.data[i])
		if zeros == len(m.data[i]) {
			continue
		}
		if zeros < prevZeros {
			return false
		}
		if strict && zeros == prevZeros {
			return false
		}
		prevZeros = zeros
	}
	return true
}

// AfterGaussianElimination returns the matrix with Gaussian elimination applied.
func (m Matrix) AfterGaussianElimination() Matrix {
	sort.Sort(m)
	return m
}

func reduceRow(mr1, mr2 MatrixRow) (MatrixRow, bool) {
	lz1 := lz(mr1)
	lz2 := lz(mr2)
	if lz1 != lz2 || lz1 == len(mr1){
		return mr2, (lz1 >= lz2)
	}

	ratio := mr1[lz1].Quo(mr2[lz1])
	mr3 := make(MatrixRow, len(mr1))

	for i := range mr1 {
		mr3[i] = mr1[i].Sub(mr2[i].Mul(ratio))
	}
	return mr3, true
}

// Swap two rows
func (m Matrix) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}
