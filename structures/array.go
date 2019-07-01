package structures

import (
	"errors"
	"fmt"
)

type Array interface {
	ToString() string
	Add(e interface{}) bool
	Get(i int) (interface{}, error)
	Set(i int, e interface{}) (bool, error)
	Search(e interface{}) int
	Delete(i int) bool
	Insert(i int, e interface{}) bool
	Size() int
}
type array struct {
	size     int
	capacity int
	arr      *[]interface{}
}

func (a *array) Size() int {
	return a.size
}
func (a *array) ToString() (str string) {
	str = fmt.Sprintf("Array: %v size:%d capacity:%d", *(a.arr), a.size, a.capacity)
	return
}
func (a *array) Add(e interface{}) bool {
	*a.arr = append(*a.arr, e)
	a.size++
	a.capacity = cap(*a.arr)
	return (*a.arr)[a.size-1] == e
}
func (a *array) Search(e interface{}) int {
	for i, v := range *a.arr {
		if v == e {
			return i
		}
	}
	return -1
}
func (a *array) Delete(index int) bool {
	if index >= 0 && index < a.size {
		*a.arr = append((*a.arr)[:index], (*a.arr)[index+1:]...)
		a.size -= 1
		return true
	}
	return false
}
func (a *array) Get(i int) (interface{}, error) {
	if i >= 0 && i < a.size {
		return (*a.arr)[i], nil
	}
	return 0, errors.New("index out of range")
}
func (a *array) Set(i int, e interface{}) (bool, error) {
	if i >= 0 && i < a.size {
		(*a.arr)[i] = e
		return (*a.arr)[i] == e, nil
	}
	return false, errors.New("index out of range")
}
func (a *array) Insert(i int, e interface{}) bool {
	if i >= 0 && i < a.size {
		*a.arr = append(*a.arr, (*a.arr)[a.size-1])
		for s := a.size; s > i; s-- {
			(*a.arr)[s] = (*a.arr)[s-1]
		}
		(*a.arr)[i] = e
		a.size++
		return true
	}
	return false
}
func NewArray(capacity int) *array {
	arr := make([]interface{}, 0, capacity)
	return &array{
		0,
		capacity,
		&arr,
	}
}
