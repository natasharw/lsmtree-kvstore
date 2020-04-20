package storage

import (
	"errors"
	"log"
)

const (
	// MaxMem : sets the maximum amount of key-value pairs before memtable flushes to disk
	MaxMem int = 1000
)

// Memtable : in-memory mutable data structure: here a sorted map of key-value pairs
type Memtable struct {
	kv  map[string][]byte
	max int
}

/*
WRITE OPERATIONS
*/

// MemtableInit : creates a memtable in local memory
func MemtableInit() (m *Memtable) {
	log.Printf("Creating a memtable (C0)")
	memAddress := &Memtable{
		kv:  make(map[string][]byte),
		max: MaxMem}
	log.Printf("Key-value store initialised with a memtable")

	return memAddress

}

// InsertToMemtable : inserts a key and value to memtable
func (mt *Memtable) InsertToMemtable(key string, value []byte) (int error) {

	log.Printf("Handling setting key \"%s\" in memtable", key)

	value, p, err := mt.SearchMemtable(key)

	if !p {
		log.Printf("Key \"%s\" does not exist in memtable already", key)
		err := mt.insertKey(key)
		if err != nil {
			return errors.New("Failed to insert key in memtable")
		}
	}

	err = mt.setValueOnKey(key, value)

	if err != nil {
		return errors.New("Failed to set value against key in memtable")
	}

	return nil
}

// insertKey : inserts key in correct place in memtable using binary search
// [TODO] - binary search to insert key in correct place
func (mt *Memtable) insertKey(key string) error {
	log.Printf("Finding correct place for key \"%s\" in memtable", key)

	return nil
}

// setValueOnKey : sets a value against a key in a memtable
func (mt *Memtable) setValueOnKey(key string, value []byte) error {
	log.Printf("Setting value of key \"%s\" in memtable", key)
	mt.kv[key] = value
	log.Printf("Successfully set key \"%s\" to %v in memtable", key, value)

	return nil
}

// IsMemtableFull : returns true if memtable entries equals max constant
func (mt *Memtable) IsMemtableFull() bool {
	log.Printf("Checking if memtable is full")
	mtEntries := len(mt.kv)
	if mtEntries < mt.max {
		log.Printf("Memtable not full. Status: %d out of max %d", mtEntries, mt.max)

		return false
	}

	log.Printf("Memtable is full. Status: %d out of max %d", mtEntries, mt.max)

	return true
}

/*
READ OPERATIONS
*/

// SearchMemtable : given a key, tries to find an entry for it in memtable
func (mt *Memtable) SearchMemtable(key string) ([]byte, bool, error) {
	log.Printf("Searching memtable for key \"%s\"", key)

	value, p := mt.kv[key]

	if !p {
		log.Printf("%s key not found in memtable", key)
		return nil, false, nil
	}

	log.Printf("Found value \"%v\" for key \"%s\"", value, key)
	return value, true, nil

}
