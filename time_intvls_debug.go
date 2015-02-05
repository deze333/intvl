// TimeIntervals debug.
//		0 : [_]                      : TimeInterval:
//		1 :     [_]                  : TimeInterval:
//		2 :         [___]            : TimeInterval:
//		3 :         [___]            : TimeInterval:
//		4 :           [_____]        : TimeInterval:
//		5 :               [___]      : TimeInterval:
//		6 :                     [_]  : TimeInterval:

package intvl

import (
	"bytes"
	"fmt"
	"sort"
	"time"
)

//------------------------------------------------------------
// Time Intervals debug
//------------------------------------------------------------

func (tis TimeIntervals) Dump() string {

	var buf bytes.Buffer
	for i, ti := range tis {
		buf.WriteString(fmt.Sprint(i))
		buf.WriteString(" : ")
		buf.WriteString("idx = ")
		buf.WriteString(fmt.Sprint(ti.idx))
		buf.WriteString(" : ")
		buf.WriteString(ti.Ts.String())
		buf.WriteString(" --- ")
		buf.WriteString(ti.Te.String())
		buf.WriteString(" : ")
		buf.WriteString(fmt.Sprint(ti.isGap, " ",
			ti.isGapLeft, " ", ti.isGapInner, " ",
			ti.isGapRight, " ", ti.isGapLeftRight))
		buf.WriteString("\n")
	}

	return buf.String()
}

func (tis TimeIntervals) String() string {

	// Check if some pointers point at same interval
	for i, ti := range tis {
		for j, ti2 := range tis {
			if i != j && ti == ti2 {
				return "WARNING: Cannot output time intervals: some time intervals pointers point as same intervals. All pointers must be unique."
			}
		}
	}

	// Arrange intervals by start
	tis = tis.SortByTs()

	// Update indexes
	for i, ti := range tis {
		ti.idx = i
	}

	// Aggregate all points
	points := make([]TimePoint, len(tis)*2)
	for i, ti := range tis {

		points[2*i] = TimePoint{
			T:    ti.Ts,
			Type: "s",
			Idx:  ti.idx,
		}
		points[2*i+1] = TimePoint{
			T:    ti.Te,
			Type: "e",
			Idx:  ti.idx,
		}
	}

	// Sort all points
	sort.Sort(TimePoints(points))

	// Aggregate equal points together
	type aggregation struct {
		t      time.Time
		points []TimePoint
	}

	aggrs := []aggregation{}
	i := 0
	var prev TimePoint
	for _, point := range points {

		if prev.T.IsZero() {
			aggr := aggregation{t: point.T, points: []TimePoint{point}}
			aggrs = append(aggrs, aggr)

		} else if point.T == prev.T {
			aggrs[i].points = append(aggrs[i].points, point)

		} else {
			aggr := aggregation{t: point.T, points: []TimePoint{point}}
			aggrs = append(aggrs, aggr)
			i++
		}
		prev = point
	}

	// Allocate line for each interval
	lines := make([][]byte, len(tis))
	for i, _ := range tis {
		lines[i] = []byte{}
	}

	// Func that returns interval inner filler char
	getIntvlSymb := func(ti *TimeInterval) byte {
		switch {
		case ti.isGapLeft:
			return '<'
		case ti.isGapRight:
			return '>'
		case ti.isGap:
			return 'G'
		default:
			return '_'
		}
	}

	// For each aggregated time position
	for i, aggr := range aggrs {

		// Hits are interval that received [ or ] marks
		hits := []int{}
		for _, p := range aggr.points {
			ti := p.FindOwner(tis)
			if ti == nil {
				continue
			}
			idx := ti.idx
			hits = append(hits, idx)
			var symb byte = ' '
			switch p.Type {
			case "s":
				symb = '['
			case "e":
				symb = ']'
			}
			lines[idx] = append(lines[idx], symb)
		}

		// Misses are intervals that didn't receive [ or ] marks
		misses := []int{}
		for _, ti := range tis {
			isHit := false
			for _, hit := range hits {
				if hit == ti.idx {
					isHit = true
					break
				}
			}
			if !isHit {
				misses = append(misses, ti.idx)
			}
		}

		// If reached last aggregation then add space and exit
		if i == len(aggrs)-1 {
			for _, idx := range misses {
				ti := tis.Find(idx)
				lines[ti.idx] = append(lines[ti.idx], ' ')
			}
			break
		}

		// Lookahead: next aggregation
		naggr := aggrs[i+1]

		// Plot hits
		for _, idx := range hits {
			ti := tis.Find(idx)
			var symb byte
			if naggr.t.After(ti.Te) {
				symb = ' '
			} else {
				symb = getIntvlSymb(ti)
			}
			lines[idx] = append(lines[idx], symb)
		}

		// Plot misses
		for _, idx := range misses {
			ti := tis.Find(idx)
			var symb byte
			if aggr.t.Before(ti.Ts) || naggr.t.After(ti.Te) {
				symb = ' '
			} else {
				symb = getIntvlSymb(ti)
			}
			lines[idx] = append(lines[idx], symb)
			lines[idx] = append(lines[idx], symb)
		}

		// Fill insides
		/*
			for _, ti := range tis {
				var symb byte
				if aggr.t.Before(ti.Ts) || aggr.t.After(ti.Te) {
					symb = ' '
				} else {
					symb = '_'
				}
				lines[ti.idx] = append(lines[ti.idx], symb)
			}
		*/
	}

	// Format string for indexes
	format := func(l int) string {
		if l < 10 {
			return "%d"
		} else if l < 100 {
			return "%2d"
		}
		return "%d"
	}(len(lines))

	// Build output
	var buf bytes.Buffer

	for i, line := range lines {
		ti := tis.Find(i)

		buf.WriteString(fmt.Sprintf(format, i))
		buf.WriteString(" : ")
		buf.WriteString(string(line))
		buf.WriteString(" : ")
		if ti != nil {
			buf.WriteString(dumpToStringLine("", ti.Dump()))
		} else {
			buf.WriteString("TimeInterval <NOT FOUND>")
		}

		buf.WriteString("\n")
	}

	return buf.String()
}
