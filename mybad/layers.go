package mybad

import (
	"fmt"
	"runtime"
)

type Layer struct {
	CallDepth    int
	FileName     string
	LineNum      int
	Message      string
	ContextKeys  []string
	ContextItems []interface{}
	BadlyFormed  bool
}

func MakeLayer(msg string, ctx []interface{}, extraDepth int) *Layer {
	var badlyFormed bool
	l := len(ctx)
	if l%2 == 1 {
		l += 1
	}
	hl := l / 2
	conKeys := make([]string, 0, hl)
	conItems := make([]interface{}, 0, hl)
	if msg == "" {
		msg = "BLANK MESSAGE"
		badlyFormed = true
	}
	callDepth := extraDepth + 2
	_, fName, lineNum, ok := runtime.Caller(callDepth)
	if !ok {
		badlyFormed = true
		conKeys = append(conKeys, "BAD EXTRA CALL DEPTH")
		conItems = append(conItems, extraDepth)
		for i := callDepth - 1; i > 0 && !ok; i-- {
			_, fName, lineNum, ok = runtime.Caller(i)
		}
	}
	layer := &Layer{
		CallDepth:    callDepth,
		FileName:     fName,
		LineNum:      lineNum,
		Message:      msg,
		ContextKeys:  conKeys,
		ContextItems: conItems,
		BadlyFormed:  badlyFormed,
	}

	if l > 0 {
		layer.AddContext(ctx)
	}
	return layer
}

func (ly *Layer) AddContext(ctx []interface{}) {
	if ln := len(ctx); ln%2 == 1 {
		ly.BadlyFormed = true
		ctx = append(ctx[:ln-1], "UNMATCHED CTX ITEM:", ctx[ln-1])
	}
	for i := 0; i < len(ctx); i += 2 {
		if key, ok := ctx[i].(string); !ok {
			ly.AppendCtxItem("BAD CTX KEY:", ctx[i])
			ly.AppendCtxItem("UNMATCHED CTX ITEM:", ctx[i+1])
			ly.BadlyFormed = true
		} else {
			ly.AppendCtxItem(key, ctx[i+1])
		}
	}
}

func (ly *Layer) AppendCtxItem(key string, item interface{}) {
	if key == "" {
		key = "BLANKKEY"
		ly.BadlyFormed = true
	}
	if maker, ok := item.(func() interface{}); ok {
		key = fmt.Sprintf("%s(MADE)", key)
		item = maker()
	}
	if str, ok := item.(string); ok {
		item = fmt.Sprintf("\"%s\"", str)
	}
	ly.ContextKeys = append(ly.ContextKeys, key)
	ly.ContextItems = append(ly.ContextItems, item)
}
