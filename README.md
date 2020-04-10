# LSM-tree key-value store

A rudimentary write-optimised key-value store using the Go standard library, with a LSM-tree (Log-Structured Merge tree) structure. Simple SET and GET key operations are accessible from HTTP endpoints. <b><i> This project is for learning only. </i></b>

<b> Why a LSM-tree design? </b>  
LSM-trees deal with fast sequential writes well as the data is append-only. Searching for existing records does not need to happen on write. There is a cost to read efficiency as data has to be collated from multiple places. This can be improved with an optimised file hierarchy through merging and compaction algorithms. Popular key-value stores like RocksDB use LSM-trees to provide efficiency for write-heavy workloads such as transactional logs.

## Pre-requisites
- Install [Go](https://golang.org/doc/install)

## Install
#### Clone repository
`$ git clone https://github.com/natasharw/lsmtree-kvstore.git`

## Usage
#### Run the server
`$ go run main.go`
#### Set a key
Go to `localhost:8000/get?key=yourkey` in browser
#### Get a key
Go to `localhost:8000/set?yourkey=yourvalue` in browser

## Set key explained
* Key-value pair initially cached in local memory in a memtable (C0)
* The memtable is a mutable data structure which sorts key-value pairs
* Periodically, memtable (C0) is flushed to disk, becoming an immutable Sorting Strings Table (SSTable)
* The flush will be triggered after a certain time period or when the cache reaches a certain size
* A SSTable contains data (key-value pairs) and an index (key-offset pairs)
* The flush of data from C0 to lower components on disk (C1+) will happen using a `MergeSort()` process. This can be implemented in several ways and the method chosen here is [TODO]
* In this simple implementation, each component is designed as a sorted array (see [Future Improvements](#future-improvements))
* SSTable layout is periodically re-arranged using a process called compaction. This prevents reads from becoming less optimised over time as the key-value pairs become less efficiently stored, due to deleted keys or data for one key being in multiple files

## Get key explained
* Check if the key exists in the in-memory memtable
* Starting from the bottom SSTable, check the index then read from relevant tables
* Merge all results and return results to the browser

## Testing efficiency

## Future improvements
#### Add Bloom filters  
A Bloom filter could be added to each SSTable to identify the keys that probably exist in that SSTable. The Bloom filters could be cached in memory and would prevent unnecessary files being read when trying to retrieve a key. Bloom filters sometimes return false positive, never false negative, results and they dramatically increase read performance from LSM-tree structures.
#### Add DELETE key functionality
#### Concurrent execution  
Use goroutines to support multiple concurrent reads or writes
#### Redesign storage components (C0, C1+) as optimised data structures
The memtable could be a balanced binary/AVL tree
