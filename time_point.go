// TimePoint is a point inside an interval.
package intvl

import "time"

//------------------------------------------------------------
// Time Point model
//------------------------------------------------------------

type TimePoint struct {
	T    time.Time
	Idx  int
	Type string
}

//------------------------------------------------------------
// Time Point methods
//------------------------------------------------------------

func (p TimePoint) FindOwner(tis []*TimeInterval) *TimeInterval {
	for _, ti := range tis {
		if ti.idx == p.Idx {
			return ti
		}
	}

	return nil
}
