package dataStructuresAndAlgorithms

import (
	"fmt"
	"sync"
)

/**
哈希表是一种使用哈希函数组织数据，以支持快速插入和搜索的数据结构。
有两种不同类型的哈希表：哈希集合和哈希映射。
哈希集合是集合数据结构的实现之一，用于存储非重复值。
哈希映射是映射 数据结构的实现之一，用于存储(key, value)键值对。
*/

/**
1. 哈希函数 一元函数 y = F(x)
2. 冲突解决  F(x1) = F(x2) (x1 != x2) 时如何解决?
3. 哈希表的构建 = 哈希函数 + 冲突解决
4. 功能上应具备 add/put、contains、remove
*/

//HashSet hash集合
type HashSet struct {
	set   []linkList //采用链表实现单个桶 解决hash冲突
	size  int
	count int
	sync.RWMutex
}

func NewHashSet() *HashSet {
	return &HashSet{
		set:   make([]linkList, 16),
		size:  16,
		count: 0,
	}
}

func (h *HashSet) Add(value int) {
	h.RWMutex.Lock()
	defer h.RWMutex.Unlock()
	if float64(h.count)/float64(h.size) > 0.75 {
		h.size = h.size << 1
		h.capacityExpansion()
	}
	index := h.hash(value)
	if h.set[index].add(value) {
		h.count++
	}
}

//hash hash运算 数学公式 y = x % size,简单的一元一次函数运用，作为哈希函数
func (h *HashSet) hash(x int) int {
	return x % h.size
}

//capacityExpansion 扩容
func (h *HashSet) capacityExpansion() {
	newLinkList := make([]linkList, h.size)
	for _, list := range h.set {
		for cur := list.node; cur != nil; cur = cur.next {
			index := h.hash(cur.value)
			newLinkList[index].add(cur.value)
		}
	}
	h.set = newLinkList
}

func (h *HashSet) Contains(value int) bool {
	h.RWMutex.RLock()
	defer h.RWMutex.RUnlock()
	index := h.hash(value)
	return h.set[index].contains(value)
}

func (h *HashSet) Remove(value int) {
	h.RWMutex.Lock()
	defer h.RWMutex.Unlock()
	index := h.hash(value)
	h.set[index].remove(value)
	h.count--
}

//HashMap hash映射
type HashMap struct {
}

func (h *HashMap) Put(key string, value int) {
	panic("implement me")
}

func (h *HashMap) Get(key string) int {
	panic("implement me")
}

func (h *HashMap) Remove(value int) {
	panic("implement me")
}

//linkList 链表
type linkList struct {
	node *node
}

//node 链表节点
type node struct {
	value int
	next  *node
	last  *node
}

func (l *linkList) add(value int) bool {
	if l.node == nil {
		l.node = &node{
			value: value,
			last:  nil,
			next:  nil,
		}
	} else {
		for cur := l.node; cur != nil; cur = cur.next {
			if cur.value == value {
				return false
			} else if cur.next == nil {
				cur.next = &node{
					value: value,
					last:  cur,
					next:  nil,
				}
				return true
			}
		}
	}
	return true
}

func (l *linkList) contains(value int) bool {
	for cur := l.node; cur != nil; cur = cur.next {
		if cur.value == value {
			return true
		}
	}
	return false
}

func (l *linkList) remove(value int) {
	for cur := l.node; cur != nil; cur = cur.next {
		if cur.value == value {
			if cur.last != nil {
				cur.last = cur.next
			} else {
				l.node = cur.next
			}
			return
		}
	}
}

func (l *linkList) read() {
	values := []int{}
	for cur := l.node; cur != nil; cur = cur.next {
		values = append(values, cur.value)
	}
	fmt.Printf("linkList values is %v \r\n", values)
}
