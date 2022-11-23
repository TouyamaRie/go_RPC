package xclient

import (
	"hash/crc32"
	"log"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constains all hashed keys
type Map struct {
	hash     Hash           //哈希算法
	replicas int            //虚拟节点倍数
	keys     []int          // Sorted//哈希环
	hashMap  map[int]string //虚拟节点和真实节点的映射
}

// New creates a Map instance
func NewHash(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys) //排序
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	log.Println(idx)
	//因为有可能找到大于最大值的节点，此时会返回len(m.keys)，求余之后就会映射到第0个节点
	log.Println(m.hashMap[m.keys[idx%len(m.keys)]])
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
