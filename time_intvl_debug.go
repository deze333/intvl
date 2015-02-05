// Methods for human readable debugging of time interval.
package intvl

import (
	"fmt"
	"time"
)

//------------------------------------------------------------
// Debug methods
//------------------------------------------------------------

// LenHuman converts length of interval to a human readable
// string, for example: 5d 3h 34m.
func (ti *TimeInterval) LenHuman() string {

	dt := ti.Len()
	d := time.Duration(dt.Hours() / 24)

	// Hours left
	dt -= d * time.Hour * 24
	h := time.Duration(dt.Hours())

	// Minutes left
	dt -= h * time.Hour
	m := time.Duration(dt.Minutes())

	return fmt.Sprintf("%3dd %2dh %02dm", d, h, m)
}

// String converts interval to string.
func (ti *TimeInterval) String() string {
	return dumpToStringLine("TimeInterval", ti.Dump())
}

// Dump provides raw kv array of fields and their values.
func (ti *TimeInterval) Dump() []interface{} {

	dump := []interface{}{
		fmt.Sprintf("[%v --- %v] dt = %v len = %v",
			ti.Ts.UTC().Format(_runtime.timeLayoutDebug),
			ti.Te.UTC().Format(_runtime.timeLayoutDebug),
			ti.Dt, ti.Te.Sub(ti.Ts)),
	}

	if ti.DtMode != "" {
		dump = append(dump, fmt.Sprintf("dtMode = %v", ti.DtMode))
	}

	if ti.Name != "" {
		dump = append(dump, ti.Name)
	}

	return dump
}
