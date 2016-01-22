package mybad

// This is probably all rubbish but mabye I might need to variably change
// runtime depth check level
type Checker func(depth int, err error, msg string, ctx ...interface{}) (my *MuleError, bad bool)

// DeepCheck is for if you want to add Check to some other subroutine and
// adjust the calldepth checked for file line and add context automatically
func DeepCheck(depth int, err error, msg string, ctx ...interface{}) (my *MuleError, bad bool) {
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
	layer := MakeLayer(msg, ctx, depth)
	me.Layers = append(me.Layers, layer)
	return me, true
}

// ExtendedCheck is a convenience function for extending the
// FullExtendChecker lets you build function calls that can automatically
// add context to error checks and optionally increase the call depth
// checking
func FullExtendChecker(f Checker, addDepth int, addCtx ...interface{}) Checker {
	return func(depth int, err error, msg string, ctx ...interface{}) (my *MuleError, bad bool) {
		return f(depth+addDepth, err, msg, append(ctx, addCtx))
	}
}
