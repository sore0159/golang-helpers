package group

// SQLGrouper is a full handler of SQLers
// See DeleteGrouper, SelectGrouper,
// InsertGrouper, UpdateGrouper
// for sub-handler types
type SQLGrouper interface {
	SQLTable() string
	PKCols() []string

	DeleteList() []SQLer

	SelectCols() []string
	New() SQLer

	UpdateCols() []string
	UpdateList() []SQLer

	InsertCols() []string
	InsertScanCols() []string
	InsertList() []SQLer
}

type DeleteGrouper interface {
	PKCols() []string
	SQLTable() string
	DeleteList() []SQLer
}

type SelectGrouper interface {
	SQLTable() string
	SelectCols() []string
	New() SQLer
}
type UpdateGrouper interface {
	SQLTable() string
	UpdateCols() []string
	PKCols() []string
	UpdateList() []SQLer
}
type InsertGrouper interface {
	SQLTable() string
	InsertCols() []string
	InsertScanCols() []string
	InsertList() []SQLer
}
