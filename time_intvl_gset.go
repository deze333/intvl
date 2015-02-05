// Get/set operations of time interval.
package intvl

import "time"

//------------------------------------------------------------
// Time Interval qualities
//------------------------------------------------------------

// Ts sets starting point of interval.
func (ti *TimeInterval) Start(t time.Time) {
	ti.ts = &t
	if ti.te != nil && (*ti.te).Before(*ti.ts) {
		panic("Invalid TimeInterval: end before start")
	}
	ti.Ts = t
}

// Ts sets starting point of interval.
func (ti *TimeInterval) End(t time.Time) {
	ti.te = &t
	if ti.ts != nil && (*ti.te).Before(*ti.ts) {
		panic("Invalid TimeInterval: end before start")
	}
	ti.Te = t
}
