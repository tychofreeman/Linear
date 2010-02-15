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

// Now put in the operations

func TestMatrixAdditionFailsIfDifferentRowCount(t *testing.T) {
	m1 := MakeMatrix(4, 4)
	m2 := MakeMatrix(5, 4)
	_, b := m1.Add(m2)
	if b {
		t.Fail()
	}
}

func TestMatrixAdditionFailsIfDifferentColumnCount(t *testing.T) {
	m1 := MakeMatrix(4, 4)
	m2 := MakeMatrix(4, 5)
	_, b := m1.Add(m2)
	if b {
		t.Fail()
	}
	
}

func TestMatrixAdditionSucceedsIfSameRowCountAndSameColCount(t *testing.T) {
	m1 := MakeMatrix(4, 5)
	m2 := MakeMatrix(4, 5)
	_, b := m1.Add(m2)
	if !b {
		t.Fail()
	}
}

func TestZeroMatrixIsEqualToZeroMatrix(t *testing.T) {
	m1 := ZeroMatrix(4, 4)
	m2 := ZeroMatrix(4, 4)
	if !m1.Equals(m2) {
		t.Fail()
	}
}

func nonZeroMatrix4x4() Matrix {
	m := ZeroMatrix(4, 4)
	m.AddRow(1, 2, 3, 4)
	m.AddRow(2, 2, 3, 4)
	m.AddRow(3, 3, 3, 4)
	m.AddRow(4, 4, 4, 4)
	return m
}

func TestNonZeroMatrixIsEqualToItself(t *testing.T) {
	m := nonZeroMatrix4x4()
	if !m.Equals(m) {
		t.Fail()
	}
}

func TestNonZeroMatrixIsEqualToMatrixWithSameValues(t *testing.T) {
	m1 := nonZeroMatrix4x4()
	m2 := nonZeroMatrix4x4()
	if !m1.Equals(m2) {
		t.Fail()
	}
}

func TestNonZeroMatrixIsNotEqualToZeroMatrix(t *testing.T) {
	if nonZeroMatrix4x4().Equals(ZeroMatrix(4,4)) {
		t.Fail()
	}
}

func TestNonEmptyMatrixIsNotEqualToEmptyMatrix(t *testing.T) {
}

func TestMatrixIsEqualToACopyOfItself(t *testing.T) {
}

func TestMatrixAddedToAdditiveUnityMatrixIsEqualToItself(t *testing.T) {
}

func TestMatrixOfOnesAddedToNegativeOnesIsZeroMatrix(t *testing.T) {
}

func TestZeroMatrixHasZerosInAllCells(t *testing.T) {
	zero := ZeroMatrix(5,5)
	for i, r := range zero.data {
		for j, c := range r {
			if !c.IsZero() {
				t.Error("Wrong value at row,col ", i, ",", j, ": found ", c, " instead of 0/1.")
			}
		}
	}
}

func TestZeroMatrixEqualsZeroMatrixIfDimensionsAreNotEqual(t *testing.T) {
	zero1 := ZeroMatrix(4, 4)
	zero2 := ZeroMatrix(4, 5)
	if zero1.Equals(zero2) {
		t.Fail()
	}
}

func TestNewMatrixIsNotEqualToZeroMatrix(t *testing.T) {
	empty := MakeMatrix(4, 4)
	zero := ZeroMatrix(4, 4)
	if zero.Equals(empty) {
		t.Error("zero.Equals(empty) should be false")
	}
	if empty.Equals(zero) {
		t.Error("empty.Equals(zero) should be false")
	}
}

func TestZeroMatrixEqualsZeroMatrixIfDimensionsAreEqual(t *testing.T) {
	zero1 := ZeroMatrix(4, 4)
	zero2 := ZeroMatrix(4, 4)
	if !zero1.Equals(zero2) {
		t.Fail()
	}
}

func TestNewMatrixIsDegenerate(t *testing.T) {
	m := MakeMatrix(5, 5)
	if !m.IsDegenerate() {
		t.Fail()
	}
}

func TestEmptyMatrixIsNotDegenerate(t *testing.T) {
	m := EmptyMatrix()
	if m.IsDegenerate() {
		t.Fail()
	}
}

func TestZeroMatrixIsNotDegenerate(t *testing.T) {
	zero := ZeroMatrix(4, 4)
	if zero.IsDegenerate() {
		t.Fail()
	}
}

func TestSingleCallOnSetCellCreatesADegenerateMatrix(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.SetCell(2, 2, 1)
	if !m.IsDegenerate() {
		t.Fail()
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
