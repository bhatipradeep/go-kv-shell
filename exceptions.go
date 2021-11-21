package gokvshell

type TransactionStackFullError struct{}

func (transactionStackFullError TransactionStackFullError) Error() string {
	return "Transaction Stack is Full !"
}

type TransactionStackEmptyError struct{}

func (transactionStackFullEmpty TransactionStackEmptyError) Error() string {
	return "Transaction Stack is Empty !"
}
