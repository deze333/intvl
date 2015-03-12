// TimeIntervals operations.
package intvl

//------------------------------------------------------------
// Time Intervals operations
//------------------------------------------------------------

// Gaps analyses time intervals within given bounds.
// Returns new time intervals marked with results of
// analysis as meta information.
func (tis TimeIntervals) AnalyzeRelativeTo(bounds *TimeInterval) (res TimeIntervals) {

	// Find all gaps
	gaps := NewTimeIntervals(bounds)

	for _, ti := range tis {
		gaps = gaps.Exclude(ti)
	}

	// Special case:
	// target interval returned untouched
	if len(gaps) == 1 && gaps[0].IsEqual_TsTe(bounds) {
		ti := gaps[0]
		ti.isGap = true
		ti.isGapLeftRight = true
		all := append(tis, ti)
		res = NewTimeIntervals(all...)
		return
	}

	// Mark edge and non-edge gaps
	for _, ti := range gaps {
		ti.isGap = true

		switch {
		case ti.Ts.Equal(bounds.Ts):
			ti.isGapLeft = true
		case ti.Te.Equal(bounds.Te):
			ti.isGapRight = true
		default:
			ti.isGapInner = true
		}
	}

	// Combine gaps and no gaps
	all := append(tis, gaps...)
	res = NewTimeIntervals(all...)
	return

	/*
		// Find edge gaps: left and right
		idxLeft := -1
		idxRight := -1
		for i, ti := range gaps {
			if ti.Ts.Equal(bounds.Ts) {
				idxLeft = i
			}
			if ti.Te.Equal(bounds.Te) {
				idxRight = i
			}
			if idxLeft != -1 && idxRight != -1 {
				break
			}
		}

		if idxLeft != -1 {
			left = gaps[idxLeft]
		}

		if idxRight != -1 {
			left = gaps[idxRight]
		}

		// Inner gaps
		for i, ti := range gaps {
			if i != idxLeft && i != idxRight {
				inner = append(inner, ti)
			}
		}
	*/

	return
}

// Excludes runs exlude of excl interval against
// each of tis intervals and returns resulting modified tis.
func (tis TimeIntervals) Exclude(excl *TimeInterval) (res TimeIntervals) {

	res = TimeIntervals{}
	for _, ti := range tis {
		rems := ti.Exclude(excl)

		// Extra bug protection:
		// Remove intervals with zero length
		for _, rem := range rems {
			if rem.Len() != 0 {
				res = append(res, rem)
			}
		}
	}

	return
}

// Analyzes intervals for duplicates and overlaps.
// Returns: origs - intervals that are clean,
// dups - intervals that duplicate originals,
// overs - parts of intervals that overlap with originals
// and need to be deleted.
// Name of each interval is preserved and can be used as meta data.
// Each overlap name contains both names, comma-separated.
func (tis TimeIntervals) AnalyzeOverlaps() (origs, dups, overs []*TimeInterval) {

	nextOrigIdx := 0
	for i, ti := range tis {

		if i != nextOrigIdx {
			continue
		}

		origs = append(origs, ti)

		// Look for each next intervals
		for j, tiNext := range tis[i+1:] {

			// Stop if next doesn't overlap
			if tiNext.Ts.Sub(ti.Te) >= 0 {
				nextOrigIdx = i + j + 1

				/*
					if j > 0 {
						fmt.Println()
					}
				*/

				break
			}

			// Duplicate ?
			if tiNext.IsEqual(ti) {
				dups = append(dups, tiNext)

				/*
					if j == 0 {
						fmt.Println(i, "ORIGINAL")
						fmt.Println(ti)
					}

					fmt.Println(j, "DUP")
					fmt.Println(tiNext)
				*/

				continue
			}

			// Overlap ? Next starts before current ends
			if ti.Te.Sub(tiNext.Ts) > 0 {
				over := tiNext.TrimRight(ti.Te)
				over.Name = ti.Name + "," + tiNext.Name
				overs = append(overs, over)

				/*
					if j == 0 {
						fmt.Println(i, "ORIGINAL")
						fmt.Println(ti)
					}

					fmt.Println(j, "TRIM")
					fmt.Println(tiNext)
				*/
			}

		}
	}

	return
}
