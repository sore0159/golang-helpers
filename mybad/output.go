package mybad

import (
	"fmt"
	"strings"
)

// Error should output something suitable for end user display
// and a logfile for someone who doesn't know this is a MuleError
func (me *MuleError) Error() string {
	parts := make([]string, len(me.Layers)+1)
	for i, layer := range me.Layers {
		parts[i+1] = layer.Message
	}
	parts[0] = me.BaseError.Error()
	return strings.Join(parts, "||")
}

// LogError returns detailed information formatted for
// a single-line logfile, but still with full file and
// context information
func (me *MuleError) LogError() string {
	parts := make([]string, len(me.Layers)+1)
	for i, layer := range me.Layers {
		var badStr, ctxStr string
		if layer.BadlyFormed {
			badStr = "(BADFORMED)"
		}
		if len(layer.ContextItems) > 0 {
			subparts := make([]string, len(layer.ContextKeys))
			for i, key := range layer.ContextKeys {
				var item interface{}
				if len(layer.ContextItems) < i-1 {
					item = "MISSING CONTEXTITEM"
				} else {
					item = layer.ContextItems[i]
				}
				subparts[i] = fmt.Sprintf("\"%s\"=%+v", key, item)
			}
			ctxStr = fmt.Sprintf(", %s", strings.Join(subparts, ", "))
		}
		parts[i+1] = fmt.Sprintf("%s(%02d %s), msg=\"%s\"%s", badStr, layer.LineNum, layer.FileName, layer.Message, ctxStr)

	}
	parts[0] = me.BaseError.Error()
	return strings.Join(parts, "||")
}

// MuleError outputs a detailed report of the error
// including internal file info and line breaks.
// MuleError is _extremely_ verbose.
func (me *MuleError) MuleError() string {
	const start = "\n=========== BEGIN MULE ERROR ===========\n"
	const end = "============ END MULE ERROR ============\n"
	if me == nil {
		return fmt.Sprintf("%s<nil>%s", start, end)
	}
	base := fmt.Sprintf(
		"NUM LAYERS: %02d\n-------- BASE ERROR ---------\n%s\n", len(me.Layers), me.BaseError.Error())
	parts := make([]string, len(me.Layers))
	for i, layer := range me.Layers {
		header := fmt.Sprintf("--------  LAYER %02d   --------\n", i+1)
		//const tailer = "--------------------\n"
		parts[i] = fmt.Sprintf("%s%s", header, layer.MuleOutput())
	}
	layers := strings.Join(parts, "")

	return fmt.Sprintf("%s%s%s%s", start, base, layers, end)
}

// MuleOutput is verbosely formatted information not really
// suitable for a logfile or end user display
func (ly *Layer) MuleOutput() string {
	var badStr, ctxStr string
	if ly.BadlyFormed {
		badStr = "+++++++++(BADLY FORMED LAYER)+++++++++\n"
	}
	if len(ly.ContextKeys) > 0 {
		parts := make([]string, len(ly.ContextKeys))
		for i, key := range ly.ContextKeys {
			var item interface{}
			if len(ly.ContextItems) < i+1 {
				item = "++++++++++++++ MISSING CONTEXT ITEM +++++++++++++"
			} else {
				item = ly.ContextItems[i]
			}
			parts[i] = fmt.Sprintf("||CONTEXT KEY|| \"%s\"\n||CONTEXT ITEM||-------\n%+v\n-----------------------\n", key, item)
		}
		ctxStr = strings.Join(parts, "")
	} else {
		ctxStr = "No context given\n"
	}
	return fmt.Sprintf("%sFile:%s \nCaught on line %d (Call Depth: %02d)\nMessage: \"%s\" \n%s", badStr, ly.FileName, ly.LineNum, ly.CallDepth, ly.Message, ctxStr)
}
