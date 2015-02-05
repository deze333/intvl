// TimeIntervals is an array of intervals.
package intvl

import "sort"

//------------------------------------------------------------
// Time Intervals model
//------------------------------------------------------------

type TimeIntervals []*TimeInterval

// NewTimeIntervals creates sorted array of intervals.
func NewTimeIntervals(tis ...*TimeInterval) TimeIntervals {

	// Normalize all to UTC
	for _, ti := range tis {
		ti.UTC()
	}

	// Create and sort
	intvls := TimeIntervals(tis)
	sort.Sort(TimeIntervals_ByTs(intvls))

	// Discover gaps
	//...

	// Set inner indexes
	for i, ti := range intvls {
		ti.idx = i
	}

	return intvls
}

// Clones time intervals.
func (tis TimeIntervals) Clone() TimeIntervals {
	clone := make([]*TimeInterval, len(tis))
	copy(clone, tis)
	return clone
}

// Finds time interval by its index position.
func (tis TimeIntervals) Find(idx int) *TimeInterval {
	for _, ti := range tis {
		if ti.idx == idx {
			return ti
		}
	}
	return nil
}

// Adds interval to array of intervals.
func (tis TimeIntervals) Add(ti *TimeInterval) TimeIntervals {

	return NewTimeIntervals(append(tis, ti)...)
}

// Sorts intervals by time start.
func (tis TimeIntervals) SortByTs() TimeIntervals {
	sortable := TimeIntervals_ByTs(tis)
	sort.Sort(sortable)
	return TimeIntervals(sortable)
}

// Sorts intervals by time end.
func (tis TimeIntervals) SortByTe() TimeIntervals {
	sortable := TimeIntervals_ByTe(tis)
	sort.Sort(sortable)
	return TimeIntervals(sortable)
}

//------------------------------------------------------------
// Time Intervals sort by Ts model
//------------------------------------------------------------

type TimeIntervals_ByTs []*TimeInterval

func (s TimeIntervals_ByTs) Len() int {
	return len(s)
}

func (s TimeIntervals_ByTs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s TimeIntervals_ByTs) Less(i, j int) bool {
	return s[i].Ts.Before(s[j].Ts)
}

//------------------------------------------------------------
// Time Intervals sort by Te model
//------------------------------------------------------------

type TimeIntervals_ByTe []*TimeInterval

func (s TimeIntervals_ByTe) Len() int {
	return len(s)
}

func (s TimeIntervals_ByTe) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s TimeIntervals_ByTe) Less(i, j int) bool {
	return s[i].Te.Before(s[j].Te)
}
