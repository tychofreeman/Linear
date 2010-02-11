package linear
import (
	"testing"
	"bignum"
	"reflect"
	"fmt"
)

func FailIf(t *testing.T, cond bool) {
	if cond {
		t.Fail()
	}
}

type Tst struct {
	t *testing.T
}

func Fail(t *testing.T) (test *Tst) {
	test = new(Tst)
	test.t = t
	return
}

func (t Tst) If(msg string, pred bool) {
	if pred {
		t.t.Error(msg)
	}
}

func intsAreEqual(expected, actual int) (msg string, pred bool) {
	msg = fmt.Sprintf("Should not have found %d, but did", actual)
	if actual == expected {
		pred = true
	}
	return
}

func intsAreNotEqual(expected, actual int) (msg string, pred bool) {
	msg = fmt.Sprintf("Expected %d; Actual %d", expected, actual)
	if actual != expected {
		pred = true
	}
	return
}

func rationalsAreNotEqual(expected, actual *bignum.Rational) (msg string, pred bool) {
	msg = fmt.Sprintf("Expected %o; Actual %o", expected, actual)
	if expected.Cmp(actual) != 0 {
		pred = true
	}
	return
}

func TestFailIfEqualToWithUnequalInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreEqual(6, 5))
	if t2.Failed() {
		t.Error()
	}
}

func TestFailIfEqualToWithEqualInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreEqual(5, 5))
	if !t2.Failed() {
		t.Error()
	}
}

func TestFailIfNotEqualToWithUnequalInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreNotEqual(6, 5))
	if !t2.Failed() {
		t.Error()
	}
}

func TestFailIfNotEqualToWithEqualInts(t *testing.T) {
	t2 := new(testing.T)
	Fail(t2).If(intsAreNotEqual(5, 5))
	if t2.Failed() {
		t.Error()
	}
}

func TestValueToRationalWithEmptyString(t *testing.T) {
	rational, pred := valueToRational(reflect.NewValue(""))
	if !pred {
		t.Error("Converting an empty string to a *bignum.Rational should always be allowed.")
	}
	if rational == nil {
		t.Error("Converting an empty string to a *bignum.Rational should never result in a nil Rational pointer.")
	} else {
		Fail(t).If(rationalsAreNotEqual(bignum.Rat(0, 1), rational))
	}
}

func TestValueToRationalWithOne(t *testing.T) {
	rational, pred := valueToRational(reflect.NewValue(1))
	if !pred {
		t.Error("Converting '1' to a *bignum.Rational should always return true.")
	}
	if rational == nil {
		t.Error("Converting '1' to a *bignum.Rational should never result in a nil Rational pointer.")
	} else {
		Fail(t).If(rationalsAreNotEqual(bignum.Rat(1, 1), rational))
	}
}

func TestValueToRationalWithFractionString(t *testing.T) {
	oneFifth := "1/5"
	sv := reflect.NewValue(oneFifth)
	expected, _, _ := bignum.RatFromString(oneFifth, 10)
	rational, success := valueToRational(sv)
	if !success {
		t.Errorf("Converting the string %s to a *bignum.Rational should always return true.", oneFifth)
	}
	if rational != nil {
		Fail(t).If(rationalsAreNotEqual(rational, expected))
	} else {
		t.Errorf("Converting the string %s to a *bignum.Rational should never result in a nil Rational pointer.", oneFifth)
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
	oldRowCount := m.nullRowCount()
	if oldRowCount != 0 {
		t.Fatal("Could not proceed with test as the matrix is not correctly full.")
	}
	
	success := m.AddRow(5, 5, 3, 6)
	if success {
		t.Error("Should not return true when adding a row to a full matrix")
	}
	newRowCount := m.nullRowCount()
	Fail(t).If(intsAreNotEqual(0, newRowCount))
}

func TestNonEmptyMatrixWithMissingRowShouldBeIncomplete(t *testing.T) {
	t.Fail()
}

func TestNonEmptyMatrixWithFullDataShouldBeIncomplete(t *testing.T) {
	t.Fail()
}

func TestMatrixWithTooFewRowsShouldBeIncopmlete(t *testing.T) {
	t.Fail()
}

func TestMatrixWithTooManyRowsShouldNotBeNormal(t *testing.T) {
	t.Fail()
}

func TestMatrixWithCorrectRowsAndColsShouldBeNormal(t *testing.T) {
	t.Fail()
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
