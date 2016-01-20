package mydb

type Querier interface {
	// SQLVal returns a pointer to a scanner corresponding to the given string
	SQLValScan(string) interface{}
	// SQLVal returns the value for psq to read corresponding to the given string
	SQLVal(string) interface{}
}
