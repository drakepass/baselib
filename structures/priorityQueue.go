package structures

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// typed in another file loopqueue
// type Loopqueue interface {
//		GetFront() 				(interface{},error)
//		GetSize()				int
//		Enqueue(interface{})	error
//		Dequeue()				(interface{},error)
//		IsEmpty()				bool
// }
type pqueue struct {
	e     interface{}
	queue []interface{}
	size  int
}

func (p *pqueue) Enqueue(e interface{}) error {
	p.queue = append(p.queue, e)
	p.size++
	p.shiftup(p.size - 1)
	return nil
}

func (p *pqueue) GetFront() (interface{}, error) {
	if p.size == 0 {
		err := errors.New("there is no element in the priority queue")
		return nil, err
	}
	return p.queue[0], nil
}

func (p *pqueue) GetSize() int {
	return p.size
}

func (p *pqueue) IsEmpty() bool {
	return p.size == 0
}

func (p *pqueue) Dequeue() (interface{}, error) {
	if p.size == 0 {
		err := errors.New("there is no element in the priority queue")
		return nil, err
	}
	ret, err := p.GetFront()
	p.queue[0] = p.queue[p.size-1]
	p.queue = p.queue[:p.size-1]
	p.size--
	p.shiftdown(0)
	return ret, err
}

func (p *pqueue) shiftdown(k int) {
	left := 2*k + 1
	right := 2*k + 2
	if left > p.size-1 {
		return
	}
	if right < p.size && p.compare(p.queue[left], p.queue[right]) == -1 {

		if p.compare(p.queue[k], p.queue[right]) == -1 {
			p.queue[k], p.queue[right] = p.queue[right], p.queue[k]
			p.shiftdown(right)
		}
	} else {
		if p.compare(p.queue[k], p.queue[left]) == -1 {
			p.queue[k], p.queue[left] = p.queue[left], p.queue[k]
			p.shiftdown(left)
		}
	}
}

func (p *pqueue) shiftup(k int) {
	if k == 0 {
		return
	}
	fk := k / 2
	if p.compare(p.queue[fk], p.queue[k]) == -1 {
		p.queue[fk], p.queue[k] = p.queue[k], p.queue[fk]
		p.shiftup(fk)
	}
}
func (p *pqueue) compare(a, b interface{}) int {
	v := reflect.TypeOf(a)
	switch v.Kind() {
	case reflect.Int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return -1
		}

	case reflect.String:
		return strings.Compare(a.(string), b.(string))
	default:
		panic("can not compare unkown type value by priority queue")
	}
}

func (p *pqueue) ToString() {
	fmt.Println(p.queue)
}

func NewPriorityQueue(capacity int) Loopqueue {
	return &pqueue{
		nil,
		make([]interface{}, 0, capacity),
		0,
	}
}
