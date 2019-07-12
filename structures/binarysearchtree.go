package structures

import (
	"fmt"
)

type Bst interface {
	Add(interface{})
	Delete(interface{})
	DeleteMax()
	DeleteMin() *node
	Find(interface{}) bool
	ToString()
	IsEmpty() bool
	GetSize() int
}

type node struct {
	e           interface{}
	left, right *node
}

type binarySearchTree struct {
	size    int
	element *node
}

func (B *binarySearchTree) Add(e interface{}) {
	if B.size == 0 {
		B.element = &node{
			e,
			nil,
			nil,
		}
		B.size++
	} else {
		B.doADD(B.element, e) //two ways for traversal using recursion
		//doAdd(B,B.element,e)
	}
}

func (B *binarySearchTree) doADD(b *node, e interface{}) *node {
	if b == nil {
		B.size++
		return &node{
			e,
			nil,
			nil,
		}
	}
	if e.(int) == (b.e).(int) {
		return b
	}
	if e.(int) < (b.e).(int) {
		b.left = B.doADD(b.left, e)
	} else {
		b.right = B.doADD(b.right, e)
	}
	return b
}

func doAdd(B *binarySearchTree, b *node, e interface{}) {
	if e.(int) == (b.e).(int) {
		return
	}
	if e.(int) < (b.e).(int) {
		if b.left == nil {
			b.left = &node{
				e,
				nil,
				nil,
			}
			B.size++
			return
		} else {
			doAdd(B, b.left, e)
			return
		}
	}
	if e.(int) > (b.e).(int) {
		if b.right == nil {
			b.right = &node{
				e,
				nil,
				nil,
			}
			B.size++
			return
		} else {
			doAdd(B, b.right, e)
			return
		}
	}
}

func (B *binarySearchTree) FindMin(el *node) *node {
	if el.left == nil {
		return el
	}
	return B.FindMin(el.left)
}

func (B *binarySearchTree) FindMax(el *node) *node {
	if el.right == nil {
		return el
	}
	return B.FindMax(el.right)
}

func (B *binarySearchTree) deleteone(node *node, e interface{}) *node {
	if (node.e).(int) > e.(int) {
		node.left = B.deleteone(node.left, e)
		return node
	} else if node.e.(int) < e.(int) {
		node.right = B.deleteone(node.right, e)
		return node
	} else {
		if node.left != nil && node.right != nil {
			minNode := B.FindMin(node.right)
			_ = B.toDeleteMin(node.right)
			if node.right != minNode {
				minNode.right = node.right
			}
			minNode.left = node.left
			return minNode
		} else if node.left == nil && node.right == nil {
			B.size--
			return nil
		} else if node.left == nil {
			newnode := node.right
			B.size--
			return newnode
		} else {
			newnode := node.left
			B.size--
			return newnode
		}
	}
}

func (B *binarySearchTree) Delete(e interface{}) {
	if !B.Find(e) {
		return
	}
	B.element = B.deleteone(B.element, e)
}

func (B *binarySearchTree) toDeleteMax(el *node) {
	if el.right.right == nil {
		if el.right.left == nil {
			el.right = nil
		} else {
			el.right = el.right.left
		}
		B.size--
		return
	}
	B.toDeleteMax(el.right)
}

func (B *binarySearchTree) DeleteMax() {
	if B.element.right == nil {
		//panic("this root of bst is the max value")
		B.toDeleteMax(B.element.left)
	} else {
		B.toDeleteMax(B.element)
	}

}

func (B *binarySearchTree) toDeleteMin(el *node) *node {
	if el.left != nil {
		el.left = B.toDeleteMin(el.left)
		return el
	} else {
		if el.right == nil {
			B.size--
			return nil
		} else {
			newnode := el.right
			B.size--
			return newnode
		}
	}

}

func (B *binarySearchTree) DeleteMin() *node {
	if B.element.left == nil {
		//panic("this root of bst is the min value")
		return nil
	}
	ret := B.FindMin(B.element)
	B.element = B.toDeleteMin(B.element)
	return ret
}

func (B *binarySearchTree) Find(e interface{}) bool {
	return B.tofind(B.element, e)
}

func (B *binarySearchTree) tofind(element *node, e interface{}) bool {
	if element == nil {
		return false
	}
	if element.e == e {
		return true
	}
	if element.e.(int) > e.(int) {
		return B.tofind(element.left, e)
	} else {
		return B.tofind(element.right, e)
	}
}

//recursion to preorder traversal
func toshow(b *node, depth int) {
	depth++
	str := ""
	for i := 0; i < depth; i++ {
		str = str + "-"
	}
	fmt.Println(str, b.e)
	if b.left != nil {
		toshow(b.left, depth)
	}
	if b.right != nil {
		toshow(b.right, depth)
	}
}

//none recursion to preorder traversal
func (B *binarySearchTree) toshowNR(b *node) {
	stack := NewStack(100)
	stack.Push(b)
	for !stack.IsEmpty() {
		el, err := stack.Pop()
		if err != nil {
			panic(err)
		}
		em := el.(*node)
		fmt.Println(em.e)
		if em.right != nil {
			stack.Push(em.right)
		}
		if em.left != nil {
			stack.Push(em.left)
		}
	}
}

func (B *binarySearchTree) toshowBreadth(b *node) {
	//queue := structures.NewLoopQueue(10)
	queue := NewLoopQueue(10)
	err := queue.Enqueue(b)
	if err != nil {
		panic(err)
	}
	for !queue.IsEmpty() {
		e, _ := queue.Dequeue()
		el := e.(*node)
		fmt.Println(el.e)
		if el.left != nil {
			queue.Enqueue(el.left)
		}
		if el.right != nil {
			queue.Enqueue(el.right)
		}
	}
}

//two way for preorder traversal and the last is another way using breadth first
func (B *binarySearchTree) ToString() {
	toshow(B.element, 0)
	fmt.Println("-----------------------")
	//B.toshowNR(B.element)
	//fmt.Println("-----------------------------")
	B.toshowBreadth(B.element)
	fmt.Println("bst size: ", B.size)
}

func (B *binarySearchTree) IsEmpty() bool {
	return B.size == 0
}

func (B *binarySearchTree) GetSize() int {
	return B.size
}

func NewBst() Bst {
	return &binarySearchTree{
		0,
		nil,
	}
}
