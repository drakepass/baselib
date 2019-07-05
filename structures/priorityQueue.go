package structures

import (
	"errors"
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
	if right < p.size && p.queue[left].(int) < p.queue[right].(int) {
		if p.queue[k].(int) < p.queue[right].(int) {
			p.queue[k], p.queue[right] = p.queue[right], p.queue[k]
			p.shiftdown(right)
		}
	} else {
		if p.queue[k].(int) < p.queue[left].(int) {
			p.queue[k], p.queue[left] = p.queue[left], p.queue[k]
			p.shiftdown(left)
		}
	}
}
func (p *pqueue) GetFront() (interface{}, error) {
	if p.size == 0 {
		err := errors.New("there is no element in the priority queue")
		return nil, err
	}
	return p.queue[0], nil
}
func (p *pqueue) shiftup(k int) {
	if k == 0 {
		return
	}
	fk := k / 2
	if (p.queue[fk]).(int) < (p.queue[k]).(int) {
		p.queue[fk], p.queue[k] = p.queue[k], p.queue[fk]
		p.shiftup(fk)
	}
}

//func (p *pqueue) ToString() {
//	fmt.Println(p.queue)
//}
func NewPriorityQueue(capacity int) *pqueue {
	return &pqueue{
		nil,
		make([]interface{}, 0, capacity),
		0,
	}
}
