package sql

// SELECT creates a selector for function chaining
// and eventual Compile()
//
// There is no error checking so you need to be sure
// to set everything you want before Compiling or else
// you'll probably get strange output
//
// Methods:
// FROM, WHERE, ORDER, Compile
func SELECT(cols ...string) *selector {
	return &selector{SelectCols: cols}
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
	return Conjunction{Joiner: "AND", Parts: conds}
}
func OR(conds ...Condition) Condition {
	return Conjunction{Joiner: "OR", Parts: conds}
}
