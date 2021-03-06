/* 
	Simple matrix operations
*/
package linear

import (
	"sort"
)

import . "big"


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
			result.data[i][j] = new(Rat).Add(m.data[i][j], addend.data[i][j])
		}
	}
	return result, true
}

// Multiply given matrix by another matrix.
func (m Matrix) Multiply(m2 Matrix) (Matrix, bool) {
	if m.IsDegenerate() || m2.IsDegenerate() {
		return EmptyMatrix(), false
	}
	if !m.hasComplementaryDimension(m2) {
		return EmptyMatrix(), false
	}

	result := ZeroMatrix(m.rows, m2.cols)
	for i := 0; i < m.cols; i++ {
		for j := 0; j < m2.rows; j++ {
			// TODO: It would be nice not to use a string to communicate the Rational across a channel...
			ch := make(chan *Rat)
			go func() {
				col := m.getCol(j)
				row := m2.getRow(i)
				ch <- col.multiply(row).sumAll()
			}()
			result.data[i][j] = <- ch
			// Get col j from m and row i from m2
			// Multiply the two vectors, and add the values.
		}
	}
	return result, true
}

// Count leading zeros
func lz(mr MatrixRow) (lz int) {
	for _, v := range mr {
		if v != nil && v.Sign() != 0 {
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
	for i, row1 := range m.data {
		for j, row2 := range m.data[i+1 : m.rows] {
			m.data[j+i+1], _ = reduceRow(row1, row2)
		}
	}
	return m
}

func reduceRow(mr1, mr2 MatrixRow) (MatrixRow, bool) {
	lz1 := lz(mr1)
	lz2 := lz(mr2)
	if lz1 != lz2 || lz1 == len(mr1) {
		return mr2, (lz1 >= lz2)
	}

	ratio := new(Rat).Quo(mr1[lz1], mr2[lz1])
	mr3 := make(MatrixRow, len(mr1))

	for i := range mr1 {
		val := new(Rat).Mul(mr2[i], ratio)
		mr3[i] = new(Rat).Sub(mr1[i], val)
	}
	return mr3, true
}

// Swap two rows
func (m Matrix) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}
