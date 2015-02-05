// Basic qualities of time intervals.
package intvl

//------------------------------------------------------------
// Time Intervals qualities
//------------------------------------------------------------

// UTC converts all times to UTC.
func (tis TimeIntervals) UTC() {

	for _, ti := range tis {
		ti.UTC()
	}
}
