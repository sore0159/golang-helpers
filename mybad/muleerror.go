package mybad

// MuleError is a handler for error chains in my programs.
// Often when something goes wrong a chain of actions
// fail, and MuleError allows me to add context as we
// pass the buck up the line.
//
// MuleErrors should probably not be constructed manually:
// They are for when an error is caught, not generated
// (even if a generated error is immediately caught).
type MuleError struct {
	BaseError error
	Layers    []*Layer
}

func NewMuleError(err error) *MuleError {
	return &MuleError{
		BaseError: err,
		Layers:    make([]*Layer, 0, 2),
	}
}

// You can add context to a MuleError as part of the Checking, or
// if you need more space or want complicated context generation
// after you know something is wrong, you can AddContext before
// passing it along.
func (me *MuleError) AddContext(ctx ...interface{}) {
	if l := len(me.Layers); l == 0 {
		layer := MakeLayer("ATTEMPTED TO ADD CONTEXT TO LAYERLESS ERROR", ctx, 0)
		layer.BadlyFormed = true
		me.Layers = []*Layer{layer}
	} else {
		me.Layers[l-1].AddContext(ctx)
	}
}

// BaseIs is a helper for checking if a captured error is of
// a certain kind: like FileNotExist or whatever
// Usage Example:
//
// file, err := os.Open("shoppinglist.txt")
// if my, bad := mybad.Check(err, "shopping list file creation"); bad {
//		if my.BaseIs(os.ErrNotExist) {
//			makeShoppingFile()
//		} else {
//			return my
//		}
// }
//
// You can currently also just access my.BaseError manually
// if you want
func (me *MuleError) BaseIs(err error) bool {
	return me.BaseError == err
}

func (me *MuleError) Grab(m2 *MuleError) {
	me.Layers = append(me.Layers, m2.Layers...)
}
