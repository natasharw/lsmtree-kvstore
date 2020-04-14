package storage

import "errors"

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
	return &Memtable{
		m:   make(map[string][]byte),
		max: MaxMemEntries}
}

// InsertToMemtable : inserts a key and value to memtable
func InsertToMemtable(key string, value []byte) (int error) {
	// if !memtable {
	// 	memtable = MemtableInit()
	// }

	// [TODO} "if key does not exist in the memtable then":
	_, err := InsertKeyInOrder(key)
	// "else do nothing"

	if err != nil {
		return errors.New("failed to set key in memtable")
	}

	_, err = SetValueOnKey(key, value)

	if err != nil {
		return errors.New("failed to set pair in memtable")
	}
	return nil
}

// InsertKeyInOrder : inserts key in correct place in memtable using binary search
func InsertKeyInOrder(key string) (int, error) {
	return 0, nil
}

// SetValueOnKey : sets a value against a key in a memtable
func SetValueOnKey(key string, int []byte) (int, error) {
	return 0, nil
}

// IsMemtableFull : returns true if memtable entries equals max constant
func IsMemtableFull(m *Memtable) bool {
	// [TO DO]
	mEntries := 1
	if mEntries < MaxMemEntries {
		return false
	}

	return true
}

/*
READ OPERATIONS
*/

// SearchMemtable : given a key, tries to find an entry for it in memtable
func SearchMemtable(key string, m *Memtable) ([]byte, error) {
	//[TODO] func (m *Memtable) func SearchMemtable(key []byte, m *Memtable) ([]byte, error) {??
	value := []byte{1, 2, 3}

	// [TODO] - fix this so it works (using a byte to search for a key in m[key])
	//value, ok := m.m[key] - [TODO] GET THIS BACK
	// if !ok {
	// 	return nil, errors.New("key not found in memtable")
	// }
	return value, nil
}
