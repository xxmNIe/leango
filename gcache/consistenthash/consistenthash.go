package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	//key 的hash
	keys    []int
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			has := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, has)
			m.hashMap[has] = key
		}
	}
	sort.Ints(m.keys)
	//fmt.Println(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	has := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= has
	})
	//fmt.Println("hash is :", has, "  idx is :", idx)
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
