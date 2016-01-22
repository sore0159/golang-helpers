package mybad

// Check is a simple default error check to be directly called
// that returns a MuleError and a bool indicating if any error
// was found.
// Example Usage:
//
// dinner, err := oven.Cook(food)
// if my, bad :=  mybad.Check(err, "OPPS!", "oventemp", oven.Temp); bad {
//		my.AddContext("recoverable", dinner.StillEdible)
//		return my
// }
//
// Context arguments should be alternating string keys
// and referenced items.  For lazy resolution of context
// items, pass a func() interface{} and Check will
// call the func and use the resulting interface as the
// context item, tagging the context key with (MADE)
// Context args can be part of the initial check or can
// be added later with AddContext(...interface{})
func Check(err error, msg string, ctx ...interface{}) (my *MuleError, bad bool) {
	if err == nil {
		return nil, false
	}
	var me *MuleError
	var ok bool
	if me, ok = err.(*MuleError); !ok {
		me = NewMuleError(err)
	} else {
		if me == nil {
			return nil, false
		}
	}
	layer := MakeLayer(msg, ctx, 0)
	me.Layers = append(me.Layers, layer)
	return me, true
}

// BuildCheck builds a check function that automatically
// includes context of your choosing
// Example Usage:
//
// var localCheck = mybad.BuildCheck("package", "mydinner", "time", func() interface{} { return time.Now() })
//
// Check functions built this way may then be used in the same fashion
// as mybad.Check
func BuildCheck(autoCtx ...interface{}) func(err error, msg string, ctx ...interface{}) (my *MuleError, bad bool) {
	return func(err error, msg string, ctx ...interface{}) (my *MuleError, bad bool) {
		if err == nil {
			return nil, false
		}
		var me *MuleError
		var ok bool
		if me, ok = err.(*MuleError); !ok {
			me = NewMuleError(err)
		} else {
			if me == nil {
				return nil, false
			}
		}
		ctx = append(ctx, autoCtx...)
		layer := MakeLayer(msg, ctx, 0)
		me.Layers = append(me.Layers, layer)
		return me, true
	}
}
