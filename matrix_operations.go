package linear

import (
	//"reflect"
)

func (m Matrix) Add(addend Matrix) (Matrix, bool) {
	return EmptyMatrix(), false
	/*
	if m.rows != addend.rows || m.cols != addend.cols {
		//log.Exitf("Could not add matrix1(%d,%d) to matrix2(%d,%d)", m.rows, m.cols, addend.rows, addend.cols);
		return emptyMatrix, false
	}
	return Matrix{}, true
	*/
}
