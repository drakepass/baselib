package structures

import "fmt"

type SegmentTree interface {
	Query(int, int) interface{}
	Update(int)
	ToString()
}
type segmentArray struct {
	arr  []interface{}
	tree []interface{}
}

func (s *segmentArray) Query(l, r int) interface{} {

}

func (s *segmentArray) Update(index int) {
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
