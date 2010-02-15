package linear

import (
	//"reflect"
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
	return EmptyMatrix(), true
}
