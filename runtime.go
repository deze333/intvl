// Package level settings.
package intvl

//------------------------------------------------------------
// Runtime model & instance
//------------------------------------------------------------

type Runtime struct {
	timeLayoutDebug string
}

var _runtime *Runtime

//------------------------------------------------------------
// Initialize
//------------------------------------------------------------

// Initializes package runtime.
func init() {
	_runtime = &Runtime{
		timeLayoutDebug: TIME_LAYOUT_DEBUG,
	}
}

//------------------------------------------------------------
// Setters
//------------------------------------------------------------

// Sets current runtime time output format.
func Runtime_TimeLayout_Debug(layout string) {
	_runtime.timeLayoutDebug = layout
}
