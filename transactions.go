package main

// To-Do make it generic for any datatype
type StringStringMap map[string]string

// A global key-value store for global access
var GlobalStore = make(StringStringMap)

// A transaction having its local store
type Transaction struct {
	store StringStringMap //local store
	next  *Transaction
}

// A transaction stack for nested transactions
type TransactionStack struct {
	top   *Transaction
	size  int
	limit int
}
