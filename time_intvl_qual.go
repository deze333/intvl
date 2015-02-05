// Basic qualities of time interval.
package intvl

import "time"

//------------------------------------------------------------
// Time Interval qualities
//------------------------------------------------------------

// Length of interval.
func (ti *TimeInterval) Len() time.Duration {
	return ti.Te.Sub(ti.Ts)
}

// UTC converts all times to UTC.
func (ti *TimeInterval) UTC() {

	ti.Ts = ti.Ts.UTC()
	ti.Te = ti.Te.UTC()

	for i, t := range ti.Times {
		ti.Times[i] = t.UTC()
	}
}
