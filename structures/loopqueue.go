package structures

import (
	"errors"
	"fmt"
)

type Loopqueue interface {
	GetFront() (interface{}, error)
	GetSize() int
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	IsEmpty() bool
	ToString()
}
type queue struct {
	front    int
	tail     int
	size     int
	loopq    *[]interface{}
	capacity int
}

func (q *queue) GetFront() (interface{}, error) {
	if q.IsEmpty() == false {
		return (*q.loopq)[q.front], nil
	}
	err := errors.New("the queue is empty")
	return nil, err
}
func (q *queue) GetSize() int {
	return q.size
}
func (q *queue) Enqueue(e interface{}) error {
	if (q.tail+1)%q.capacity != q.front {
		(*q.loopq)[q.tail] = e
		q.tail = (q.tail + 1) % q.capacity
		q.size++
		return nil
	} else {
		return q.reSize(e)
		//return q.reEnqueue(e) another way
	}
}
func (q *queue) reSize(e interface{}) error {
	sli := make([]interface{}, q.capacity*2, q.capacity*2)
	for i := 0; i < q.size; i++ {
		sli[i] = (*q.loopq)[(q.front+i)%q.capacity]
	}
	sli[q.size] = e
	q.front = 0
	q.size++
	q.tail = q.size
	q.capacity = q.capacity * 2
	*q.loopq = sli
	return nil
}
func (q *queue) reEnqueue(e interface{}) error {
	if q.tail < q.front {
		sli := make([]interface{}, q.capacity*2, q.capacity*2)
		sli1 := make([]interface{}, 0)
		sli1 = append(sli1, (*q.loopq)[q.front:]...)
		sli1 = append(sli1, (*q.loopq)[:(q.tail)]...)
		sli1 = append(sli1, e)
		copy(sli, sli1)
		*q.loopq = sli
		q.front = 0
		q.size++
		q.tail = q.size
		q.capacity = cap(*q.loopq)
		return nil
	} else {
		sli := make([]interface{}, 0, q.capacity*2)
		sli = append(sli, (*q.loopq)[q.front:q.tail]...)
		sli = append(sli, e)
		sli1 := make([]interface{}, q.capacity*2, q.capacity*2)
		copy(sli1, sli)
		*q.loopq = sli1
		q.front = 0
		q.size++
		q.tail = q.size
		q.capacity = cap(*q.loopq)
		return nil
	}

}
func (q *queue) ToString() {
	fmt.Println(*q.loopq)
	fmt.Println("front: ", q.front, "tail:", q.tail, " size: ", q.size, " capacity: ", q.capacity)
	return
}
func (q *queue) Dequeue() (interface{}, error) {
	if q.front != q.tail {
		v := (*q.loopq)[q.front]
		q.front = (q.front + 1) % q.capacity
		q.size--
		return v, nil
	}
	return nil, errors.New("the queue is empty")
}
func (q *queue) IsEmpty() bool {
	if q.front != q.tail {
		return false
	}
	return true
}

func NewLoopQueue(capacity int) Loopqueue {
	loopq := make([]interface{}, capacity, capacity)
	loopqueue := &queue{
		0,
		0,
		0,
		&loopq,
		capacity,
	}
	return loopqueue
}
