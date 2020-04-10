package storage

type SSTable struct {
	//define what it is
	d map[string][]byte
}

/*
WRITE OPERATIONS
*/

func SSTableInit(m *Memtable) (ss *SSTable) {
	return &SSTable{d: make(map[string][]byte)}
}

func MergeSort(*SSTable, *LsmTree) {
	// takes a newly created SSTable and merges it into the existing lsm tree
}

/*
READ OPERATIONS
*/

func SearchSSTable(key []byte) []byte {
	// given a key, search in the SSTable index to see if the key exists and return to
	return []byte{1, 2, 3}
}

/*
OTHER OPERATIONS
*/

func Compaction() {
	// reorganize the data to make it more efficient
	// undertake level-based compaction
	// when data reaches growthFactor --> merge into row below
	// this
}
