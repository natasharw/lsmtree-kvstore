package storage

import (
	"log"
	"os"
)

const (
	// grwthFctr : ratio between levels of LSM-tree e.g. C2 = C1 * fanout
	grwthFctr int = 2
)

// Storage : generic store must be able to get and set key-value pairs
type Storage interface {
	Get(key int) (int, error)
	Set(key int, value int) error
}

// Pair : a key-value pair
type Pair struct {
	key   int
	value int
}

// LsmNode : a file / SSTable or the buffer
type LsmNode struct {
	fname    string
	lvl      int
	kv       []Pair
	approxSz int
	maxSz    int
	minKey   int // the min key in the file
	maxKey   int // the max key in the file
}

// LsmLevel : a conceptual hierachical level which comprises of a bunch of nodes
type LsmLevel struct {
	lvl       int
	files     []*LsmNode // an array of pointers to actual files within level
	fileCount int
	next      *LsmLevel
}

// LsmTree : the whole tree
type LsmTree struct {
	lvls      []LsmLevel
	lvlCount  int
	grwthFctr int
	entries   int
}

// LsmTreeInit : initalises a LSM tree
func LsmTreeInit() *LsmTree {
	log.Printf("Instantiating a new LSM tree...")

	memtable := MemtableInit()
	initComponents := []*LsmNode{}
	initComponents = append(initComponents, memtable)

	l0 := LsmLevel{
		lvl:       0,
		files:     initComponents,
		fileCount: 1,
		next:      nil}

	initLevels := []LsmLevel{}
	initLevels = append(initLevels, l0) // init a memtable component, and this is only initial level

	log.Printf("Initial attributes: Memtable added as buffer (CO/LO), level growth factor %d", grwthFctr)
	lsmtree := &LsmTree{
		lvls:      initLevels,
		lvlCount:  1,
		grwthFctr: grwthFctr,
		entries:   0}

	log.Printf("Creating directory lsmtree/ for disk storage")
	os.Mkdir("lsmtree", 0777)

	log.Printf("LSM tree succesfully instantiated")
	return lsmtree
}

// Set : Main function to set a key in the key-value store
func (lsmtree *LsmTree) Set(key int, value int) error {
	log.Printf("Setting key %v to value %v", key, value)

	pair := new(Pair)
	pair.key = key
	pair.value = value
	lsmtree.InsertToMemtable(pair)

	f := lsmtree.IsMemtableFull()
	if f {
		err := lsmtree.Flush()
		if err != nil {
			return err
		}
	}

	log.Printf("Congrats! %v: %v successfully stored", key, value)

	return nil
}

// Flush : flushes memtable to disk as SSTable
func (lsmtree *LsmTree) Flush() error {
	log.Printf("Starting flush...")

	memtable := lsmtree.lvls[0].files[0]
	sst := lsmtree.SSTableFromMemtable(memtable)
	lsmtree.lvls[0].files[0] = MemtableInit() // create new empty memtable
	log.Printf("Stored memtable and fresh empty memtable created. Ready for write")

	lsmtree.WriteSSTableToDisk(sst)

	return nil
}

// Get : main function to get a key's value from the key-value store
func (lsmtree *LsmTree) Get(key int) (int, error) {
	log.Printf("Getting key %v", key)

	memtable := lsmtree.lvls[0].files[0]
	result, present, err := memtable.SearchNode(key)

	if err != nil {
		return 0, err
	}

	if !present {
		result, err = lsmtree.searchDisk(key)
		if err != nil {
			return 0, err
		}
	}

	value := result.value

	return value, nil
}

// SearchNode : searches in the set of key-value pairs of a node for a given key
func (lsmnode *LsmNode) SearchNode(key int) (*Pair, bool, error) {
	log.Printf("Searching node for key %v...", key)

	pairs := lsmnode.kv

	for i := range pairs {
		if pairs[i].key == key {
			log.Printf("Found value %v for key %v", pairs[i].value, key)
			return &pairs[i], true, nil
		}
	}
	// TODO - cursor to skip over file to relevant starting place

	log.Printf("Key %v not found in memtable", key)

	return nil, false, nil
}

// searchDisk : performs search on-disk for a given key
func (lsmtree *LsmTree) searchDisk(key int) (*Pair, error) {
	log.Printf("Starting disk search...")

	indices, err := lsmtree.loadIndices()
	if err != nil {
		panic(err)
	}

	ir, err := lsmtree.searchIndices(indices)
	result := lsmtree.getFromDisk(ir, key)

	return result, nil
}

// loadIndices : loads indices of SStables into local memory
func (lsmtree *LsmTree) loadIndices() (map[string][]byte, error) {
	log.Printf("Loading indices...")

	indices := make(map[string][]byte)

	// TODO ; for each SSTable, load its block index (key-off pairs) into a structure in local memory
	// load all -> search one by one or load one -> search one ?

	return indices, nil
}

// searchIndices : searches through key-offset pairs of all sstables in local memory
func (lsmtree *LsmTree) searchIndices(map[string][]byte) (map[string]map[string][]byte, error) {
	log.Printf("Searching indices...")

	var r map[string]map[string][]byte

	// TODO - search through SStables from top level to bottom
	result := r

	return result, nil
}

// getFromDisk : returns a value for a key given a ssblock identifier, and key's approx location
func (lsmtree *LsmTree) getFromDisk(ir map[string]map[string][]byte, key int) *Pair {
	log.Printf("Getting key-value from disk...")

	sst := new(LsmNode) // TODO
	results := sst.GetPair(key)

	// for _, r := range ir {
	// 	for tbl := range r {
	// 		print(tbl)
	// 		ss := new(SSTable) // TODO dummy overwrite -> change this to the search indices result
	// 		results = ss.GetValue(key)
	// 	}
	// }

	return results
}
