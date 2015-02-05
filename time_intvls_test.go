package intvl

import (
	"fmt"
	"testing"
	"time"
)

//------------------------------------------------------------
// Tests for Time Intervals
//------------------------------------------------------------

// Tests intervals.
func TestIntervals(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	// Interval 1
	t1 := t0.Add(-4 * time.Hour)
	t2 := t0.Add(+4 * time.Hour)

	ti1 := &TimeInterval{
		Name:   "One",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 2, overlaps right
	t1 = t0.Add(+0 * time.Hour)
	t2 = t0.Add(+8 * time.Hour)

	ti2 := &TimeInterval{
		Name:   "Two",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 3, far right
	t1 = t0.Add(+9 * time.Hour)
	t2 = t0.Add(+12 * time.Hour)

	ti3 := &TimeInterval{
		Name:   "Three",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 4, overlaps left
	t1 = t0.Add(-10 * time.Hour)
	t2 = t0.Add(-2 * time.Hour)

	ti4 := &TimeInterval{
		Name:   "Four",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 5, overlaps left, same points
	t1 = t0.Add(-10 * time.Hour)
	t2 = t0.Add(-2 * time.Hour)

	ti5 := &TimeInterval{
		Name:   "Five",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 6, far left
	t1 = t0.Add(-16 * time.Hour)
	t2 = t0.Add(-12 * time.Hour)

	ti6 := &TimeInterval{
		Name:   "Six",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 7, far x2 left
	t1 = t0.Add(-48 * time.Hour)
	t2 = t0.Add(-36 * time.Hour)

	ti7 := &TimeInterval{
		Name:   "Seven",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Interval 8, cover them all
	t1 = t0.Add(-72 * time.Hour)
	t2 = t0.Add(+72 * time.Hour)

	ti8 := &TimeInterval{
		Name:   "Eight",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// String
	fmt.Println(ti1)
	fmt.Println(ti2)
	fmt.Println(ti3)
	fmt.Println(ti4)
	fmt.Println(ti5)
	fmt.Println(ti6)
	fmt.Println(ti7)
	fmt.Println(ti8)

	// String: steps
	tis := NewTimeIntervals(ti1, ti2, ti3, ti4, ti5, ti6, ti7)
	fmt.Println()
	fmt.Println(tis.String())

	// String: steps and cover-all
	tis = NewTimeIntervals(ti1, ti2, ti3, ti4, ti5, ti6, ti7, ti8)
	fmt.Println()
	fmt.Println(tis.String())

	// Parses from strings
	fromStrings := func(times ...string) {

		layout := "2006 Jan 2 15:04:05 MST"
		tis, err := Parse_TimeIntervals(layout, times...)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println(tis.String())
	}

	// Same time zones
	fromStrings(
		"2014 Dec 14 20:00:00 UTC --- 2014 Dec 16 04:00:00 UTC",
		"2014 Dec 15 04:00:00 UTC --- 2014 Dec 15 08:00:00 UTC",
		"2014 Dec 15 08:00:00 UTC --- 2014 Dec 15 16:00:00 UTC",
		"2014 Dec 15 16:00:00 UTC --- 2014 Dec 15 20:00:00 UTC",
		"2014 Dec 15 20:00:00 UTC --- 2014 Dec 16 00:00:00 UTC",
	)

	// Mixed time zones
	fromStrings(
		"2014 Dec 14 20:00:00 PST --- 2014 Dec 16 04:00:00 PST",
		"2014 Dec 15 04:00:00 PST --- 2014 Dec 15 08:00:00 PST",
		"2014 Dec 15 08:00:00 PST --- 2014 Dec 15 16:00:00 PST",
		"2014 Dec 15 16:00:00 PST --- 2014 Dec 15 20:00:00 PST",
		"2014 Dec 15 20:00:00 UTC --- 2014 Dec 16 00:00:00 UTC",
	)

}

// Tests intervals exclude operation.
func TestIntervals_Exclude(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	// Interval 1
	t1 := t0.Add(-8 * time.Hour)
	t2 := t0.Add(+8 * time.Hour)

	ti1 := &TimeInterval{
		Name:   "One",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	tis := NewTimeIntervals(ti1)
	fmt.Println("Before exclude:")
	fmt.Println(tis)

	// Excluding Interval 1
	t1 = t0.Add(-5 * time.Hour)
	t2 = t0.Add(-4 * time.Hour)

	tiExcl := &TimeInterval{
		Name:   "Excl One",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	tis = tis.Exclude(tiExcl)
	fmt.Println("After exclude:")
	fmt.Println(tis)

	// Excluding Interval 2
	tiExcl.Ts = t0.Add(+5 * time.Hour)
	tiExcl.Te = t0.Add(+7 * time.Hour)

	tis = tis.Exclude(tiExcl)
	fmt.Println("After exclude:")
	fmt.Println(tis)

	// Excluding Interval 3, right edge
	tiExcl.Ts = t0.Add(+5 * time.Hour)
	tiExcl.Te = t0.Add(+16 * time.Hour)

	tis = tis.Exclude(tiExcl)
	fmt.Println("After exclude:")
	fmt.Println(tis)
}

// Tests intervals gaps operation.
func TestIntervals_Gaps(t *testing.T) {

	t0 := time.Date(2015, time.March, 15, 12, 0, 0, 0, time.UTC)
	var dt time.Duration = 30 * time.Minute
	dtmode := ""

	// Interval to exclude from
	t1 := t0.Add(-8 * time.Hour)
	t2 := t0.Add(+8 * time.Hour)

	ti := &TimeInterval{
		Name:   "MAIN",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Excluding Interval 1
	t1 = t0.Add(-5 * time.Hour)
	t2 = t0.Add(-4 * time.Hour)

	tie1 := &TimeInterval{
		Name:   "Invl A",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Excluding Interval 2
	t1 = t0.Add(-3 * time.Hour)
	t2 = t0.Add(-1 * time.Hour)

	tie2 := &TimeInterval{
		Name:   "Invl B",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Excluding Interval 3
	t1 = t0.Add(+2 * time.Hour)
	t2 = t0.Add(+6 * time.Hour)

	tie3 := &TimeInterval{
		Name:   "Invl C",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	fmt.Println("Analyze gaps for:")
	fmt.Println(NewTimeIntervals(ti, tie1, tie2, tie3))

	tis := NewTimeIntervals(tie1, tie2, tie3)
	tisRes := tis.AnalyzeRelativeTo(ti)
	fmt.Println("Gap analysys:")
	fmt.Println(tisRes)

	fmt.Println("Gaps Inner:")
	fmt.Println(NewTimeIntervals(tisRes.GapsInner()...))

	fmt.Println("Gap Left:", tisRes.GapLeft())

	fmt.Println("Gap Right:", tisRes.GapRight())

	fmt.Println("Gaps All:")
	fmt.Println(NewTimeIntervals(tisRes.Gaps()...))

	// Special cases

	// Intervals fully consume target interval
	fmt.Println()
	fmt.Println("Special case A: intervals fully consume target interval")

	// Excluding Interval 1
	t1 = t0.Add(-20 * time.Hour)
	t2 = t0.Add(0 * time.Hour)

	tie1 = &TimeInterval{
		Name:   "Excl One",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Excluding Interval 2
	t1 = t0.Add(0 * time.Hour)
	t2 = t0.Add(+20 * time.Hour)

	tie2 = &TimeInterval{
		Name:   "Excl Two",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	tis = NewTimeIntervals(tie1, tie2)
	fmt.Println("From:")
	fmt.Println(ti)
	fmt.Println("Exclude:")
	fmt.Println(tis)

	tisRes = tis.AnalyzeRelativeTo(ti)
	fmt.Println("Result composite:")
	fmt.Println(tisRes)

	fmt.Println("Gaps Inner:")
	fmt.Println(NewTimeIntervals(tisRes.GapsInner()...))

	fmt.Println("Gap Left:", tisRes.GapLeft())

	fmt.Println("Gap Right:", tisRes.GapRight())

	fmt.Println("Gaps All:")
	fmt.Println(NewTimeIntervals(tisRes.Gaps()...))

	// Special case

	// Intervals fully consume target interval
	fmt.Println()
	fmt.Println("Special case B: intervals lie outsite of target interval")

	// Excluding Interval 1
	t1 = t0.Add(-20 * time.Hour)
	t2 = t0.Add(-15 * time.Hour)

	tie1 = &TimeInterval{
		Name:   "Excl One",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	// Excluding Interval 2
	t1 = t0.Add(+15 * time.Hour)
	t2 = t0.Add(+20 * time.Hour)

	tie2 = &TimeInterval{
		Name:   "Excl Two",
		Ts:     t1,
		Te:     t2,
		Dt:     dt,
		DtMode: dtmode,
	}

	tis = NewTimeIntervals(tie1, tie2)
	fmt.Println("From:")
	fmt.Println(ti)
	fmt.Println("Exclude:")
	fmt.Println(tis)

	tisRes = tis.AnalyzeRelativeTo(ti)
	fmt.Println("Result composite:")
	fmt.Println(tisRes)

	fmt.Println("Gaps Inner:")
	fmt.Println(NewTimeIntervals(tisRes.GapsInner()...))

	fmt.Println("Gap Left:", tisRes.GapLeft())
	fmt.Println("Gap Right:", tisRes.GapRight())
	fmt.Println("Gap Left to Right:", tisRes.GapLeftRight())

	fmt.Println("Gaps All:")
	fmt.Println(NewTimeIntervals(tisRes.Gaps()...))

	// Special case

	// Intervals are empty, only target interval exists
	fmt.Println()
	fmt.Println("Special case C: intervals are empty only target interval exists")

	tis = NewTimeIntervals()
	fmt.Println("From:")
	fmt.Println(ti)
	fmt.Println("Exclude:")
	fmt.Println(tis)

	tisRes = tis.AnalyzeRelativeTo(ti)
	fmt.Println("Result composite:")
	fmt.Println(tisRes)

	fmt.Println("Gaps Inner:")
	fmt.Println(NewTimeIntervals(tisRes.GapsInner()...))

	fmt.Println("Gap Left:", tisRes.GapLeft())
	fmt.Println("Gap Right:", tisRes.GapRight())
	fmt.Println("Gap Left to Right:", tisRes.GapLeftRight())

	fmt.Println("Gaps All:")
	fmt.Println(NewTimeIntervals(tisRes.Gaps()...))

}
