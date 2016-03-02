// Package sql handles query generation and argument management.
// It is not for collecting data or talking to a db, just
// for organizing data into sql.
//
// Usage example:
//	import "mule/mydb/sql"
//
//	func() {
//		s := sql.SELECT("name", "date").FROM("birthdays").WHERE(
//			sql.AND(sql.EQ("month", 12), sql.IN("status", 1,2,3))
//		).ORDER("date DESC")
//
//		query, args := s.Compile()
//		rows, err := db.Query(query, args...)
package sql
