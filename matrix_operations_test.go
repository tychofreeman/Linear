package linear
import (
	"testing"
	"bignum"
)

func unitMatrix(rows int) Matrix {
	cols := rows
	m := ZeroMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		m.SetCell(i, i, 1)
	}
	return m
}

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

func TestMatrixWithNoZerosShouldNotBeEchelonForm(t *testing.T) {
	m := nonZeroMatrix(5, 5)
	if m.IsEchelonForm() {
		t.Fail()
	}
}

func TestMatrixWithNonZeroEntriesDirectlyAboveOnAnotherShouldNotBeEchelonForm(t *testing.T) {
	m := MakeMatrix(4, 4)
	m.AddRow(1,2,3,4)
	m.AddRow(0,1,2,3)
	m.AddRow(0,1,2,3)
	m.AddRow(0,0,1,2)

	if m.IsEchelonForm() {
		t.Fail()
	}
}

func TestSwitchRows(t *testing.T) {
	m := nonZeroMatrix(5, 5)
	m.switchRows(0, 4)
	if m.data[4][0].Cmp(bignum.Rat(1,1)) != 0 {
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
		if !n.IsZero() {
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
	m1 := nonZeroMatrix(5,4)
	m2 := nonZeroMatrix(4,4)
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
		unitMatrix(4).Print("Unit(4) = ");
		m.Print("Unit(4).Multiplied(Unit(4)) = ");
		t.Fail()
	}
}

func TestGaussianEquivalentOfUnitMatrixEqualsUnitMatrix(t *testing.T) {
	m := unitMatrix(4).GetGaussianEquivalent()
	if !m.Equals(unitMatrix(4)) {
		unitMatrix(4).Print("unitMatrix(4) = ");
		m.Print("unitMatrix(4).Gaussian() = ");
		t.Fail()
	}
}
