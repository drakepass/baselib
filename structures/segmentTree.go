package structures

import "fmt"

type SegmentTree interface {
	Query(int, int) interface{}
	Update(int, interface{})
	ToString()
}
type segmentArray struct {
	arr  []interface{}
	tree []interface{}
}

func (s *segmentArray) Query(l, r int) interface{} {
	return s.toQuery(0, 0, len(s.arr)-1, l, r)
}
func (s *segmentArray) toQuery(index, left, right, l, r int) interface{} {
	if left == l && right == r {
		return s.tree[index]
	}
	mid := left + (right-left)/2
	nleft := index*2 + 1
	nright := index*2 + 2
	if r <= mid {
		return s.toQuery(nleft, left, mid, l, r)
	}
	if l > mid {
		return s.toQuery(nright, mid+1, right, l, r)
	}
	a := s.toQuery(nleft, left, mid, l, mid)
	b := s.toQuery(nright, mid+1, right, mid+1, r)
	return a.(int) + b.(int)
}

func (s *segmentArray) Update(index int, e interface{}) {
	s.toUpdate(0, 0, len(s.arr)-1, index, e)
}
func (s *segmentArray) toUpdate(index, left, right, key int, e interface{}) {
	if left == right {
		s.arr[key] = e
		s.tree[index] = e
		return
	}
	mid := left + (right-left)/2
	leftindex := 2*index + 1
	rightindex := 2*index + 2
	if key <= mid {
		s.toUpdate(leftindex, left, mid, key, e)
	} else {
		s.toUpdate(rightindex, mid+1, right, key, e)
	}
	s.arr[key] = e
	s.tree[index] = s.tree[leftindex].(int) + s.tree[rightindex].(int)
	return
}
func (s *segmentArray) ToString() {
	fmt.Println(s.arr)
	fmt.Println("------------")
	fmt.Println(s.tree)
}
func (s *segmentArray) buildSegmentTree(index, l, r int) {
	if l == r {
		s.tree[index] = s.arr[l]
		return
	}
	mid := l + (r-l)/2
	left := 2*index + 1
	right := 2*index + 2
	s.buildSegmentTree(left, l, mid)
	s.buildSegmentTree(right, mid+1, r)
	s.tree[index] = s.tree[left].(int) + s.tree[right].(int)
	return
}
func NewSegmentTree(a ...interface{}) SegmentTree {
	size := len(a) * 4
	var segmentTree SegmentTree
	var tree = make([]interface{}, size, size)
	s := &segmentArray{
		a,
		tree,
	}
	s.buildSegmentTree(0, 0, len(a)-1)
	segmentTree = s
	return segmentTree
}
