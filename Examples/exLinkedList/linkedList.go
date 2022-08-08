package lindelist

import "fmt"

// wrong index error
var ErrWrongIndex = fmt.Errorf("wrong index")

// linked list struct
type IntLList struct {
	size int
	Head *IntNode
}

// linked list node
type IntNode struct {
	Value int
	Next  *IntNode
}

// create new linked list node
func NewIntNode(val int) *IntNode {
	return &IntNode{val, nil}
}

// create new int linked list
func NewLinkedList() *IntLList {
	return &IntLList{0, nil}
}

// return size of linked list
func (l IntLList) Size() int {
	return l.size
}

// get element of linked list by index
func (l IntLList) Get(idx int) (*IntNode, error) {
	if idx < 0 || idx > l.size {
		return nil, ErrWrongIndex
	}

	node := l.Head

	for i := 1; i <= idx; i++ {
		node = node.Next
	}

	return node, nil

}

// set number into linked list by index
func (l *IntLList) Set(num, idx int) error {

	node, err := l.Get(idx)
	if err != nil {
		return err
	}

	node.Value = num

	return nil
}

// add element at start linked list
func (l *IntLList) AddStart(num int) {

	newNode := NewIntNode(num)

	if l.Head == nil {
		l.Head = newNode
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.size++

}

// insert element into idx position
func (l *IntLList) Insert(num, idx int) error {

	if idx == 0 {
		l.AddStart(num)
		return nil
	}

	node, err := l.Get(idx - 1)
	if err != nil {
		return err
	}

	newNode := NewIntNode(num)
	newNode.Next = node.Next
	node.Next = newNode
	l.size++

	return nil

}

// remove linked list element by idx
func (l *IntLList) Remove(idx int) error {

	node, err := l.Get(idx - 1)
	if err != nil {
		return err
	}

	node.Next = node.Next.Next
	l.size--

	return nil
}

func (l IntLList) Print() {

	if l.Head == nil {
		fmt.Println("linked list is empty")
		return
	}

	node := l.Head

	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}
