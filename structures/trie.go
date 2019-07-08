package structures

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

type Trie interface {
	Set(interface{}, interface{})
	Get(interface{}) interface{}
	Delete(interface{})
	ToString()
	GetSize() int
}
type tree struct {
	value  interface{}
	c      rune
	isWord bool
	next   map[rune]*tree
}
type tenode struct {
	root *tree
	size int
}

func (t *tenode) GetSize() int {
	return t.size
}

func getString(str interface{}) (tmp string) {
	tp := reflect.TypeOf(str)
	if tp.Kind() == reflect.String {
		tmp = str.(string)
	} else {
		tmpint := str.(int)
		tmp = strconv.Itoa(tmpint)
	}
	return
}
func (t *tenode) Set(index, value interface{}) {
	indexstr := getString(index)
	t.toSet(t.root.next, indexstr, value)
}
func (t *tenode) toSet(Next map[rune]*tree, index string, value interface{}) {
	i := 0
	for _, v := range index {
		if _, ok := Next[v]; !ok {
			isWord := false
			var vl interface{}
			if i == utf8.RuneCountInString(index)-1 {
				isWord = true
				vl = value
				t.size++
			}
			Next[v] = &tree{
				vl,
				v,
				isWord,
				make(map[rune]*tree),
			}
		} else if i == utf8.RuneCountInString(index)-1 {
			Next[v].value = value
			Next[v].isWord = true
			return
		}
		Next = Next[v].next
		i++
	}
	return
}
func (t *tenode) ToString() {
	fmt.Println(t.root.next)
	fmt.Println(t.root.next['a'].next['b'])
	fmt.Println(t.root.next['a'].next['b'].next['c'])
}
func (t *tenode) Get(index interface{}) interface{} {
	indexstr := getString(index)
	Next := t.root.next
	i := 0
	for _, v := range indexstr {
		if _, ok := Next[v]; !ok {
			return false
		} else if i == utf8.RuneCountInString(indexstr)-1 {
			if Next[v].isWord == true {
				return Next[v].value
			}
			return false
		} else {
			Next = Next[v].next
		}
		i++
	}
	return false
}
func (t *tenode) Delete(index interface{}) {
	indexstr := getString(index)
	t.dodelete(t.root.next, indexstr, utf8.RuneCountInString(indexstr)-1, 1)
}

func (t *tenode) dodelete(Next map[rune]*tree, indexstr string, until, isd int) {
	if until < 0 {
		return
	}
	i := 0
	for _, v := range indexstr {
		if _, ok := Next[v]; ok {
			if i == until {
				if isd == 1 {
					t.size--
					if len(Next[v].next) == 0 {
						delete(Next, v)
					} else {
						Next[v].isWord = false
						return
					}

				} else if Next[v].isWord == false {
					if len(Next[v].next) == 0 {
						delete(Next, v)
					}
				}
				break
			}
			Next = Next[v].next
		} else {
			return
		}
		i++
	}
	t.dodelete(t.root.next, indexstr, until-1, 0)
}
func NewTrie() Trie {
	tree := &tree{
		nil,
		0,
		false,
		make(map[rune]*tree),
	}
	return &tenode{
		tree,
		0,
	}
}
