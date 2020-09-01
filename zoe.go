/*
An interactive shell that allows access to a transactional, in-memory 
key/value store.
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
		panic(errors.New("stack underflow"))
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

/*func (s *TransactionStack) Display() {
	fmt.Println("---Transaction Stack---")
	var node *Transaction
	if s.top == nil {
		panic(errors.New("Stack Underflow !!"))
	} else {
		fmt.Printf("top ->")
		node = s.top
		for node != nil {
			fmt.Printf("\t%v\n", node.value)
			node = node.next
		}
	}
}
*/
/*Get value of key from Store */
func Get(key string, ActiveTransaction *TransactionStack) {
	s := ActiveTransaction.Peek()
	if s == nil {
		if val, ok := GlobalStore[key]; ok {
		    fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	} else {
		if val, ok := s.store[key]; ok {
		    fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	}
}

/*Set key to value */
func Set(key string, value string, ActiveTransaction *TransactionStack) {
	// Get key:value store from active transaction
	s := ActiveTransaction.Peek()
	if s == nil {
		GlobalStore[key] = value
	} else {
		s.store[key] = value
	}
}

/*Count returns the number of keys that have been set to the specified value.*/
func Count(value string){
	var count int = 0
	ActiveTransaction := &TransactionStack{}
	s := ActiveTransaction.Peek()
	for _, v := range s.store {
		if v == value {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}

/*Delete value from Store */
func Delete(key string, ActiveTransaction *TransactionStack) {
	s := ActiveTransaction.Peek()
	if s == nil {
		delete(GlobalStore, key)
	} else {
		delete(s.store, key)
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
		} else if userAction[0] == "END" {
			items.PopTransaction()
		} else if userAction[0] == "SET" {
			Set(userAction[1], userAction[2], items)
		} else if userAction[0] == "GET" {
			Get(userAction[1], items)
		} else if (userAction[0] == "DELETE") {
			Delete(userAction[1], items)
		} else if (userAction[0] == "COUNT") {
			Count(userAction[1])
		} else {
			fmt.Printf("ERROR: Unrecognised Operation %s", userAction[0])
		}
	}
}