package hashtable

import (
	"fmt"
)

const (
	KEY_SIZE   int = 256
	VALUE_SIZE int = 1048576
)

/** >> HashTable **/

type HashTable struct {
	size    int
	buckets []*LinkedList
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*LinkedList, size),
	}
}

// Hash function: simple modulo hashing
func (ht *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) % ht.size
	}
	return hash
}

// Insert a key-value pair into the hash table
func (ht *HashTable) Insert(key string, value string) {
	kbytes, vbytes := []byte(key), []byte(value)
	if len(kbytes) >= KEY_SIZE {
		panic(fmt.Sprintf("Key length must be of %d bytes", KEY_SIZE))
	}
	if len(vbytes) >= VALUE_SIZE {
		panic(fmt.Sprintf("Value length must be of %d bytes", VALUE_SIZE))
	}
	index := ht.hash(key)
	if ht.buckets[index] == nil {
		ht.buckets[index] = CreateLinkedList()
	}
	ht.buckets[index].insert(kbytes, vbytes)
}

// Get the value associated with a key
func (ht *HashTable) Get(key string) (string, bool) {
	index := ht.hash(key)
	if ht.buckets[index] == nil {
		return "", false
	}
	v, e := ht.buckets[index].search([]byte(key))
	if e {
		return "", e
	}
	return string(v), false
}

func (ht *HashTable) Update(key string, value string) bool {
	index := ht.hash(key)
	if ht.buckets[index] == nil {
		return false
	}
	k, v := []byte(key), []byte(value)
	ht.buckets[index].update(k, v)
	return true
}

// Delete a key-value pair from the hash table
func (ht *HashTable) Delete(key string) {
	index := ht.hash(key)
	if ht.buckets[index] != nil {
		ht.buckets[index].delete([]byte(key))
	}
}

/** << End HashTable **/
