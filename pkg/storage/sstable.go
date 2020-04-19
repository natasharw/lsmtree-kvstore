package storage

import (
	"fmt"
	"log"
	"os"
)

// SSTable : represents a Sorted String Table on disk
// d : key-value pairs (sequence of blocks); bs : block size; index : key-offset pairs
type SSTable struct {
	// [TODO] - change this to sequence of blocks
	fname string
	kv    map[string][]byte
	bs    int
	index map[string][]byte
}

type block struct {
	d map[string][]byte
}

const (
	// blockSz = size of blocks in kb
	blockSz = 32
)

/*
WRITE OPERATIONS
*/

// NewSSTable : creates a new SStable from a given memtable
func (mt *Memtable) NewSSTable() (ss *SSTable) {
	// need to create a new memtable instance while I am doing this
	log.Printf("Creating a new SSTable")
	log.Printf("SSTable attributes. Block size: BLOCK_SIZE")

	//TODO group records into blocks and write to disk

	return &SSTable{
		fname: "",
		kv:    mt.kv,
		bs:    blockSz,
		index: make(map[string][]byte)}
}

// WriteSSTableToDisk : writes a SStable to disk
func (ss *SSTable) WriteSSTableToDisk() error {
	log.Printf("Writing SSTable to disk")

	ss.fname = ss.createFileName()
	file, err := os.Create(ss.fname)

	if err != nil {
		panic(err)
	}

	nb, err := file.Write([]byte{1, 2, 3})
	log.Printf("Number of bytes written: %d", nb)

	// TODO encode sstable data to bytes?
	// if n, err := file.Write(); err != nil {
	// 	return err
	// }

	ss.MergeSort()

	return nil
}

// createFileName : set up file name where data will initially be written to
func (ss *SSTable) createFileName() string {

	return "foo.data"
}

// MergeSort : merges an SStable into existing disk structure
func (ss *SSTable) MergeSort() error {
	// TODO - implement merge sort - combine with compaction func?

	return nil
}

/*
READ OPERATIONS
*/

// // SearchSSTable : search in index of one SSTable for given key
// func (*SSTable) SearchSSTable(key string) []byte {
// 	return []byte{1, 2, 3}
// }

// GetValue : read disk from approximate starting point to locate key and return its value
func (ss *SSTable) GetValue(key string) []byte {
	file, err := os.Open(ss.fname)
	bArr := make([]byte, 10)
	value, err := file.Read(bArr)
	file.Close()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", value)

	return []byte{1, 2}
}
