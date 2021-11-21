package gokvshell

// To-Do make it generic for any datatype
type StringStringMap map[string]string

// A global key-value store for global access
var GlobalStore = make(StringStringMap)

// A transaction having its local store
type Transaction struct {
	store StringStringMap //local store
}

// A transaction stack for nested transactions
type TransactionStack struct {
	stack    []*Transaction // custom stack element
	topIndex int            // Top empty index of stack
	limit    int            // Limit to size of stack
}

// Initializes new transaction
func NewTransaction() *Transaction {
	return &Transaction{store: make(StringStringMap)}
}

// Initializes new transaction stack
func NewTransactionStack(limit int) *TransactionStack {
	return &TransactionStack{stack: make([]*Transaction, limit), topIndex: 0, limit: limit}
}

// Push a transaction to custom stack
func (transactionStack *TransactionStack) pushTransaction(transaction *Transaction) error {
	if transactionStack.topIndex == transactionStack.limit {
		return TransactionStackFullError{}
	}
	transactionStack.stack[transactionStack.topIndex] = transaction
	transactionStack.topIndex += 1
	return nil
}

// Push a top transaction from custom stack
func (transactionStack *TransactionStack) popTransaction() error {
	if transactionStack.topIndex == 0 {
		return TransactionStackEmptyError{}
	}
	// Release memory of unwanted poped transaction
	topTransaction, _ := transactionStack.topTransaction()
	topTransaction = nil

	transactionStack.topIndex -= 1
	return nil
}

// Get top transaction from custom stack
func (transactionStack *TransactionStack) topTransaction() (*Transaction, error) {
	if transactionStack.topIndex == 0 {
		return nil, TransactionStackEmptyError{}
	}
	return transactionStack.stack[transactionStack.topIndex-1], nil
}
