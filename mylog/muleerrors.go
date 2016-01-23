package mylog

import (
	"mule/mybad"
)

// Println does a special check for the first arg
// to see if it is a MuleError; if so, it calls
// *MuleError.LogError() rather than .Error()
func (lg *Logger) Println(v ...interface{}) {
	if len(v) > 0 {
		if me, ok := v[0].(*mybad.MuleError); ok {
			v[0] = me.LogError()
		}
	}
	lg.Logger.Println(v...)
}

// Print does a special check for the first arg
// to see if it is a MuleError; if so, it calls
// *MuleError.LogError() rather than .Error()
func (lg *Logger) Print(v ...interface{}) {
	if len(v) > 0 {
		if me, ok := v[0].(*mybad.MuleError); ok {
			v[0] = me.LogError()
		}
	}
	lg.Logger.Print(v...)
}
