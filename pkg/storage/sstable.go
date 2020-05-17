package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// /*
// WRITE OPERATIONS
// */

// SSTableFromMemtable : creates a new SSTable node given memtable data
func (lsmtree *LsmTree) SSTableFromMemtable(memtable *LsmNode) (sst *LsmNode) {

	log.Printf("Creating new SSTable")

	fname := sst.createFileName()

	sstable := &LsmNode{
		fname:    fname,
		lvl:      memtable.lvl + 1,
		kv:       memtable.kv,
		approxSz: memtable.approxSz,
		maxSz:    memtable.maxSz,
		minKey:   memtable.minKey,
		maxKey:   memtable.maxKey}

	return sstable
}

// createFileName : set up file name where data will be written to
func (sst *LsmNode) createFileName() string {
	log.Printf("Creating file name")

	fileName := filepath.Join(
		"lsmtree",
		fmt.Sprintf("%v.dat", time.Now().Unix()),
	)
	log.Printf("Filename created: %v", fileName)

	return fileName
}

// WriteSSTableToDisk : writes a SStable to disk
func (lsmtree *LsmTree) WriteSSTableToDisk(sst *LsmNode) error {
	log.Printf("Writing SSTable to disk...")

	log.Printf("Creating new file")
	file, err := os.Create(sst.fname)
	if err != nil {
		panic(err)
	}

	log.Printf("Encoding SSTable data")
	sstEncoded, err := json.Marshal(sst)
	if err != nil {
		return err
	}

	log.Printf("Checking status of level one of LSM tree")

	if len(lsmtree.lvls) < 2 { // TODO - change to map index check
		log.Printf("Level one of LSM tree does not exist yet. Flushing immediately")
		nb, err := file.Write(sstEncoded)
		if err != nil {
			return err
		}

		log.Printf("Number of bytes written: %d", nb)
	} else {
		lsmtree.MergeSSTable(sst)
	}

	return nil
}

// MergeSSTable : recursively merge an SStable into existing disk structure
func (lsmtree *LsmTree) MergeSSTable(sst *LsmNode) error {

	desiredLevel := sst.lvl + 1
	maxComponents := desiredLevel * lsmtree.grwthFctr

	// base case
	if maxComponents > len(lsmtree.lvls[1].files) {
		// sstable can be moved to level below with no drama.
		return nil
	}

	// TODO - else we need to flowing merge the components on the level below until one of them has not got max components
	return nil
}

/*
READ OPERATIONS
*/

// GetPair : read disk from approximate starting point to locate key and return its pair
func (sst *LsmNode) GetPair(key int) *Pair {

	log.Printf("Creating key-value from SSTable: %s", sst.fname)

	log.Printf("Opening file")
	file, err := os.Open("lsmtree/1589638784.dat") // TODO - change this to open a sst.fname
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		} // close the file as destructor
	}()

	log.Printf("Creating new file scanner")
	scanner := bufio.NewScanner(file)
	log.Printf("Scanning...")
	scanner.Scan()

	data := scanner.Bytes() // text := scanner.Text()

	sstable := LsmNode{}

	log.Printf("Decoding")
	decoded := json.Unmarshal(data, sstable)
	fmt.Printf("%v", decoded)
	pair, presence, err := sstable.SearchNode(key) // TODO - sstable is decoded data

	if !presence {
		log.Printf("No key-value pair for key %v found :(", key)
		return nil
	}

	return pair
}
