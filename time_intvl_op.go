// Operations in time intervals.
package intvl

import "time"

//------------------------------------------------------------
// Time Interval operations
//------------------------------------------------------------

// Trims interval from left leaving what's on the right of t.
//	    [     this      ]
// 	--------|----------------->
// 	        t
//	     xxx
// Result:
// 	        [    ti     ]
func (this *TimeInterval) TrimLeft(t time.Time) (ti *TimeInterval) {

	// t is outside of interval, nothing happens
	if t.Sub(this.Ts) <= 0 || t.Sub(this.Te) >= 0 {
		return this
	}

	ti = this.CloneMin()
	ti.Ts = t

	return
}

// Trims interval from left leaving what's on the left of t.
//	    [      this      ]
// 	---------|----------------->
// 	         t
//	          xxxxxxxxxxxx
// Result:
// 	    [ ti ]
func (this *TimeInterval) TrimRight(t time.Time) (ti *TimeInterval) {

	// t is outside of interval, nothing happens
	if t.Sub(this.Ts) <= 0 || t.Sub(this.Te) >= 0 {
		return this
	}

	ti = this.CloneMin()
	ti.Te = t

	return
}

// Exclude excludes other interval.
func (this *TimeInterval) Exclude(other *TimeInterval) []*TimeInterval {

	res := []*TimeInterval{}

	// No overlap at all
	if other.IsBefore(this) || other.IsAfter(this) {
		return append(res, this.CloneMin())
	}

	// Other fully covers
	if this.IsContainedBy(other) {
		return res
	}

	// Other is fully inside
	if this.IsContains(other) {
		switch {

		case this.Ts.Equal(other.Ts):
			// Both starts equal, only right gap

			right := this.CloneMin()
			right.Ts = other.Te
			right.Te = this.Te
			res = append(res, right)

		case this.Te.Equal(other.Te):
			// Both ends equal, only left gap

			left := this.CloneMin()
			left.Ts = this.Ts
			left.Te = other.Ts
			res = append(res, left)

		default:
			// Other padded from left and right

			left := this.CloneMin()
			left.Ts = this.Ts
			left.Te = other.Ts

			right := this.CloneMin()
			right.Ts = other.Te
			right.Te = this.Te

			res = append(res, left)
			res = append(res, right)
		}

		return res
	}

	// Other ends inside
	if other.IsEndsInside(this) {

		ti := this.CloneMin()
		ti.Ts = other.Te

		return append(res, ti)
	}

	// Other starts inside
	if other.IsStartsInside(this) {

		ti := this.CloneMin()
		ti.Ts = this.Ts
		ti.Te = other.Ts

		return append(res, ti)
	}

	// Other
	return res
}

// Splits interval into shorter subintervals of len dur.
// Extends last subinterval to reach the end of source interval.
//
//  Size:          |     |
//
// 	Source:        [     |     |     | ]
// 	Result:        [ dur | dur | dur+  ]
//
// 	Source:        [ ]
// 	Result:        [ ]
func (ti *TimeInterval) Split(dur time.Duration) (tis []*TimeInterval) {

	if ti.Len() == 0 {
		return
	}

	if ti.Len() <= dur {
		tis = append(tis, ti)
		return
	}

	num := ti.Len() / dur
	var i time.Duration
	var ts, te time.Time

	ts = ti.Ts
	for i = 0; i < num; i++ {

		if i < num-1 {
			te = ts.Add(dur)
		} else {
			te = ti.Te
		}

		sub := ti.Clone()
		sub.Ts = ts
		sub.Te = te
		tis = append(tis, sub)

		ts = te
	}

	return
}

// Splits interval into shorter subintervals of len dur,
// starting from interval's right boundary and moving leftwards.
// Extends last subinterval to full dur size.
//
//	Direction:    <------------------|
// 	Source:            [ |     |     ]
// 	Result:        [ dur | dur | dur ]
func (ti *TimeInterval) SplitExtend_Leftwards(dur time.Duration) (tis []*TimeInterval) {

	if ti.Len() == 0 {
		return
	}

	num := ti.Len() / dur
	if ti.Len()%dur != 0 {
		num++
	}

	var i time.Duration
	var ts, te time.Time

	te = ti.Te
	for i = 0; i < num; i++ {

		ts = te.Add(-dur)

		sub := ti.Clone()
		sub.Ts = ts
		sub.Te = te
		tis = append(tis, sub)

		te = ts
	}

	return
}

// Splits interval into shorter subintervals of len dur,
// starting from interval's left boundary and moving rightwards.
// Extends last subinterval to full dur size.
//
//	Direction:     |----------------->
// 	Source:        [     |     | ]
// 	Result:        [ dur | dur | dur ]
func (ti *TimeInterval) SplitExtend_Rightwards(dur time.Duration) (tis []*TimeInterval) {

	if ti.Len() == 0 {
		return
	}

	num := ti.Len() / dur
	if ti.Len()%dur != 0 {
		num++
	}

	var i time.Duration
	var ts, te time.Time

	ts = ti.Ts
	for i = 0; i < num; i++ {

		te = ts.Add(dur)

		sub := ti.Clone()
		sub.Ts = ts
		sub.Te = te
		tis = append(tis, sub)

		ts = te
	}

	return
}
