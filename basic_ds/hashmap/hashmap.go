package hashmap

import (
	"math"
	"sync"
)

var hashAlgorithm = func(key []byte) uint64 {
	h := New64()
	h.Write(key)
	return h.Sum64()
}

const (
	expandFactor = 0.75
)

type keyPairs struct {
	key   string
	value interface{}
	next  *keyPairs
}

type HashMap struct {
	array   []*keyPairs
	cap     int
	len     int
	capMask int
	lock    sync.Mutex
}

func NewHashMap(cap int) *HashMap {
	defaultCap := 1 << 4
	if cap <= defaultCap {
		cap = defaultCap
	} else {
		cap = 1 << (int(math.Ceil(math.Log2(float64(cap)))))
	}

	m := new(HashMap)
	m.array = make([]*keyPairs, cap, cap)
	m.cap = cap
	m.capMask = cap - 1

	return m
}

func (m *HashMap) Capiticy() int {
	return m.cap
}

func (m *HashMap) Len() int {
	return m.len
}

func (m *HashMap) hashIndex(key string, mask int) int {
	hash := hashAlgorithm([]byte(key))

	index := hash & uint64(mask)

	return int(index)
}

func (m *HashMap) Put(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capMask)

	element := m.array[index]

	if element == nil {
		m.array[index] = &keyPairs{
			key:   key,
			value: value,
		}
	} else {
		var lastPairs *keyPairs

		for element != nil {
			if element.key == key {
				element.value = value
				return
			}
			lastPairs = element
			element = element.next
		}

		lastPairs.next = &keyPairs{
			key:   key,
			value: value,
		}
	}

	newLen := m.len + 1

	if float64(newLen)/float64(m.cap) >= expandFactor {
		newM := new(HashMap)

		newM.array = make([]*keyPairs, 2*m.cap, 2*m.cap)
		newM.cap = 2 * m.cap
		newM.capMask = 2*m.cap - 1

		for _, pairs := range m.array {
			for pairs != nil {
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.next
			}
		}

		m.array = newM.array
		m.cap = newM.cap
		m.capMask = newM.capMask

	}
	m.len = newLen
}

func (m *HashMap) Get(key string) (value interface{}, ok bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capMask)

	element := m.array[index]

	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.next
	}
	return nil, false
}

func (m *HashMap) Delete(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	index := m.hashIndex(key, m.capMask)

	element := m.array[index]

	if element == nil {
		return
	}

	if element.key == key {
		m.array[index] = element.next
		m.len = m.len - 1
		return
	}

	nextElement := element.next
	for nextElement != nil {
		if nextElement.key == key {
			element.next = nextElement.next
			m.len = m.len - 1
			return
		}
		element = nextElement
		nextElement = nextElement.next
	}
}

func (m *HashMap) Range(visit func(key string, value interface{})) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for _, pairs := range m.array {
		for pairs != nil {
			visit(pairs.key, pairs.value)
			pairs = pairs.next
		}
	}

}
