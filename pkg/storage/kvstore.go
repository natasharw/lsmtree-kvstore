package storage

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	// fanout : ratio between levels of LSM-tree e.g. C2 = C1 * fanout
	fanout int = 2
)

// Store : generic store must be able to get and set key-value pairs
type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value string) error
}

// KVStore : represents the entirety of a key-value store
type KVStore struct {
	buffer   *Memtable
	fanout   int
	manifest manifest
}

// manifest : contains indexes and meta-data about LSMtree
type manifest struct {
	components  map[int]int
	sstableList map[int]string
}

// KVStoreInit : instantiates a new key-value store with some default values
func KVStoreInit() *KVStore {
	log.Printf("Instantiating a new key-value store...")
	log.Printf("Setting attributes: Memtable as buffer and level growth factor of %d", fanout)

	mem := MemtableInit()
	mani := manifestInit()

	kvs := &KVStore{
		buffer:   mem,
		fanout:   fanout,
		manifest: mani}

	log.Printf("Setting initial manifest data")
	kvs.manifest.components[0] = 1 // memtable buffer represents one component in top level of kvs

	log.Printf("Creating directory lsmtree/ for SSTable files")
	os.Mkdir("lsmtree", 0777)

	log.Printf("Key-value store succesfully instantiated")
	return kvs
}

// ManifestInit : creates a new manifest file, empty apart from C0
func manifestInit() manifest {
	log.Printf("Creating manifest")

	m := manifest{
		components:  make(map[int]int),
		sstableList: make(map[int]string)}
	log.Printf("Manifest successfully created")

	return m

}

/*
WRITE OPERATIONS
*/

// Set : Main function to set a key in the key-value store
func (kvs *KVStore) Set(key string, value string) error {
	log.Printf("Setting key %s", key)
	mt := kvs.buffer

	val := Encode(value)
	mt.InsertToMemtable(key, val)

	f := mt.IsMemtableFull()
	if f {
		err := kvs.Flush()
		if err != nil {
			return err
		}
	}

	log.Printf("Congrats! %s: %s successfully stored", key, value)

	return nil
}

// Flush : flushes memtable to disk as SSTable
func (kvs *KVStore) Flush() error {

	log.Printf("Starting flush...")
	ss := SSTableInit(kvs.buffer.kv)
	kvs.buffer = MemtableInit()
	log.Printf("Stored memtable and fresh empty memtable created")

	kvs.WriteSSTableToDisk(ss)

	return nil
}

// Compaction : takes a set of files sorted by key and returns a new set of non-overlapping files sorted by key
// might be called when level reaches certain threshold or for periodic reorganising of files on disk to maintain efficient
func (kvs *KVStore) Compaction(inputFilesDir string) (string, error) {

	log.Printf("Starting compaction...")
	log.Printf("Compacting files stored at tmp file dir %s", inputFilesDir)

	log.Printf("Creating new tmp dir to store output")
	now := time.Now().Unix()
	outputFilesDir := fmt.Sprintf("outputFilesDir-%v", now)
	os.Mkdir("outputFilesDir-", 0777)
	log.Printf("New tmp dir created: %s", outputFilesDir)

	// TODO - DO THE COMPACTION (MERGE SORT)

	log.Printf("Compaction complete. Returning new tmp file dir and remvoing old tmp file dir", outputFilesDir)
	defer os.RemoveAll("inputFilesDir")

	return outputFilesDir, nil

	// if l1 is None: return l2
	// if l2 is None: return l1
	// if l1.next is None and l2.next is None:
	// 	if l1.val > l2.val:
	// 		return ListNode(l2.val, next=l1)
	// 	else:
	// 		return ListNode(l1.val, next=l2)
	// elif l1.val < l2.val:
	// 	return ListNode(l1.val,self.mergeTwoLists(l1.next,l2))
	// elif l2.val < l1.val:
	// 	return ListNode(l2.val,self.mergeTwoLists(l1,l2.next))
	// elif l1.val == l2.val:
	// 	return ListNode(l1.val,self.mergeTwoLists(l1.next,l2))
	// return nil
}

/*
READ OPERATIONS
*/

// Get : main function to get a key's value from the key-value store
func (kvs *KVStore) Get(key string) ([]byte, error) {
	var results []byte

	results, present, err := kvs.buffer.SearchMemtable(key)
	if err != nil {
		return nil, err
	}

	if !present {
		results, err = kvs.searchDisk(key)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

// searchDisk : performs search on-disk for a given key
func (kvs *KVStore) searchDisk(key string) ([]byte, error) {
	log.Printf("Starting disk search...")

	indices, err := kvs.loadIndices()
	if err != nil {
		panic(err)
	}

	ir, err := kvs.searchIndices(indices)
	result := kvs.getFromDisk(ir, key)

	return result, nil
}

// loadIndices : loads indices of SStables into local memory
func (kvs *KVStore) loadIndices() (map[string][]byte, error) {

	log.Printf("Loading indices...")

	indices := make(map[string][]byte)

	// TODO ; for each SSTable, load its block index (key-off pairs) into a structure in local memory
	// load all -> search one by one or load one -> search one ?

	return indices, nil
}

// searchIndices : searches through key-offset pairs of all sstables in local memory
func (kvs *KVStore) searchIndices(map[string][]byte) (map[string]map[string][]byte, error) {

	log.Printf("Searching indices...")

	var r map[string]map[string][]byte

	// TODO - search through SStables from top level to bottom

	result := r

	return result, nil
}

// getFromDisk : returns a value for a key given a ssblock identifier, and key's approx location
func (kvs *KVStore) getFromDisk(ir map[string]map[string][]byte, key string) []byte {

	log.Printf("Getting key-value from disk...")
	// TODO

	var results []byte

	sst := new(SSTable)
	results = sst.GetValue(key)

	// for _, r := range ir {
	// 	for tbl := range r {
	// 		print(tbl)
	// 		ss := new(SSTable) // TODO dummy overwrite -> change this to the search indices result
	// 		results = ss.GetValue(key)
	// 	}
	// }

	return results
}
