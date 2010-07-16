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

func (mr MatrixRow) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for _, e := range mr {
			ch <- e
		}
		close(ch)
	}()
	return ch
}

func (mr MatrixData) Less(l, r int) bool {
	// TODO: This needs to do a comparison against the matrix rows...
	for i := 0; i < len(mr[l]); i++ {
		cmp := mr[l][i].Cmp(mr[r][i])
		if cmp > 0 {
			return true;
		}
		if cmp < 0 {
			return false;
		}
	}
	return true;
}
