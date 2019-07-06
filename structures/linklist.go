package structures

type Linklist interface {
	GetSize() int
	GetFirst() interface{}
	Add(index int, e interface{})
	Delete(index int)
	Update(index int, e interface{}) error
	GetLast() interface{}
	GetAt(index int) interface{}
	AddFirst(e interface{}) // this function can be used in queue structure
	AddLast(e interface{})  // this function can be used in queue and stack structure
	IsEmpty() bool
	DeleteAll(e interface{})
}

type list struct {
	e    interface{}
	next *list
}

type linklist struct {
	size    int
	element *list
}

func (l *linklist) GetSize() int {
	return l.size
}

func (l *linklist) GetAt(index int) interface{} {
	if 0 <= index && index < l.size {
		cur := *l.element
		for i := 0; i < index+1; i++ {
			cur = *cur.next
		}
		return cur.e
	}
	panic("the index is out of range")
}

func (l *linklist) GetFirst() interface{} {
	return l.GetAt(0)
}
func (l *linklist) GetLast() interface{} {
	return l.GetAt(l.size)
}

func (l *linklist) Add(index int, e interface{}) {
	cur := l.element
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.next = &list{
		e,
		cur.next,
	}

	l.size++
	return
}

func (l *linklist) Delete(index int) {
	if 0 <= index && index < l.size {
		cur := l.element
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		cur.next = (cur.next).next
		l.size--
		return
	}
	panic("the index is out of range")
}

func (l *linklist) DeleteAll(e interface{}) {
	//deleteElement(l,l.element,e)  采用递归方式删除
	element := l.element
	for {
		if element.next == nil {
			break
		}
		if element.next.e == e {
			element.next = element.next.next
			l.size--
		} else {
			element = element.next
		}
	}
	return
}

func (l *linklist) Update(index int, e interface{}) error {
	if 0 <= index && index < l.size {
		cur := l.element
		for i := 0; i < index+1; i++ {
			cur = cur.next
		}
		cur.e = e
		return nil
	}
	panic("the index is out of range")
}

func (l *linklist) AddFirst(e interface{}) {
	l.Add(0, e)
	return

}

func (l *linklist) AddLast(e interface{}) {
	l.Add(l.size, e)
	return
}

func (l *linklist) IsEmpty() bool {
	return l.size == 0
}

func deleteElement(L *linklist, l *list, e interface{}) *list {
	if l.next == nil {
		return nil
	}
	l.next.next = deleteElement(L, l.next, e)
	if l.next.e == e {
		L.size--
		return l.next.next
	} else {
		return l.next
	}

}

func NewLinklist() Linklist {
	element := list{
		nil,
		nil,
	}
	return &linklist{
		0,
		&element,
	}
}
