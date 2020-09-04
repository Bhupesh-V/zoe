/*
Zoe is an interactive shell that allows access to a transactional non-persistent in-memory key/value store.
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

/*GlobalStore holds the (global) variables*/
var GlobalStore = make(map[string]string)

/*Map string:string*/
type Map = map[string]string

/*Transaction points to a key:value store*/
type Transaction struct {
	store Map 			// every transaction has its own local store
	next  *Transaction
}

/*TransactionStack is maintained as a list of active/suspended transactions */
type TransactionStack struct {
	top  *Transaction
	size int 			// more meta data can be saved like Stack limit etc.
}

/*PushTransaction creates new active transaction*/
func (ts *TransactionStack) PushTransaction() {
	// Push a new Transaction, this is the current active transaction
	temp := Transaction{store : make(Map)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

/*PopTransaction creates delete a transaction*/
func (ts *TransactionStack) PopTransaction() {
	// Pop the Transaction from stack, no longer active
	if ts.top == nil {
		// basically stack underflow
		fmt.Printf("ERROR: No Active Transactions\n")
	} else {
		node := &Transaction{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
	}
}

/*Peek returns active transaction*/
func (ts *TransactionStack) Peek() *Transaction {
	return ts.top
}

/*RollBackTransaction clears all keys SET within a transaction*/
func (ts *TransactionStack) RollBackTransaction() {
	if ts.top == nil {
		fmt.Printf("ERROR: No Active Transaction\n")
	} else {
		// this is optimized by the Go1.11 compiler :o
		for key := range ts.top.store {
			delete(ts.top.store, key)
		}
	}
}

/*Commit write(SET) changes to the store with TranscationStack scope
Also write changes to disk/file, if data needs to persist after the shell closes
*/
func (ts *TransactionStack) Commit() {
	ActiveTransaction := ts.Peek()
	if ActiveTransaction != nil {
		for key, value := range ActiveTransaction.store {
			GlobalStore[key] = value
			if ActiveTransaction.next != nil {
				// update the parent transaction
				ActiveTransaction.next.store[key] = value
			}
		}
	} else {
		fmt.Printf("INFO: Nothing to commit\n")
	}
}

/*Get value of key from Store (WIP)*/
func Get(key string, T *TransactionStack) {
	ActiveTransaction := T.Peek()
	var node *Transaction
	if ActiveTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
		    fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	} else {
		node = ActiveTransaction
		if val, ok := node.store[key]; ok {
		    fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	}
}

/*Set key to value */
func Set(key string, value string, T *TransactionStack) {
	// Get key:value store from active transaction
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		GlobalStore[key] = value
	} else {
		ActiveTransaction.store[key] = value
	}
}

/*Count returns the number of keys that have been set to the specified value.*/
func Count(value string, T *TransactionStack){
	var count int = 0
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		for _, v := range GlobalStore {
			if v == value {
				count++
			}
		}
	} else {
		for _, v := range ActiveTransaction.store {
			if v == value {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)
}

/*Delete value from Store */
func Delete(key string, T *TransactionStack) {
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		delete(GlobalStore, key)
	} else {
		delete(ActiveTransaction.store, key)
	}
	fmt.Printf("%s deleted\n", key)
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	items := &TransactionStack{}
	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		// split the text into operation strings
		operation := strings.Fields(text)
		switch operation[0] {
		case "BEGIN": 		items.PushTransaction()
		case "ROLLBACK": 	items.RollBackTransaction()
		case "COMMIT": 		items.Commit(); items.PopTransaction()
		case "END": 		items.PopTransaction()
		case "SET": 		Set(operation[1], operation[2], items)
		case "GET": 		Get(operation[1], items)
		case "DELETE": 		Delete(operation[1], items)
		case "COUNT": 		Count(operation[1], items)
		case "STOP": 		os.Exit(0)
		default:
			fmt.Printf("ERROR: Unrecognised Operation %s\n", operation[0])
		}
	}
}