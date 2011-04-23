package linear

import (
	"testing"
)

import . "big"

func TestZeroMatrixShouldBeEchalonForm(t *testing.T) {
	zero := ZeroMatrix(4, 4)
	if !zero.IsEchelonForm() {
		t.Fail()
	}
}

func TestUnitMatrixShouldBeEchelonForm(t *testing.T) {
	unit := unitMatrix(5)
	if !unit.IsEchelonForm() {
		t.Fail()
	}
}

func TestUnitMatrixShouldBeReducedEchelonForm(t *testing.T) {
	unit := unitMatrix(5)
	if !unit.IsReducedEchelonForm() {
		t.Fail()
	}
}

func TestMatrixWithNoZerosShouldBeEchelonForm(t *testing.T) {
	m := nonZeroMatrix(5, 5)
	if !m.IsEchelonForm() {
		t.Fail()
	}
}

func TestMatrixWithNoZerosShouldNotBeReducedEchelonForm(t *testing.T) {
	m := nonZeroMatrix(5, 5)
	if m.IsReducedEchelonForm() {
		t.Fail()
	}
}

func TestMatrixWithNonZeroEntriesDirectlyAboveOnAnotherShouldBeEchelonForm(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1, 2, 3, 4)
	m.AddRow(0, 1, 2, 3)
	m.AddRow(0, 1, 2, 3)
	m.AddRow(0, 0, 1, 2)

	if !m.IsEchelonForm() {
		t.Fail()
	}
}

func TestMatrixWithNonZeroEntriesDirectlyAboveOnAnotherShouldNotBeReducedEchelonForm(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1, 2, 3, 4)
	m.AddRow(0, 1, 2, 3)
	m.AddRow(0, 1, 2, 3)
	m.AddRow(0, 0, 1, 2)

	if m.IsReducedEchelonForm() {
		t.Fail()
	}
}

func TestSwitchRows(t *testing.T) {
	m := nonZeroMatrix(5, 5)
	m.Swap(0, 4)
	if m.data[4][0].Cmp(NewRat(1, 1)) != 0 {
		t.Fail()
	}
}

func TestReduceZeroRow(t *testing.T) {
	m := ZeroMatrix(4, 4)
	m1 := m.data[0]
	m2 := m.data[1]
	_, success := reduceRow(m1, m2)
	if !success {
		t.Fail()
	}
}

func TestReduceEqualRow(t *testing.T) {
	m := nonZeroMatrix(4, 4)
	m1 := m.data[0]
	m2 := m.data[0]
	m3, success := reduceRow(m1, m2)
	if !success {
		t.Fail()
	}
	for i, n := range m3 {
		if n.Sign() != 0 {
			t.Errorf("Expected 0 at index %d; found %v", i, n)
		}
	}
}

func TestDegenerateMatrixCannotBeMultipliedByAnotherMatrix(t *testing.T) {
	deg := MakeMatrix(5, 5)
	m := nonZeroMatrix(5, 5)
	_, success := deg.Multiply(m)
	if success {
		t.Fail()
	}
}

func TestMatrixCannotBeMultipliedByDegenerateMatrix(t *testing.T) {
	deg := MakeMatrix(5, 5)
	m := nonZeroMatrix(5, 5)
	_, success := m.Multiply(deg)
	if success {
		t.Fail()
	}
}

func TestMatrixCannotBeMultipliedByMatrixWithWrongDimensions(t *testing.T) {
	m1 := nonZeroMatrix(5, 4)
	m2 := nonZeroMatrix(4, 4)
	_, success := m1.Multiply(m2)
	if success {
		t.Fail()
	}
}

func TestUnitMatrixMutlipliedByUnitMatrixShouldReturnTrue(t *testing.T) {
	_, success := unitMatrix(4).Multiply(unitMatrix(4))
	if !success {
		t.Fail()
	}
}

func TestUnitMatrixMutlipliedByUnitMatrixEqualsUnitMatrix(t *testing.T) {
	m, _ := unitMatrix(4).Multiply(unitMatrix(4))
	if !m.Equals(unitMatrix(4)) {
		unitMatrix(4).Print("Unit(4) = ")
		m.Print("Unit(4).Multiplied(Unit(4)) = ")
		t.Fail()
	}
}

