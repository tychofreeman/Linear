package linear
import (
	"testing"
)

func FailIf(t *testing.T, cond bool) {
	if cond {
		t.Fail()
	}
}

func TestMakeMatrixShouldReturnIncompleteMatrix(t *testing.T) {
	m := MakeMatrix(5,5)
	FailIf(t, m.IsComplete())
}

func TestMatrixWithZeroRowsOrColumnsShouldBeEmpty(t *testing.T) {
	m1 := Matrix{rows: 0}
	FailIf(t, !m1.IsEmpty())

	m2 := Matrix{cols: 0}
	FailIf(t, !m2.IsEmpty())
}
func TestMatrixWithNonZeroRowsAndColsShouldBeNotEmpty(t *testing.T) {
	m := Matrix{rows: 1, cols: 1}
	FailIf(t, m.IsEmpty())
}

func TestNonEmptyMatrixWithNilDataShouldntBeComplete(t *testing.T) {
	m := MakeMatrix(10, 4)
	FailIf(t, m.IsComplete())
}

func TestAddRowShouldIncrementRowCount(t *testing.T) {
	
}
	
func TestNonEmptyMatrixWithMissingRowShouldBeIncomplete(t *testing.T) {
}

func TestNonEmptyMatrixWithFullDataShouldBeIncomplete(t *testing.T) {
}

func TestMatrixWithTooFewRowsShouldBeIncopmlete(t *testing.T) {
}

func TestMatrixWithTooManyRowsShouldNotBeNormal(t *testing.T) {
}
func TestMatrixWithCorrectRowsAndColsShouldBeNormal(t *testing.T) {
}

// Now put in the operations

func TestMatrixAdditionFailsIfDifferentRowCount(t *testing.T) {
	m1 := Matrix {nil, 5, 10}
	m2 := Matrix {nil, 6, 10}
	m, b := m1.Add(m2)
	if !b {
		t.Error("Should not return true if row counts differ")
	}

	if m.IsEmpty() {
		t.Error("Should not return non-nil matrix if row counts differ")
	}
}

func TestMatrixAdditionFailsIfDifferentColumnCount(t *testing.T) {
	m1 := Matrix {nil, 10, 5}
	m2 := Matrix {nil, 10, 6}
	m1.Add(m2)
}

func TestMatrixAdditionFailsIfMissingData(t *testing.T) {
/*
	one := bignum.Rat(1,1)

	m1 := Matrix {nil, 0, 0}
	m2 := Matrix {
			MatrixData{
				MatrixRow{one, one, one},
				MatrixRow{one, one, one},
			},
			2, 3 }
	m1.Add(m2)
*/
}
