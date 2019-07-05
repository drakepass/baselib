package structures

type BstSet interface {
	Add(int)
	Remove(int)
	IsEmpty() bool
	GetSize() int
	Contains(int) bool
}

type set struct {
	b Bst
}

func (s *set) Add(e int) {
	s.b.Add(e)
}

func (s *set) Remove(e int) {
	s.b.Delete(e)
}

func (s *set) IsEmpty() bool {
	return s.b.IsEmpty()
}

func (s *set) Contains(e int) bool {
	return s.b.Find(e)
}

func (s *set) GetSize() int {
	return s.b.GetSize()
}

func NewBstSet() *set {
	var b Bst
	b = NewBst()
	return &set{
		b,
	}
}
