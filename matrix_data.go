package linear

import (
	"bignum"
)

type MatrixRow []* bignum.Rational
type MatrixData []MatrixRow

func (md MatrixData) Reduce(pred func(MatrixRow) bool) (out MatrixData) {
	tmpOut := make(MatrixData, len(md))
	count := 0
	for i := 0; i < len(md); i++ {
		if pred(md[i]) {
			tmpOut[i] = md[i]
			count += 1
		}
	}
	out = make(MatrixData, count)
	for i:= 0; i < len(out); i++ {
		out[i] = tmpOut[i]
	}
	return
}
