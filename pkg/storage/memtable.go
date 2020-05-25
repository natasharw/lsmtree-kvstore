package storage

import (
	"errors"
	"log"
)

var initPairs []*Pair

// MemtableInit : creates a memtable in local memory
func MemtableInit() (m *LsmNode) {
	log.Printf("Creating a memtable (C0)")

	memtable := &LsmNode{
		fname:    "",
		lvl:      0,
		kv:       initPairs,
		approxSz: 0,
		maxSz:    MaxMem,
		minKey:   0,
		maxKey:   0}
	log.Printf("New memtable successfully created")

	return memtable
}

/*
WRITE OPERATIONS
*/

// InsertToMemtable : inserts a key and value to memtable
func (lsmtree *LsmTree) InsertToMemtable(pair *Pair) error {

	log.Printf("Handling setting key %v in memtable", pair.key)

	memtable := lsmtree.lvls[0].files[0]
	existing, present, err := memtable.SearchNode(pair.key)

	if !present {
		err := lsmtree.insertPairToMemtable(pair)
		if err != nil {
			return errors.New("Failed to insert key in memtable")
		}

		lsmtree.lvls[0].files[0].updateMinOrMax(pair.key)

	} else {
		err = lsmtree.setValueOnPair(existing, pair.value)
		if err != nil {
			return errors.New("Failed to set value against key in memtable")
		}
	}

	return nil
}

// insertKey : inserts key in correct place in memtable using binary search
func (lsmtree *LsmTree) insertPairToMemtable(pair *Pair) error {
	log.Printf("Finding correct place for key %v in memtable...", pair.key)

	memtable := lsmtree.lvls[0].files[0]
	memtable.kv = append(memtable.kv, pair)

	memtable.approxSz++
	// [TODO] - binary search to insert key in correct place

	return nil
}

// setValueOnKey : sets a value against a key in a memtable
func (lsmtree *LsmTree) setValueOnPair(pair *Pair, value int) error {
	log.Printf("Setting value of key %v in memtable", pair.key)
	pair.value = value
	log.Printf("Successfully set key %v to %v", pair.key, value)

	return nil
}

func (node *LsmNode) updateMinOrMax(key int) {
	if key < node.minKey {
		node.minKey = key
		log.Printf("Updated min key of node to %v", key)
	} else if key > node.maxKey {
		node.maxKey = key
		log.Printf("Updated max key of node to %v", key)
	}
}

// IsMemtableFull : returns true if memtable entries equals max constant
func (lsmtree *LsmTree) IsMemtableFull() bool {
	log.Printf("Checking if memtable is full")

	memtable := lsmtree.lvls[0].files[0]
	if memtable.approxSz < memtable.maxSz {
		log.Printf("Memtable not full. Approx size: %v Max size: %v", memtable.approxSz, memtable.maxSz)

		return false
	}
	log.Printf("Memtable is full. Approx size: %v Max size: %v", memtable.approxSz, memtable.maxSz)

	return true
}

// /*
// READ OPERATIONS - stored in Lsm file
// */
