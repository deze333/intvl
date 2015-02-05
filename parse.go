// TimeInterval parsing functions.
package intvl

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/deze333/xparam"
)

//------------------------------------------------------------
// Time Interval parsing functions
//------------------------------------------------------------

// Parses time interval from supplied string.
// "layout --- layout"
// "layout --- layout @ dt" where dt is like 15m, 7h, 30d
func Parse_TimeInterval(layout, str string) (ti *TimeInterval, err error) {

	var dt time.Duration

	// Extract duration
	if strings.Index(str, "@") != -1 {
		strs := strings.Split(str, " @ ")
		if strs[1] != "" {
			dt, err = xparam.StringToDuration(strs[1])
			if err != nil {
				err = errors.New("Unparseable duration: " + err.Error())
				return
			}
			str = strs[0]
		}
	}

	// Process start/end time

	// Split into strings
	strs := strings.Split(str, " --- ")
	if len(strs) != 2 {
		err = errors.New("Unparseable string, doesn't match 'layout --- layout': " + str)
		return
	}

	// Parse time start
	var ts, te time.Time
	ts, err = time.Parse(layout, strs[0])
	if err != nil {
		return
	}

	// Parse time end
	te, err = time.Parse(layout, strs[1])
	if err != nil {
		return
	}

	// Verify ts < te
	if !ts.Before(te) {
		err = errors.New("Invalid TimeInterval: Ts must be before Te")
		return
	}

	ti = &TimeInterval{Ts: ts, Te: te, Dt: dt}
	return
}

// Parses into array of time intervals from supplied strings.
// "layout --- layout", "layout --- layout", ...
func ParseMany_TimeInterval(layout string, strs ...string) (tis []*TimeInterval, err error) {

	var ti *TimeInterval
	for i, str := range strs {

		if ti, err = Parse_TimeInterval(layout, str); err == nil {
			tis = append(tis, ti)
		} else {
			err = fmt.Errorf(
				"Unparseable time interval at position %v: %v",
				i, err)
			return
		}
	}

	return
}

// Parses into array of time intervals from supplied strings.
// "layout --- layout", "layout --- layout", ...
func Parse_TimeIntervals(layout string, strs ...string) (tis TimeIntervals, err error) {

	var tiis []*TimeInterval
	tiis, err = ParseMany_TimeInterval(layout, strs...)
	if err != nil {
		return
	}

	tis = NewTimeIntervals(tiis...)
	return
}
