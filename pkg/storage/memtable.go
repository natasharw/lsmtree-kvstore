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
	kv       map[string][]byte
	approxSz int
	max      int
}

/*
WRITE OPERATIONS
*/

// MemtableInit : creates a memtable in local memory
func MemtableInit() (m *Memtable) {
	log.Printf("Creating a memtable (C0)")
	memAddress := &Memtable{
		kv:       make(map[string][]byte),
		approxSz: 0,
		max:      MaxMem}
	log.Printf("Key-value store initialised with a memtable")

	return memAddress

}

// InsertToMemtable : inserts a key and value to memtable
func (mt *Memtable) InsertToMemtable(key string, value []byte) (int error) {

	log.Printf("Handling setting key \"%s\" in memtable", key)

	_, present, err := mt.SearchMemtable(key)

	if !present {
		err := mt.insertKey(key)
		mt.approxSz++

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
func (mt *Memtable) insertKey(key string) error {
	log.Printf("Finding correct place for key \"%s\" in memtable", key)

	// [TODO] - binary search to insert key in correct place

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

	if mt.approxSz < mt.max {
		log.Printf("Memtable not full. Approx size: %d Max: %d", mt.approxSz, mt.max)

		return false
	}

	log.Printf("Memtable is full. Approx size: %d Max: %d", mt.approxSz, mt.max)

	return true
}

/*
READ OPERATIONS
*/

// SearchMemtable : given a key, tries to find an entry for it in memtable
func (mt *Memtable) SearchMemtable(key string) ([]byte, bool, error) {
	log.Printf("Searching memtable for key \"%s\"", key)

	value, present := mt.kv[key]

	if !present {
		log.Printf("Key %s not found in memtable", key)
		return nil, false, nil
	}

	log.Printf("Found value \"%s\" (\"%v\") for key \"%s\"", value, value, key)
	return value, true, nil

}
