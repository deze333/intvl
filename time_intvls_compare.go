// TimeIntervals comparison.
package intvl

//------------------------------------------------------------
// Time Intervals comparison.
//------------------------------------------------------------

// Compares two time intervals for full matching.
func (tis TimeIntervals) IsEqual(other TimeIntervals) bool {

	if len(tis) != len(other) {
		return false
	}

	// Each interval must match exactly
	for i := 0; i < len(tis); i++ {
		if !tis[i].IsEqual(other[i]) {
			return false
		}
	}

	return true
}
