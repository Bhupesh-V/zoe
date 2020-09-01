/*
An interactive shell that allows access to a transactional non-persistent in-memory key/value store.
*/

package main

import (
	"fmt"
	"os"
	"errors"
	"bufio"
	"strings"
)

/*GlobalStore hold the global transaction operations */
var GlobalStore = make(map[string]string)

/*Transaction points to a key:value store*/
type Transaction struct {
	store map[string]string
	next  *Transaction
}

/*TransactionStack is maintained as a list of active/suspended transactions */
type TransactionStack struct {
	top  *Transaction
	size int // more meta data can be saved like Stack limit etc.
}

/*PushTransaction creates new active transaction*/
func (s *TransactionStack) PushTransaction() {
	// Push a new Transaction, this is the current active transaction
	temp := Transaction{store : make(map[string]string)}
	temp.next = s.top
	s.top = &temp
	s.size++
}

/*PopTransaction creates delete a transaction*/
func (s *TransactionStack) PopTransaction() {
	// Pop the Transaction from stack, no longer active
	if s.top == nil {
		// basically stack underflow
		panic(errors.New("No Active Transactions"))
	} else {
		node := &Transaction{}
		s.top = s.top.next
		node.next = nil
		s.size--
	}
}

/*Peek active transaction*/
func (s *TransactionStack) Peek() *Transaction {
	return s.top
}

/*RollBackTransaction removes all keys SET within a transaction*/
func (s *TransactionStack) RollBackTransaction() {
	if s.top == nil {
		panic(errors.New("No Active Transaction"))
	} else {
		for key := range s.top.store {
			delete(s.top.store, key)
		}
	}
}

/*Get value of key from Store */
func Get(key string, T *TransactionStack) {
	ActiveTransaction := T.Peek()
	var node *Transaction
	var found bool = false
	if ActiveTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
		    fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	} else {
		node = ActiveTransaction
		for node != nil {
			if val, ok := node.store[key]; ok {
			    fmt.Printf("%s\n", val)
			    found = true
			}
			node = node.next
		}
		if !found {
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
	fmt.Printf("%s deleted", key)
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	items := &TransactionStack{}
	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		// split the text into operation strings
		userAction := strings.Fields(text)
		if userAction[0] == "BEGIN" {
			items.PushTransaction()
		} else if userAction[0] == "ROLLBACK" {
			items.RollBackTransaction()
		} else if userAction[0] == "END" {
			items.PopTransaction()
		} else if userAction[0] == "SET" {
			Set(userAction[1], userAction[2], items)
		} else if userAction[0] == "GET" {
			Get(userAction[1], items)
		} else if (userAction[0] == "DELETE") {
			Delete(userAction[1], items)
		} else if (userAction[0] == "COUNT") {
			Count(userAction[1], items)
		} else {
			fmt.Printf("ERROR: Unrecognised Operation %s", userAction[0])
		}
	}
}