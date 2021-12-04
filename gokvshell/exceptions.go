package gokvshell

type TransactionStackFullError struct{}

func (transactionStackFullError TransactionStackFullError) Error() string {
	return "Transaction Stack is Full !"
}

type TransactionStackEmptyError struct{}

func (transactionStackFullEmpty TransactionStackEmptyError) Error() string {
	return "Transaction Stack is Empty !"
}

type KeyNotPresentError struct{}

func (keyNotPresentError KeyNotPresentError) Error() string {
	return "Key Not Present in store !"
}

type TransactionNotFoundError struct{}

func (transactionNotFoundError TransactionNotFoundError) Error() string {
	return "Transaction Not found !"
}

type InvalidSetArgumentsError struct{}

func (invalidSetArgumentsError InvalidSetArgumentsError) Error() string {
	return "Enter valid set argument. Ex. SET <key> <value>"
}

type InvalidGetArgumentsError struct{}

func (invalidGetArgumentsError InvalidGetArgumentsError) Error() string {
	return "Enter valid set argument. Ex. GET <key>"
}
