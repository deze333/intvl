// Comparison of time intervals.
package intvl

//------------------------------------------------------------
// Time Interval comparison
//------------------------------------------------------------

// IsEqual verifies full equality of two intervals.
func (ti *TimeInterval) IsEqual(other *TimeInterval) bool {
	if !ti.Ts.Equal(other.Ts) {
		return false
	}
	if !ti.Te.Equal(other.Te) {
		return false
	}
	if ti.Dt != other.Dt {
		return false
	}
	if ti.DtMode != other.DtMode {
		return false
	}

	return true
}

// IsEqual_TsTe checks if [Ts:Te) are equal for both intervals.
func (ti *TimeInterval) IsEqual_TsTe(other *TimeInterval) bool {
	return ti.Ts.Equal(other.Ts) && ti.Te.Equal(other.Te)
}

// IsEqual_TsTeDt checks if [Ts:Te) at given Dt are equal for both intervals.
func (ti *TimeInterval) IsEqual_TsTeDt(other *TimeInterval) bool {
	return ti.Dt == other.Dt &&
		ti.Ts.Equal(other.Ts) && ti.Te.Equal(other.Te)
}

// IsInside checks if this interval is fully inside other.
//     [   other  ]
//       [ this ]
func (this *TimeInterval) IsContainedBy(other *TimeInterval) bool {

	gapL := this.Ts.Sub(other.Ts)
	gapR := other.Te.Sub(this.Te)

	return gapL >= 0 && gapR >= 0
}

// IsInside checks if other interval is fully inside this one.
//     [   this   ]
//       [ other ]
func (this *TimeInterval) IsContains(other *TimeInterval) bool {

	gapL := other.Ts.Sub(this.Ts)
	gapR := this.Te.Sub(other.Te)

	return gapL >= 0 && gapR >= 0
}

// IsBefore checks if this interval is located before other.
//     [   this   ]
//                  [ other ]
func (this *TimeInterval) IsBefore(other *TimeInterval) bool {

	return other.Ts.Sub(this.Te) >= 0
}

// IsAfter checks if this interval is located after other.
//               [   this   ]
//     [ other ]
func (this *TimeInterval) IsAfter(other *TimeInterval) bool {

	return this.Ts.Sub(other.Te) >= 0
}

// IsStartsBefore checks if this interval starts inside other.
//     [   this
//       [ other
func (this *TimeInterval) IsStartsBefore(other *TimeInterval) bool {

	return other.Ts.Sub(this.Ts) >= 0
}

// IsStartsAfter checks if this interval starts inside other.
//               [   this
//     [ other ]
func (this *TimeInterval) IsStartsAfter(other *TimeInterval) bool {

	return this.Ts.Sub(other.Te) >= 0
}

// IsStartsInside checks if this interval starts inside other.
//         [   this   ]
//     [ other ]
func (this *TimeInterval) IsStartsInside(other *TimeInterval) bool {

	return this.Ts.Sub(other.Ts) >= 0 && other.Te.Sub(this.Ts) >= 0
}

// IsEndsInside checks if this interval ends inside other.
//     [   this   ]
//            [ other ]
func (this *TimeInterval) IsEndsInside(other *TimeInterval) bool {

	return this.Te.Sub(other.Ts) >= 0 && other.Te.Sub(this.Te) >= 0
}

// IsLeftAdjacent checks if this interval ends at other's start.
//     [   this   ]
//                [ other ]
func (this *TimeInterval) IsLeftAdjacent(other *TimeInterval) bool {

	return this.Te.Equal(other.Ts)
}

// IsRightAdjacent checks if this interval starts at other's end.
//             [   this   ]
//     [ other ]
func (this *TimeInterval) IsRightAdjacent(other *TimeInterval) bool {

	return other.Te.Equal(this.Ts)
}
