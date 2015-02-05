// TimeIntervals get/set operations.
package intvl

//------------------------------------------------------------
// Time Intervals get/set operations
//------------------------------------------------------------

// Counts all intervals that are marked as gaps of any kind.
func (tis TimeIntervals) GapCount() int {

	count := 0
	for _, ti := range tis {
		if ti.isGap {
			count++
		}
	}

	return count
}

// Finds all intervals that marked at "gap of any kind".
//
//	       [           target             ]
//		---------------------------------------->
//       <-- left]  [inner] [inner]  [right -->
func (tis TimeIntervals) Gaps() []*TimeInterval {

	var gaps []*TimeInterval
	for _, ti := range tis {
		if ti.isGap {
			gaps = append(gaps, ti)
		}
	}

	return gaps
}

// Finds all gaps that marked as "inner"
// lie inside of some target interval
// and don't touch neither of its edges.
//
//	       [             target             ]
//		---------------------------------------->
//            [ inner ] [ inner ]   [i]
func (tis TimeIntervals) GapsInner() []*TimeInterval {

	var gaps []*TimeInterval
	for _, ti := range tis {
		if ti.isGapInner {
			gaps = append(gaps, ti)
		}
	}

	return gaps
}

// Finds gap that marked as "extending left"
// past left edge of some target interval.
//
//	           [   target   ]
//		---------------------------------------->
//         <-- left ]
func (tis TimeIntervals) GapLeft() *TimeInterval {

	for _, ti := range tis {
		if ti.isGapLeft {
			return ti
		}
	}

	return nil
}

// Finds gap that marked as "extending right"
// past right edge of some target interval.
//
//	       [   target   ]
//		---------------------------------------->
//                   [ right -->
func (tis TimeIntervals) GapRight() *TimeInterval {

	for _, ti := range tis {
		if ti.isGapRight {
			return ti
		}
	}

	return nil
}

// Finds gap that marked as "left-to-right"
// relative to some target interval.
//
//	              [     target      ]
//		---------------------------------------->
//                [< left-to-right >]
func (tis TimeIntervals) GapLeftRight() *TimeInterval {

	for _, ti := range tis {
		if ti.isGapLeftRight {
			return ti
		}
	}

	return nil
}
