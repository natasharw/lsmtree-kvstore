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

// SSTable : represents a Sorted String Table on disk (kv : key-value pairs; index : key-offset pairs)
type SSTable struct {
	fname string
	kv    map[string][]byte
	index map[string][]byte
}

/*
WRITE OPERATIONS
*/

// SSTableInit : creates new SStable from a sorted set of key-value pairs
func SSTableInit(kv map[string][]byte) (sst *SSTable) {

	log.Printf("Creating new SSTable")

	fname := sst.createFileName()

	return &SSTable{
		fname: fname,
		kv:    kv,
		index: make(map[string][]byte)}

	// TODO group records into blocks?
}

// createFileName : set up file name where data will be written to
func (sst *SSTable) createFileName() string {
	log.Printf("Creating file name")

	fileName := filepath.Join(
		"lsmtree",
		fmt.Sprintf("%v.dat", time.Now().Unix()),
	)
	log.Printf("Filename created: %v", fileName)

	return fileName
}

// WriteSSTableToDisk : writes a SStable to disk
func (kvs *KVStore) WriteSSTableToDisk(sst *SSTable) error {
	log.Printf("Writing SSTable to disk...")

	log.Printf("Creating new file")
	file, err := os.Create(sst.fname)
	if err != nil {
		panic(err)
	}

	log.Printf("Encoding SSTable data...")
	sstEncoded := Encode(sst)

	log.Printf("Checking status of level one of LSM tree")
	_, present := kvs.manifest.components[1]

	if !present {
		log.Printf("Level one of LSM tree does not exist yet. Flushing immediately")
		nb, err := file.Write(sstEncoded)
		if err != nil {
			return err
		}

		log.Printf("Number of bytes written: %d", nb)
	} else {
		kvs.MergeSSTable(sst)
	}

	return nil
}

// MergeSSTable : merges an SStable into existing disk structure
func (kvs *KVStore) MergeSSTable(ssy *SSTable) error {

	maxComponents := 1 * kvs.fanout // TODO - determine when / how a level has reached max components

	if kvs.manifest.components[1] == maxComponents {
		// TODO - level 1 merged into level 2
		// C0 is now level 1
	}

	return nil
}

/*
READ OPERATIONS
*/

// GetValue : read disk from approximate starting point to locate key and return its value
func (sst *SSTable) GetValue(key string) []byte {

	log.Printf("Creating key-value from SSTable: %s", sst.fname)

	log.Printf("Opening file")
	file, err := os.Open("lsmtree/1589638784.dat") // TODO - change this to open a sst.fname
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Creating new file scanner")
	scanner := bufio.NewScanner(file)
	log.Printf("Scanning...")
	scanner.Scan()

	// text := scanner.Text()
	data := scanner.Bytes()

	sstable := SSTable{}

	log.Printf("Decoding")
	decoded := json.Unmarshal(data, sstable) //TODO
	// decoded := DecodeToSST(data)
	log.Printf("Decoded")
	fmt.Printf("%v", decoded)

	// TODO - can access the SSTable as bytes - need to decode back to SSTable format first or access key-value and decode only this?

	return data
}
