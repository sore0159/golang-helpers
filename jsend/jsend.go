// This package facilitates in serving JSON according to
// the jsend api spec found at
// http://labs.omniti.com/labs/jsend
package jsend

import (
	"mule/mybad"
)

const (
	MAXSIZE int64 = 1048576
)

var Check = mybad.BuildCheck("package", "jsend")
