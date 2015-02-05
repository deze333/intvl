package intvl

import (
	"fmt"
	"testing"
	"time"
)

//------------------------------------------------------------
// Tests for Time Interval
//------------------------------------------------------------

// Tests runtime intialization.
func TestRuntime(t *testing.T) {

	// Set debug output format
	Runtime_TimeLayout_Debug("Jan _2 15:04:05 MST")
}

// Tests base operations.
func TestBase(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	t1 := t0.Add(-4 * time.Hour)
	t2 := t0.Add(+100 * time.Hour)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	ti := &TimeInterval{
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// String
	fmt.Println(ti)

	// Length
	l := ti.Len()
	fmt.Println("Length =", l)
	fmt.Println("Length =", ti.LenHuman())
}

// Tests equality operations.
func TestEquality(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	t1 := t0.Add(-4 * time.Hour)
	t2 := t0.Add(+4 * time.Hour)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	ti := &TimeInterval{
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Equality: ts|te
	tiA := ti.Clone()

	if res := ti.IsEqual_TsTe(tiA); res != true {
		t.Error("Intervals must be equal")
	}

	tiA.Ts = ti.Ts.AddDate(0, 0, 1)
	if res := ti.IsEqual_TsTe(tiA); res != false {
		t.Error("Intervals must be NOT equal")
	}

	// Equality: ts|te|dt
	if res := ti.IsEqual_TsTeDt(tiA); res != false {
		t.Error("Intervals must be NOT equal")
	}

	tiA.Dt = ti.Dt * 2
	if res := ti.IsEqual_TsTeDt(tiA); res != false {
		t.Error("Intervals must be NOT equal")
	}
}

// Tests compare operations.
func TestCompare(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	t1 := t0.Add(-4 * time.Hour)
	t2 := t0.Add(+4 * time.Hour)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	ti := &TimeInterval{
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Exact match
	other := ti.Clone()

	if !ti.IsContains(other) || !other.IsContains(ti) {
		t.Error("Intervals must pass contain test")
	}

	// IsContains
	// ----------

	// Other is inside
	other.Ts = other.Ts.Add(+1 * time.Hour)
	other.Te = other.Te.Add(-1 * time.Hour)

	if !ti.IsContains(other) || !other.IsContainedBy(ti) {
		t.Error("Included intervals must pass contain test")
	}

	// Other extends to right
	other = ti.Clone()
	other.Ts = other.Ts.Add(+1 * time.Hour)
	other.Te = other.Te.Add(+8 * time.Hour)

	if ti.IsContains(other) || other.IsContainedBy(ti) {
		t.Error("Right extending intervals must pass contain test")
	}

	// Other extends to left
	other = ti.Clone()
	other.Ts = other.Ts.Add(-8 * time.Hour)
	other.Te = other.Te.Add(-1 * time.Hour)

	if ti.IsContains(other) || other.IsContainedBy(ti) {
		t.Error("Left extending intervals must pass contain test")
	}

	// IsBefore
	// ----------

	other = ti.Clone()
	other.Ts = t0.Add(+24 * time.Hour)
	other.Te = t0.Add(+20 * time.Hour)

	if !ti.IsBefore(other) || other.IsBefore(ti) {
		t.Error("IsBefore failed")
	}

	other = ti.Clone()
	other.Ts = t0.Add(+4 * time.Hour)
	other.Te = t0.Add(+8 * time.Hour)

	if !ti.IsBefore(other) {
		t.Error("IsBefore failed")
	}

	// IsStartsInside
	// --------------

	other = ti.Clone()
	other.Ts = t0.Add(-8 * time.Hour)
	other.Te = t0.Add(+8 * time.Hour)

	if !ti.IsStartsInside(other) {
		t.Error("IsStartsInside failed")
	}

	// IsLeftAdjacent
	// --------------

	other = ti.Clone()
	other.Ts = t0.Add(+4 * time.Hour)
	other.Te = t0.Add(+8 * time.Hour)

	if !ti.IsLeftAdjacent(other) {
		t.Error("IsLeftAdjacent failed")
	}

	// IsRightAdjacent
	// --------------

	other = ti.Clone()
	other.Ts = t0.Add(-8 * time.Hour)
	other.Te = t0.Add(-4 * time.Hour)

	if !ti.IsRightAdjacent(other) {
		t.Error("IsRightAdjacent failed")
	}

}

// Tests operations.
func TestOp(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	t1 := t0.Add(-4 * time.Hour)
	t2 := t0.Add(+4 * time.Hour)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	ti := &TimeInterval{
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	var res []*TimeInterval

	// Exclude
	// ------

	fmt.Println("Exclude = ", ti)

	// Equal
	res = ti.Exclude(ti)
	if len(res) != 0 {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// No overlap, lefter
	other := ti.Clone()
	other.Ts = t0.Add(-8 * time.Hour)
	other.Te = t0.Add(-7 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 1 || !res[0].IsEqual_TsTe(ti) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// No overlap, left adjacent
	other = ti.Clone()
	other.Ts = t0.Add(-8 * time.Hour)
	other.Te = t0.Add(-4 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 1 || !res[0].IsEqual_TsTe(ti) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// No overlap, righter
	other = ti.Clone()
	other.Ts = t0.Add(+8 * time.Hour)
	other.Te = t0.Add(+9 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 1 || !res[0].IsEqual_TsTe(ti) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// No overlap, right adjacent
	other = ti.Clone()
	other.Ts = t0.Add(+4 * time.Hour)
	other.Te = t0.Add(+8 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 1 || !res[0].IsEqual_TsTe(ti) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// Left overlap
	other = ti.Clone()
	other.Ts = t0.Add(-8 * time.Hour)
	other.Te = t0.Add(+2 * time.Hour)

	correct := ti.Clone()
	correct.Ts = t0.Add(+2 * time.Hour)
	correct.Te = t0.Add(+4 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 1 || !res[0].IsEqual_TsTe(correct) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// Right overlap
	other = ti.Clone()
	other.Ts = t0.Add(+2 * time.Hour)
	other.Te = t0.Add(+8 * time.Hour)

	correct = ti.Clone()
	correct.Ts = t0.Add(-4 * time.Hour)
	correct.Te = t0.Add(+2 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 1 || !res[0].IsEqual_TsTe(correct) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// Inside
	other = ti.Clone()
	other.Ts = t0.Add(-2 * time.Hour)
	other.Te = t0.Add(+2 * time.Hour)

	correctA := ti.Clone()
	correctA.Ts = t0.Add(-4 * time.Hour)
	correctA.Te = t0.Add(-2 * time.Hour)

	correctB := ti.Clone()
	correctB.Ts = t0.Add(+2 * time.Hour)
	correctB.Te = t0.Add(+4 * time.Hour)

	res = ti.Exclude(other)
	if len(res) != 2 || !res[0].IsEqual_TsTe(correctA) ||
		!res[1].IsEqual_TsTe(correctB) {
		t.Error("Exclude failed")
	}
	fmt.Println("Exclude =", res)

	// Split
	// ------

	src := ti.Clone()
	src.Ts = t0.Add(-2 * time.Hour)
	src.Te = t0.Add(+2 * time.Hour)

	fmt.Println("Split divisible = ", src)

	tis := src.Split(30 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	src = ti.Clone()
	src.Ts = t0.Add(-2 * time.Hour)
	src.Te = t0.Add(+2 * time.Hour)

	fmt.Println("Split indivisible = ", src)

	tis = src.Split(33 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split 1.x = ", src)

	tis = src.Split(80 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	// Split leftwards
	// ------

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split leftwards divisible = ", src)

	tis = src.SplitExtend_Leftwards(30 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split leftwards indivisible = ", src)

	tis = src.SplitExtend_Leftwards(33 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split leftwards 0.x = ", src)

	tis = src.SplitExtend_Leftwards(180 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	// Split rightwards
	// ------

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split rightwards divisible = ", src)

	tis = src.SplitExtend_Rightwards(30 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split rightwards indivisible = ", src)

	tis = src.SplitExtend_Rightwards(33 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))

	src = ti.Clone()
	src.Ts = t0.Add(-1 * time.Hour)
	src.Te = t0.Add(+1 * time.Hour)

	fmt.Println("Split rightwards 0.x = ", src)

	tis = src.SplitExtend_Rightwards(180 * time.Minute)
	fmt.Println(NewTimeIntervals(tis...))
}
