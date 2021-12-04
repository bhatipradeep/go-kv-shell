package gokvshell

// To-Do make it generic for any datatype
type StringStringMap map[string]string

// A global key-value store for global access
var GLOBAL_STORE = make(StringStringMap)

// A transaction having its local store
type Transaction struct {
	store StringStringMap //local store
}

// A transaction stack for nested transactions
type TransactionStack struct {
	stack []*Transaction // custom stack element
	size  int            // size of filled stack
	limit int            // Limit to size of stack
}

// Initializes new transaction
func NewTransaction() *Transaction {
	return &Transaction{store: make(StringStringMap)}
}

// Initializes new transaction stack
func NewTransactionStack(limit int) *TransactionStack {
	return &TransactionStack{stack: make([]*Transaction, limit), size: 0, limit: limit}
}

// Push a transaction to custom stack
func (transactionStack *TransactionStack) PushTransaction(transaction *Transaction) error {
	if transactionStack.size == transactionStack.limit {
		return TransactionStackFullError{}
	}
	transactionStack.stack[transactionStack.size] = transaction
	transactionStack.size += 1
	return nil
}

// Push a top transaction from custom stack
func (transactionStack *TransactionStack) PopTransaction() error {
	if transactionStack.size == 0 {
		return TransactionStackEmptyError{}
	}

	// To-Do
	// Release memory of unwanted poped transaction
	// topTransaction, _ := transactionStack.topTransaction()
	// topTransaction = nil

	transactionStack.size -= 1
	return nil
}

// Get top transaction from custom stack
func (transactionStack *TransactionStack) TopTransaction() (*Transaction, error) {
	if transactionStack.size == 0 {
		return nil, TransactionStackEmptyError{}
	}
	return transactionStack.stack[transactionStack.size-1], nil
}

// Get second transaction from top from custom stack
func (transactionStack *TransactionStack) PreTopTransaction() (*Transaction, error) {
	if transactionStack.size < 2 {
		return nil, TransactionNotFoundError{}
	}
	return transactionStack.stack[transactionStack.size-2], nil
}
