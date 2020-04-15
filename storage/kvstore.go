package storage

import (
	"errors"
	"fmt"
)

const (
	// LvlGrowthFactor : magnification of data to lower level of LSM-tree
	// E.g. C1 = C0 * LvlGrowthFactor
	LvlGrowthFactor int = 2
)

type storage interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
}

// Store : represents the entirety of a key-value store
type Store struct {
	buffer *Memtable
	levels []byte
	gf     []byte
}

/*
WRITE OPERATIONS
*/

// Put : Main function to set a key in the key-value store
// Returns 0 if sucessful
func Put(mt *Memtable, key string, value []byte) (int, error) {

	InsertToMemtable(mt, key, value)

	switch f := IsMemtableFull(mt); f {
	case true:
		Flush(mt)
		return fmt.Printf("%s: %s added to buffer and buffer written to disk", key, value)
	case false:
		return fmt.Printf("%s: %s added to buffer", key, value)
	}

	return 0, nil
}

// Flush : takes the current memtable and flushes to disk
func Flush(mt *Memtable) int {
	s := SSTableInit(mt)
	MergeSort(s)
	return 0
}

/*
READ OPERATIONS
*/

// Get : main function to get a key's value from the key-value store
func Get(mt *Memtable, key string) ([]byte, error) {
	mResults, err := SearchMemtable(mt, key)
	dResults, err := SearchDisk(key)

	if err != nil {
		return nil, errors.New("arghhhh")
	}

	//[TODO] catch errors properly
	results, err := MergeResults(mResults, dResults)

	if err != nil {
		return nil, errors.New("oh no")
	}

	return results, nil
}

// SearchDisk : searches SStables for a key
func SearchDisk(key string) ([]byte, error) {
	results := []byte{1, 2, 3}
	return results, nil
}

// MergeResults : combines results of memtable and disk searches to find current value of key
func MergeResults(mresults []byte, dresults []byte) ([]byte, error) {
	results := append(mresults, dresults...)
	return results, nil
}
