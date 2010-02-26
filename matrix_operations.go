package linear

import (
//	"bignum"
)


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

func (m Matrix) Multiply(m2 Matrix) (Matrix, bool) {
	if m.IsDegenerate() || m2.IsDegenerate() {
		return EmptyMatrix(), false
	}
	if !m.hasComplementaryDimension(m2) {
		return EmptyMatrix(), false
	}

	result := ZeroMatrix(m.rows, m.cols)
	return result, true
}

// Count leading zeros
func lz(mr MatrixRow) (lz int) {
	for _, i := range mr {
		if !i.IsZero() {
			break
		}
		lz += 1
	}
	return
}

// TODO: We can make this use go functions, which may speed things up...
func (m Matrix) IsEchelonForm() bool {
	prevZeros := -1
	for i := range m.data {
		zeros := lz(m.data[i])
		if zeros == len(m.data[i]) {
			continue
		}
		if zeros <= prevZeros {
			return false
		}
		prevZeros = zeros
		
	}
	return true
}

func (m Matrix) GetEchelonEquivalent() Matrix {
	nonZeros := 
	return EmptyMatrix()
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

func (m Matrix) switchRows(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}
