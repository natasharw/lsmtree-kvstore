package storage

import (
	"errors"
	"fmt"
)

const (
	MAX_NUM_LEVELS       int = 3
	LEVEL_GROWTH_FACTOR  int = 2 // magnification of data to level below e.g. C1 = C0 * 2
	MAX_MEMTABLE_ENTRIES int = 1000
)

type Storage interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
}

type LsmTree struct {
	buffer *Memtable
	levels int
	gf     int
}

/*
WRITE OPERATIONS
*/

func Put(key []byte, value []byte) (int, error) {
	// [TODO]
	status := LsmTreeExists("foo")
	if status == false {
		LsmTreeInit()
	}
	SetInMemtable(key, value)
	m := MemtableInit()
	switch f := IsMemtableFull(m); f {
	case true:
		Flush()
		return fmt.Printf("%s: %s added to buffer and buffer written to disk", key, value)
	case false:
		return fmt.Printf("%s: %s added to buffer", key, value)
	}

	return 0, nil
}

func LsmTreeExists(bar string) bool {
	// [TODO]
	return false
}

func LsmTreeInit() (lsmt *LsmTree) {
	return &LsmTree{
		buffer: MemtableInit(),
		levels: MAX_NUM_LEVELS,
		gf:     LEVEL_GROWTH_FACTOR,
	}

}

func Flush() int {
	m := MemtableInit()   // these will
	s := SSTableInit(m)   // absolutely not be
	tree := LsmTreeInit() // done here
	MergeSort(s, tree)
	return 0
}

/*
READ OPERATIONS
*/
func GetKey(key []byte) ([]byte, error) {
	m := MemtableInit() // this will absolutely not be done here
	mresults, err := SearchMemtable(key, m)
	dresults, err := SearchDisk(key)

	if err != nil {
		return nil, errors.New("arghhhh")
	}

	//[TODO] Catch errors properly
	results, err := MergeResults(mresults, dresults)

	if err != nil {
		return nil, errors.New("oh no")
	}

	return results, nil
}

func SearchDisk(key []byte) ([]byte, error) {
	// [TODO]
	// for each SSTable, SearchSSTable() and append to results slice
	results := []byte{1, 2, 3}
	return results, nil
}

func MergeResults(mresults []byte, dresults []byte) ([]byte, error) {
	// [TODO]
	// take all the results and work out what value the key currently has
	results := append(mresults, dresults...)

	return results, nil
}
