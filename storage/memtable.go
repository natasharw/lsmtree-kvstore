package storage

import (
	"errors"
	"log"
)

// Memtable : in-memory mutable data structure: sorted map of keys and values
type Memtable struct {
	m   map[string][]byte
	max int
}

const (
	// MaxMemEntries : sets the maximum amount of key-value pairs before memtable flushes to disk
	MaxMemEntries int = 1000
)

/*
WRITE OPERATIONS
*/

// MemtableInit : creates a memtable in local memory
func MemtableInit() (m *Memtable) {
	log.Printf("Creating a memtable (C0)")
	memAddress := &Memtable{
		m:   make(map[string][]byte),
		max: MaxMemEntries}
	log.Printf("Key-value store initialised with a memtable")
	return memAddress
}

// InsertToMemtable : inserts a key and value to memtable
func InsertToMemtable(mt *Memtable, key string, value []byte) (int error) {

	log.Printf("Handling setting key \"%s\" in memtable", key)

	check, err := SearchMemtable(mt, key)

	if len(check) == 0 {
		log.Printf("Key \"%s\" does not exist in memtable already", key)
		_, err := InsertKeyInOrder(mt, key)
		if err != nil {
			return errors.New("Failed to insert key in memtable")
		}
	}

	_, err = SetValueOnKey(mt, key, value)

	if err != nil {
		return errors.New("Failed to set value against key in memtable")
	}
	return nil
}

// InsertKeyInOrder : inserts key in correct place in memtable using binary search
// [TODO] - binary search to insert key in correct place
func InsertKeyInOrder(mt *Memtable, key string) (int, error) {
	log.Printf("Finding correct place for key \"%s\" in memtable", key)
	return 0, nil
}

// SetValueOnKey : sets a value against a key in a memtable
func SetValueOnKey(mt *Memtable, key string, value []byte) (int, error) {
	log.Printf("Setting value of key \"%s\" in memtable", key)
	mt.m[key] = value
	log.Printf("Successfully set key \"%s\" to %v in memtable", key, value)
	return 0, nil
}

// IsMemtableFull : returns true if memtable entries equals max constant
func IsMemtableFull(mt *Memtable) bool {
	log.Printf("Checking if memtable is full")
	mtEntries := len(mt.m)
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
func SearchMemtable(mt *Memtable, key string) ([]byte, error) {
	log.Printf("Searching memtable for key \"%s\"", key)

	//[TODO] func (m *Memtable) func SearchMemtable(key []byte, m *Memtable) ([]byte, error) ???

	value, ok := mt.m[key]
	if !ok {
		return nil, errors.New("Error finding key in memtable")
	}

	log.Printf("Found value \"%s\" for key \"%s\"", value, key)
	return value, nil
}
