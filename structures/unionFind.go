package structures

type UnionFind interface {
	Find(int) int
	Union(int, int)
	isConnected(int, int) bool
}
type unionNode struct {
	fid  int
	rank int
}
type unionTree struct {
	ele  map[int]*unionNode
	size int
}

func (u *unionTree) Find(q int) int {
	if q < 0 || q > u.size {
		panic("the param int is out of bound!")
	}
	for q != u.ele[q].fid {
		u.ele[q].fid = u.ele[u.ele[q].fid].fid //压缩节点q的层高,减少下次查询复杂度
		q = u.ele[q].fid

	}
	return q
}
func (u *unionTree) Union(p, q int) {
	pid := u.Find(p)
	qid := u.Find(q)
	if pid == qid {
		return
	}
	if u.ele[pid].rank < u.ele[qid].rank {
		u.ele[p].fid = u.ele[q].fid
	} else if u.ele[pid].rank < u.ele[qid].rank {
		u.ele[q].fid = u.ele[p].fid
	} else {
		u.ele[q].fid = u.ele[p].fid
		u.ele[pid].rank++
	}
}
func (u *unionTree) isConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
func NewUnionFind(s ...int) UnionFind {
	ele := make(map[int]*unionNode)

	for _, v := range s {
		ele[v] = &unionNode{
			v,
			1,
		}
	}
	size := len(ele)
	return &unionTree{
		ele,
		size,
	}
}
