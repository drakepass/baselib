package structures

import "errors"

type Stack interface {
	Push(interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	GetSize() int
	IsEmpty() bool
}
type arrayStack struct {
	size int
	arr  *[]interface{}
}

func (s *arrayStack) Push(e interface{}) {
	*s.arr = append(*s.arr, e)
	s.size++
}
func (s *arrayStack) Pop() (interface{}, error) {
	if bl := s.IsEmpty(); bl == true {
		err := errors.New("stack is empty")
		return nil, err
	}
	e := (*s.arr)[s.size-1]
	*s.arr = (*s.arr)[:s.size-1]
	s.size--
	return e, nil
}
func (s *arrayStack) Peek() (interface{}, error) {
	if bl := s.IsEmpty(); bl == true {
		err := errors.New("stack is empty")
		return nil, err
	}
	return (*s.arr)[s.size-1], nil
}
func (s *arrayStack) GetSize() int {
	return s.size
}
func (s *arrayStack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}
func NewStack(capacity int) Stack {
	arr := make([]interface{}, 0, capacity)
	stack := &arrayStack{
		0,
		&arr,
	}
	return stack
}
