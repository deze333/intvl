// Utilities for debugging.
package intvl

import (
	"bytes"
	"fmt"
)

//------------------------------------------------------------
// Debug utils
//------------------------------------------------------------

// DumpToString converts kv array into multi-line human readable string.
func dumpToString(name string, kvals []interface{}) string {
	var buf bytes.Buffer

	buf.WriteString("\n")
	buf.WriteString("-----------------------------------\n")
	buf.WriteString(name)
	buf.WriteString("\n")
	buf.WriteString("-----------------------------------\n")

	for i := 0; i < len(kvals); i += 2 {
		buf.WriteString(fmt.Sprint(kvals[i]))
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(kvals[i+1]))
		buf.WriteString("\n")
	}

	return buf.String()
}

// DumpToString converts kv array into single-line human readable string.
func dumpToStringLine(name string, kvals []interface{}) string {
	var buf bytes.Buffer

	buf.WriteString(name)
	buf.WriteString(" = {")

	for i := 0; i < len(kvals); i += 2 {
		buf.WriteString(fmt.Sprint(kvals[i]))
		if i+1 == len(kvals) {
			break
		}
		buf.WriteString(": ")
		buf.WriteString(fmt.Sprint(kvals[i+1]))
		if i+2 == len(kvals) {
			break
		}
		buf.WriteString(", ")
	}

	buf.WriteString(" }")
	return buf.String()
}
