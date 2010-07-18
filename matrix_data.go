/* 
	Simple matrix operations.
*/
package linear

import "exp/bignum"

// MatrixRow is an array of pointers to Rational objects.
type MatrixRow []* bignum.Rational

// MatrixData is an array of MatrixRow objects.
type MatrixData []MatrixRow

// Iter iterates through rows
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

// Iter iterates through values
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

// Less is needed for sorting.
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
