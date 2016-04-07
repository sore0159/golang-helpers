package sql

import "errors"

// INSERT creates an inserter for function chaining
// and finally Compile() or Args()
// Compile() is for query+args, args is for quickly
// parsing args without doing all the string building
// for prepared statements
//
// COLS must be called before VALS, and must be
// of the same length
// Optionally VALS may be omitted for querygens
// not tracking arguments (mass interts).  Nils
// will be provided
//
// Methods:
//		required: Compile, COLS
//		optional: RETURN, VALS
// Example:
//      c := sql.INSERT("tweets").COLS("author", "content").VALS(
//			"mule", "HELLO WORLD"
//		).RETURN("id", "likes")
//		query, args := c.Compile()
func INSERT(table string) *inserter {
	return &inserter{Table: table}
}

// UPDATE creates an updater for function chaining
// and finally Compile() or Args()
// COLS must be called before TO, and must be
// of the same length. TO can be omitted if you
// are not collecting arguments; nils will be supplied.
// Compile() is for query+args, Args() is for quickly
// parsing args without doing all the string building
// for prepared statements
//
// Methods:
//		required: Compile, COLS
//		optional: WHERE, RETURN, TO
// Example:
//      c := sql.UPDATE("cards").COLS("number", "suit").TO(3, "spade").WHERE(sql.EQ("id", 2))
//		query, args := c.Compile()
//
// Example querygen for multi-updates:
//      c2 := sql.UPDATE("days").COLS("count").WHERE(EQ("id", nil))
//		query, _ = c.Compile()
func UPDATE(table string) *updater {
	return &updater{Table: table}
}

// DELETE creates a deletor for function chaining
// and finally Compile()
// Compile() is for query+args, args is for quickly
// parsing args without doing all the string building
// for prepared statements
//
// Methods:
//		required: Compile
//		optional: WHERE
// Example:
//      c := sql.DELETE("chairs").WHERE(sql.OR(
//			sql.IN("legs", 1, 2, 3),
//			sql.COMP("age", ">", 1),
//		))
//		query, args := c.Compile()
func DELETE(table string) *dropper {
	return &dropper{Table: table}
}

// SELECT creates a selector for function chaining
// and finally Compile()
// If COLS is not called, * is selected
//
// Methods:
//		required: Compile
//		optional: COLS, WHERE, ORDER
// Example:
//      c := sql.SELECT("shirts").COLS("size", "fashion").WHERE(sql.AND(
//			sql.EQ("clean", true), sql.EQ("color", "red"),
//		))
//		query, args := c.Compile()
func SELECT(table string) *selector {
	return &selector{Table: table}
}

// P is for vardic functions to not type switch strings
type P struct {
	Col string
	Val interface{}
}

// Condititons are for WHERE clauses, and here are
// the collection of functions for creating them

// Note the IN function can't be called with "list..." for
// vals without you first converting "list" to []interface{}
func IN(col string, vals ...interface{}) Condition {
	return In{ColName: col, InVals: vals}
}
func COMP(col, op string, val interface{}) Condition {
	return Compare{Operator: op, ColName: col, ColVal: val}
}
func EQ(col string, val interface{}) Condition {
	return Compare{Operator: "=", ColName: col, ColVal: val}
}
func AND(conds ...Condition) Condition {
	return MakeConjunction("AND", conds...)
}
func OR(conds ...Condition) Condition {
	return MakeConjunction("OR", conds...)
}

// AllEQ really is my most common use case
func AllEQ(ps ...P) Condition {
	return AND(EQList(ps...)...)
}

// EQList is if we want to AND a bunch of EQs with an IN or something
func EQList(ps ...P) []Condition {
	parts := make([]Condition, len(ps))
	for i, p := range ps {
		parts[i] = EQ(p.Col, p.Val)
	}
	return parts
}

// Convert2P is a helper if you REALLY just don't care
// about type saftey and want to use ...interface{}
func Convert2P(list []interface{}) ([]P, error) {
	l := len(list)
	if l%2 != 0 {
		return nil, errors.New("odd interface args for convert2P")
	}
	them := make([]P, l/2)
	for i := 0; i < l/2; i += 1 {
		col, ok := list[2*i].(string)
		if !ok {
			return nil, errors.New("nonstring interface arg for convert2P")
		}
		them[i].Col = col
		them[i].Val = list[2*i+1]
	}
	return them, nil
}
