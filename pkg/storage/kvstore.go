package storage

import (
	"log"
)

const (
	// fanout : ratio between levels of LSM-tree e.g. C2 = C1 * fanout
	fanout int = 2
	// initL : inital levels of LSM-tree. 1 = Memtable (C0) only, 2 = Memtable (C0) + C1
	initL int = 1
)

type storage interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

// KVStore : represents the entirety of a key-value store
type KVStore struct {
	buffer *Memtable
	levels int
	fanout int
}

// NewKVStore : instantiates a new key-value store with some default values
func NewKVStore() *KVStore {
	log.Printf("Instantiating a new key-value store")
	log.Printf("Setting attributes: Memtable as buffer, initial levels as %d and level growth factor of %d", levels, fanout)

	mt := new(Memtable)
	kvs := &KVStore{
		buffer: mt,
		levels: initL,
		fanout: fanout}

	log.Printf("Key-value store instantiated")
	return kvs
}

/*
WRITE OPERATIONS
*/

// Set : Main function to set a key in the key-value store
func (kvs *KVStore) Set(key string, value []byte) error {
	log.Printf("Setting key %s", key)
	mt := kvs.buffer

	mt.InsertToMemtable(key, value)

	switch f := mt.IsMemtableFull(); f {
	case true:
		err := kvs.Flush()
		log.Printf("%s: %s added to memtable and flushed to disk", key, value)

		if err != nil {
			return err
		}

	case false:
		log.Printf("%s: %s added to memtable", key, value)
	}

	return nil
}

// Flush : flushes memtable to disk as SSTable
func (kvs *KVStore) Flush() error {
	// TODO : create a new memtable while the last one is being written?
	ss := kvs.buffer.NewSSTable()
	ss.WriteSSTableToDisk()

	return nil
}

/*
READ OPERATIONS
*/

// Get : main function to get a key's value from the key-value store
func (kvs *KVStore) Get(key string) ([]byte, error) {
	var results []byte

	results, p, err := kvs.buffer.SearchMemtable(key)

	if err != nil {
		return nil, err
	}

	if !p {
		results, err = kvs.searchDisk(key)

		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

// searchDisk : performs search on-disk for a given key
func (kvs *KVStore) searchDisk(key string) ([]byte, error) {

	indices, err := kvs.loadIndices()
	if err != nil {
		panic(err)
	}

	ir, err := kvs.searchIndices(indices)

	result := kvs.getFromDisk(ir, key)

	return result, nil
}

// loadIndices : loads indices of SStables into local memory
func (kvs *KVStore) loadIndices() (map[string][]byte, error) {
	// TODO ; for each SSTable, load its block index (key-off pairs) into a structure in local memory
	// TODO ; load all -> search one by one or load one -> search one
	indices := make(map[string][]byte)

	return indices, nil
}

// searchIndices : searches through key-offset pairs of all sstables in local memory
func (kvs *KVStore) searchIndices(map[string][]byte) (map[string]map[string][]byte, error) {
	// TO DO - search through ss-tables from most recent to least recent
	var r map[string]map[string][]byte

	result := r

	return result, nil
}

// getFromDisk : returns a value for a key given a ssblock identifier, and key's approx location
func (kvs *KVStore) getFromDisk(ir map[string]map[string][]byte, key string) []byte {
	var results []byte
	for _, r := range ir {
		for tbl := range r {
			print(tbl)
			ss := new(SSTable) // TODO dummy overwrite -> change this to the search indices result
			results = ss.GetValue(key)
		}
	}

	return results
}

/*
OTHER OPERATIONS
*/

// compaction : reorganises SSTables between levels of the tree to store efficiently
func (kvs *KVStore) compaction() error {
	return nil
	// TODO implement a mergesort between sorted arrays

	// if l1 is None: return l2
	// if l2 is None: return l1
	// if l1.next is None and l2.next is None:
	// 	if l1.val > l2.val:
	// 		return ListNode(l2.val, next=l1)
	// 	else:
	// 		return ListNode(l1.val, next=l2)
	// elif l1.val < l2.val:
	// 	return ListNode(l1.val,self.mergeTwoLists(l1.next,l2))
	// elif l2.val < l1.val:
	// 	return ListNode(l2.val,self.mergeTwoLists(l1,l2.next))
	// elif l1.val == l2.val:
	// 	return ListNode(l1.val,self.mergeTwoLists(l1.next,l2))
	// return nil
}