func TestGaussianEquivalentOfUnitMatrixEqualsUnitMatrix(t *testing.T) {
	m := unitMatrix(4).AfterGaussianElimination()
	if !m.Equals(unitMatrix(4)) {
		unitMatrix(4).Print("unitMatrix(4) = ")
		m.Print("unitMatrix(4).Gaussian() = ")
		t.Fail()
	}
}

func TestGaussianEquivalentOfReorderedUnitMatrixEqualsUnitMatrix(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(0, 0, 0, 1)
	m.AddRow(0, 1, 0, 0)
	m.AddRow(1, 0, 0, 0)
	m.AddRow(0, 0, 1, 0)

	ge := m.AfterGaussianElimination()
	if !ge.Equals(unitMatrix(4)) {
		t.Fail()
	}
}

func TestGaussianEquivalentReordersRegardlessOfScaleOfRow(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(0, 0, 0, 5)
	m.AddRow(6, 0, 1, 10)
	m.AddRow(0, 10, 0, 10)
	m.AddRow(0, 0, 11, 27)

	mReordered := MakeMatrix(4, 4)
	mReordered.AddRow(6, 0, 1, 10)
	mReordered.AddRow(0, 10, 0, 10)
	mReordered.AddRow(0, 0, 11, 27)
	mReordered.AddRow(0, 0, 0, 5)
	ge := m.AfterGaussianElimination()
	if !ge.Equals(mReordered) {
		t.Fail()
	}
}

func TestNonReducedEchelonFormMatrixIsReducedEchelonFormAfterGaussianElimination(t *testing.T) {
	m := MakeMatrix(5, 5)
	m.AddRow(1, 2, 3, 4, 5)
	m.AddRow(2, 3, 4, 5, 6)
	m.AddRow(0, 1, 2, 3, 4)
	m.AddRow(0, 1, 1, 2, 3)
	m.AddRow(0, 0, 0, 0, 1)

	age := m.AfterGaussianElimination()

	if !age.IsReducedEchelonForm() {
		t.Fail()
	}
}

func TestDuplicateRowsAreEliminatedAfterGaussianElimination(t *testing.T) {
	m := MakeMatrix(2, 2)
	m.AddRow(1, 2)
	m.AddRow(1, 2)

	actual := m.AfterGaussianElimination()

	expected := MakeMatrix(2, 2)
	expected.AddRow(1, 2)
	expected.AddRow(0, 0)

	if !actual.Equals(expected) {
		t.Fail()
	}
}

func TestRowsWhichAreMultiplesAreEliminatedAfterGaussianElimination(t *testing.T) {
	m := MakeMatrix(2, 2)
	m.AddRow(1, 2)
	m.AddRow(2, 4)

	actual := m.AfterGaussianElimination()

	expected := MakeMatrix(2, 2)
	expected.AddRow(2, 4)
	expected.AddRow(0, 0)

	if !actual.Equals(expected) {
		t.Fail()
	}
}

func TestGaussianEliminationProducesCorrectResult(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(5, 6, 9, 3)
	m.AddRow(1, 1, 1, 0)
	m.AddRow(0, 1, 3, 1)
	m.AddRow(0, 0, 1, 1)

	expected := MakeMatrix(4, 4)
	expected.AddRow(5, 6, 9, 3)
	expected.AddRow(0, 1, 4, 3)
	expected.AddRow(0, 0, 1, 2)
	expected.AddRow(0, 0, 0, 1)

	actual := m.AfterGaussianElimination()

	if !actual.Equals(expected) {
		t.Fail()
	}
}

func TestUnitMatrixMultipliedByAEqualsA(t *testing.T) {
	a := MakeMatrix(4, 4)
	a.AddRow(5, 9, 3, 10)
	a.AddRow(7, 8, 15, 9)
	a.AddRow(1, 3, 5, 7)
	a.AddRow(11, 13, 17, 23)

	unit := unitMatrix(4)

	result, _ := unit.Multiply(a)

	if !result.Equals(a) {
		t.Fail()
	}
}
