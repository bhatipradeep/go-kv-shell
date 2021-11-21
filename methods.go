package gokvshell

// Get value for a key
func (transactionStack *TransactionStack) get(key string) (string, error) {
	topTransaction, err := transactionStack.topTransaction()
	if err != nil {
		if val, ok := GLOBAL_STORE[key]; ok {
			return val, nil
		} else {
			return "", KeyNotPresentError{}
		}
	} else {
		return topTransaction.store[key], nil
	}
}

// Set value of a key
func (transactionStack *TransactionStack) set(key string, val string) {
	topTransaction, err := transactionStack.topTransaction()
	if err != nil {
		GLOBAL_STORE[key] = val
	} else {
		topTransaction.store[key] = val
	}

}

// delete key-value
func (transactionStack *TransactionStack) delete(key string) {
	topTransaction, err := transactionStack.topTransaction()
	if err != nil {
		delete(GLOBAL_STORE, key)
	} else {
		topTransaction.store[key] = ""
	}
}

// Commit changes made during the transaction
func (transactionStack *TransactionStack) commit() error {
	topTransaction, err := transactionStack.topTransaction()
	if err != nil {
		return err
	} else {
		for key, val := range topTransaction.store {
			if val == "" {
				delete(GLOBAL_STORE, key)
			} else {
				GLOBAL_STORE[key] = val
			}
		}
		preTopTransaction, preTopTransactionErr := transactionStack.preTopTransaction()
		if preTopTransactionErr == nil {
			for key, val := range topTransaction.store {
				if val == "" {
					delete(preTopTransaction.store, key)
				} else {
					preTopTransaction.store[key] = val
				}

			}
		}
		return nil
	}
}

// Rollback changes made during transaction
func (transactionStack *TransactionStack) rollback() error {
	topTransaction, err := transactionStack.topTransaction()
	if err != nil {
		return err
	} else {
		for key, _ := range topTransaction.store {
			delete(topTransaction.store, key)
		}
		return nil
	}
}
