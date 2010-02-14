package linear
import (
	"testing"
	"bignum"
)

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

func TestEmptyMatrixShouldBeComplete(t *testing.T) {
	FailIf(t, !EmptyMatrix().IsComplete())
}

func TestNonEmptyMatrixWithNilDataShouldntBeComplete(t *testing.T) {
	m := MakeMatrix(10, 4)
	FailIf(t, m.IsComplete())
}

func TestAddRowOnBlankMatrixShouldIncrementNonNullRowCount(t *testing.T) {
	m := MakeMatrix(10, 4)
	oldRowCount := m.nullRowCount()
	m.AddRow(5, 5, 3, 6)
	newRowCount := m.nullRowCount()
	Fail(t).If(intsAreNotEqual(oldRowCount - 1, newRowCount))
}

func TestAddRowOnMatrixWithTwoRowsShouldIncrementNonNullRowCount(t *testing.T) {
	m := MakeMatrix(10, 4)
	m.AddRow(5, 5, 3, 6)
	m.AddRow(5, 5, 3, 6)
	oldRowCount := m.nullRowCount()
	m.AddRow(5, 5, 3, 6)
	newRowCount := m.nullRowCount()
	Fail(t).If(intsAreNotEqual(oldRowCount - 1, newRowCount))
	Fail(t).If(intsAreNotEqual(3, 10 - newRowCount))
}

func TestAddRowOnCompleteMatrixShouldReturnFalse(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1, 1, 1, 1)
	m.AddRow(1, 1, 1, 1)
	m.AddRow(1, 1, 1, 1)
	m.AddRow(1, 1, 1, 1)
	
	success := m.AddRow(5, 5, 3, 6)
	if success {
		t.Error("Should not return true when adding a row to a full matrix")
	}
}

func TestAddRowOnCompleteMatrixShouldHaveZeroNullRows(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1, 1, 1, 1)
	m.AddRow(1, 1, 1, 1)
	m.AddRow(1, 1, 1, 1)
	m.AddRow(1, 1, 1, 1)
	newRowCount := m.nullRowCount()
	Fail(t).If(intsAreNotEqual(0, newRowCount))
}

func TestNonEmptyMatrixWithMissingRowShouldBeIncomplete(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1,1,1,1)
	m.AddRow(1,1,1,1)
	m.AddRow(1,1,1,1)
	FailIf(t, m.IsComplete())
}

func TestMatrixWithNoMissingRowsShouldBeComplete(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1,1,1,1)
	m.AddRow(1,1,1,1)
	m.AddRow(1,1,1,1)
	m.AddRow(1,1,1,1)
	FailIf(t, !m.IsComplete())
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

func TestSetCellOnValidAddrShouldPopulateColumns(t *testing.T) {
	m1 := MakeMatrix(4,4)
	m1.SetCell(2, 2, 5)
	if len(m1.data[2]) == 0 {
		t.Fatalf("Missing column data!!")
	}
}

func TestSetCellOnValidAddrShouldReturnTrue(t *testing.T) {
	m1 := MakeMatrix(4,4)
	if !m1.SetCell(2, 2, 5) {
		t.Fail()
	}
}

func TestSetCellOnInvalidAddrShouldReturnFalse(t *testing.T) {
	m := MakeMatrix(2,2)
	if m.SetCell(5,5,10) {
		t.Fail()
	}
}

func TestSetCellOnValidAddrWithIntShouldSetData(t *testing.T) {
	m1 := MakeMatrix(4,4)
	m1.SetCell(2, 2, 5)
	if m1.data[2][2].Cmp(bignum.Rat(5,1)) != 0 {
		t.Error("Did not set cell to correct value.")
	}
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
