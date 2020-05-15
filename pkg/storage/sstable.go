package storage

import (
	"fmt"
	"log"
	"os"
)

// SSTable : represents a Sorted String Table on disk (kv : key-value pairs; index : key-offset pairs)
type SSTable struct {
	fname string
	kv    map[string][]byte
	index map[string][]byte
}

/*
WRITE OPERATIONS
*/

// NewSSTable : creates a new SStable from a given memtable
func (mt *Memtable) NewSSTable() (ss *SSTable) {

	log.Printf("Creating a new SSTable")

	// TODO group records into blocks and write to disk
	// create a new memtable instance while I am doing this

	return &SSTable{
		fname: "",
		kv:    mt.kv,
		index: make(map[string][]byte)}
}

// WriteSSTableToDisk : writes a SStable to disk
func (kvs *KVStore) WriteSSTableToDisk(ss *SSTable) error {
	log.Printf("Writing SSTable to disk")

	log.Printf("Creating file for SSTable")
	ss.fname = ss.createFileName()
	file, err := os.Create(ss.fname)

	log.Printf("Encoding SSTable data to bytes")
	ssEncoded := Encode(ss)

	if err != nil {
		panic(err)
	}

	log.Printf("Checking status of level one of LSM tree")
	_, present := kvs.components[1]

	if !present {
		log.Printf("Level one of LSM tree does not exist yet. Flushing immediately")
		nb, err := file.Write(ssEncoded)

		if err != nil {
			return err
		}

		log.Printf("Number of bytes written: %d", nb)
	} else {
		kvs.MergeSSTable(ss)
	}

	return nil
}

// createFileName : set up file name where data will be written to
func (ss *SSTable) createFileName() string {
	// TODO
	return "0100.data"
}

// MergeSSTable : merges an SStable into existing disk structure
func (kvs *KVStore) MergeSSTable(ss *SSTable) error {

	maxComponents := 1 * kvs.fanout // TODO - determine when / how a level has reached max components

	if kvs.components[1] == maxComponents {
		// TODO - level 1 merged into level 2
		// C0 is now level 1
	}

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
