package sql

// INSERT creates an inserter for function chaining
// and finally Compile()
// COLS must be called before VALS, and must be
// of the same length
//
// Methods:
//		required: Compile, COLS, VALS
//		optional: RETURN
// Example:
//      c := sql.INSERT("tweets").COLS("author", "content").VALS(
//			"mule", "HELLO WORLD"
//		).RETURN("id")
//		query, args := c.Compile()
func INSERT(table string) *inserter {
	return &inserter{Table: table}
}

// UPDATE creates an updater for function chaining
// and finally Compile()
// COLS must be called before TO, and must be
// of the same length
//
// Methods:
//		required: Compile, COLS, TO
//		optional: WHERE, RETURN
// Example:
//      c := sql.UPDATE("cards").COLS("number", "suit").TO(3, "spade").WHERE(sql.EQ("id", 2))
//		query, args := c.Compile()
func UPDATE(table string) *updater {
	return &updater{Table: table}
}

// DELETE creates a deletor for function chaining
// and finally Compile()
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
