// TimeInterval represents slice of time line
// within [Ts, Te] where both Ts, Te are inclusive.
// Dt represents time granularity for the interval.
// DtMode is "" for homogenous distribution
// or "DISCRETE", in which case Times[] contain
// specific moments in the interval.
package intvl

import "time"

//------------------------------------------------------------
// Time Interval model
//------------------------------------------------------------

type TimeInterval struct {
	Ts     time.Time     `bson:"ts,omitempty"        json:"ts,omitempty"`
	Te     time.Time     `bson:"te,omitempty"        json:"te,omitempty"`
	Dt     time.Duration `bson:"dt,omitempty"        json:"dt,omitempty"`
	DtMode string        `bson:"dtMode,omitempty"    json:"dtMode,omitempty"`
	Times  []time.Time   `bson:"times,omitempty"     json:"times,omitempty"`

	// Meta information
	Name string `bson:"name,omitempty"     json:"name,omitempty"`

	// Not exposed
	ts *time.Time
	te *time.Time

	// TimeIntervals inner index
	idx int

	// Gap analysis data
	isGap          bool
	isGapInner     bool
	isGapLeft      bool
	isGapRight     bool
	isGapLeftRight bool

	// NOTE When adding/renaming fields, don't forget to update .Clone()
}

//------------------------------------------------------------
// Methods
//------------------------------------------------------------

// Clone produces deep copy of interval.
func (ti *TimeInterval) Clone() *TimeInterval {

	clone := &TimeInterval{
		Ts:             ti.Ts,
		Te:             ti.Te,
		Dt:             ti.Dt,
		DtMode:         ti.DtMode,
		Name:           ti.Name,
		isGap:          ti.isGap,
		isGapInner:     ti.isGapInner,
		isGapLeft:      ti.isGapLeft,
		isGapRight:     ti.isGapRight,
		isGapLeftRight: ti.isGapLeftRight,
	}

	if len(ti.Times) >= 0 {
		times := make([]time.Time, len(ti.Times))
		copy(times, ti.Times)
		clone.Times = times
	}

	return clone
}

// CloneMin produces minimalistic deep copy of interval.
// Only basic fields are preserved. No meta data copied.
func (ti *TimeInterval) CloneMin() *TimeInterval {

	clone := &TimeInterval{
		Ts:     ti.Ts,
		Te:     ti.Te,
		Dt:     ti.Dt,
		DtMode: ti.DtMode,
	}

	if len(ti.Times) >= 0 {
		times := make([]time.Time, len(ti.Times))
		copy(times, ti.Times)
		clone.Times = times
	}

	return clone
}
