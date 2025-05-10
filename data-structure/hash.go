package datastructure

import "sync"

// the size of a map should never be zero because hashing uses modulo
// we have 3 options
// 1. let user define the size of slice
// 2. have a constructor func for initialization e.g NewHashMap(size int) ans if size is not provided default to a size
// 3. use go sync.Once
type MapNode struct {
	Key   string
	Value interface{}
}

type SetNode struct {
	Value interface{}
	Next  *SetNode
}

type HashMap struct {
	Bucket []*MapNode
	Slots  int // number of spaces
	Size   int // number of filed spaces
	once   sync.Once
}

type HashSet struct {
	Hash []*SetNode
	Size int
}

type IHash interface {
	initHashMap(...int)
	hashKey(key interface{}) int
	// initHashSet()
	Set()
	Get()
	Remove()
	Len()
	Keys()
	Values()
}

func NewHashMap(slots int) *HashMap {
	var defaultSlots = 4
	if slots > 0 {
		defaultSlots = slots
	}
	m := &HashMap{
		Bucket: make([]*MapNode, defaultSlots),
		Slots:  defaultSlots,
	}
	return m
}

func (m *HashMap) initHashMap(slots ...int) {
	m.once.Do(func() {
		if len(slots) > 0 {
			m.Slots = slots[0]
		} else {
			m.Slots = 4
		}
		m.Bucket = make([]*MapNode, m.Slots)

	})
}

func (m *HashMap) hashKey(key interface{}) int {
	switch key := key.(type) {
	case string:
		asciiTotal := 0
		for _, char := range key {
			asciiTotal += int(char)
		}
		return asciiTotal % m.Slots
	}
	return 0
}

func (m *HashMap) checkCollision(key string, pos int) bool {
	return m.Bucket[pos] != nil && m.Bucket[pos].Key != key // setting another key at an index that is taken by another
}

func (m *HashMap) resize() {
	// resize and rehash
	oldSlice := m.Bucket
	slots := m.Slots * 2
	m.Bucket = make([]*MapNode, m.Slots, slots)
	m.Slots = slots
	m.Size = 0
	for _, h := range oldSlice {
		m.Set(h.Key, h.Value)
	}
}

func (m *HashMap) Set(key string, value interface{}) {
	m.initHashMap()
	pos := m.hashKey(key)

	// check for collision
	if m.checkCollision(key, pos) {
		m.resize()
	}

	m.Bucket[pos] = &MapNode{
		Key:   key,
		Value: value,
	}
	m.Size++
}

func (m *HashMap) Get(key string) interface{} {
	pos := m.hashKey(key)
	item := m.Bucket[pos]
	if item != nil && item.Key == key {
		return item.Value
	}
	return nil
}

func (m *HashMap) Remove(key string) {
	pos := m.hashKey(key)
	item := m.Bucket[pos]
	if item != nil && item.Key == key {
		m.Bucket = append(m.Bucket[:pos], m.Bucket[pos+1:]...)
		m.Size--
	}
}

func (m *HashMap) Len(key string) int {
	return m.Size
}

func (m *HashMap) Keys() []string {
	var keys []string
	for _, item := range m.Bucket {
		if item != nil {
			keys = append(keys, item.Key)
		}
	}
	return keys
}

func (m *HashMap) Values() []interface{} {
	var values []interface{}
	for _, item := range m.Bucket {
		if item != nil {
			values = append(values, item.Value)
		}
	}
	return values
}