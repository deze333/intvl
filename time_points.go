// TimePoints is a sortable array of points.
package intvl

//------------------------------------------------------------
// Time Points model
//------------------------------------------------------------

type TimePoints []TimePoint

//------------------------------------------------------------
// Time Points sort
//------------------------------------------------------------

func (s TimePoints) Len() int {
	return len(s)
}

func (s TimePoints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s TimePoints) Less(i, j int) bool {
	return s[i].T.Before(s[j].T)
}
