package storage

import (
	"fmt"
)

type Memtable struct {
	m map[string][]byte
}

/*
WRITE OPERATIONS
*/

func MemtableInit() (m *Memtable) {
	return &Memtable{m: make(map[string][]byte)}
}

func SetInMemtable(key []byte, value []byte) error {
	// [TODO] : get the next line working
	// m.m[key] := value
	memtableentry := value
	fmt.Printf("%v", memtableentry) // [TODO] placeholder only - remove

	return nil
}

func IsMemtableFull(m *Memtable) bool {
	// [TO DO]
	m_entries := 1
	if m_entries < MAX_MEMTABLE_ENTRIES {
		return false
	} else {
		return true
	}
}

/*
READ OPERATIONS
*/

func SearchMemtable(key []byte, m *Memtable) ([]byte, error) {
	//[TODO] func (m *Memtable) func SearchMemtable(key []byte, m *Memtable) ([]byte, error) {??
	value := key

	// [TODO] - fix this so it works (using a byte to search for a key in m[key])
	//value, ok := m.m[key] - [TODO] GET THIS BACK
	// if !ok {
	// 	return nil, errors.New("key not found in memtable")
	// }
	return value, nil
}
