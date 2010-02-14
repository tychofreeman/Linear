package linear

import (
	"bignum"
)

type MatrixRow []* bignum.Rational
type MatrixData []MatrixRow

func (md MatrixData) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for _, e := range md {
			ch <- e
		}
		close(ch)
	}()
	return ch
}
