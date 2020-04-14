package storage

// SSTable : represents a Sorted String Table on disk
type SSTable struct {
	d map[string][]byte
}

/*
WRITE OPERATIONS
*/

// SSTableInit : creates a new SStable from a given memtable
func SSTableInit(m *Memtable) (ss *SSTable) {
	return &SSTable{d: make(map[string][]byte)}
}

// MergeSort : and merges an SStable into existing disk structure
func MergeSort(*SSTable) (int, error) {
	return 0, nil
}

/*
READ OPERATIONS
*/

// SearchSSTable : search in index of one SSTable for given key
func SearchSSTable(key string) []byte {
	return []byte{1, 2, 3}
}

/*
OTHER OPERATIONS
*/

// Compaction : reorganises SSTables in order to
func Compaction() (int, error) {
	// reorganize the data to make it more efficient
	// undertake level-based compaction
	// when data reaches growthFactor --> merge into row below
	return 0, nil
}
