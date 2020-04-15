# LSM-tree key/value store

An attempt at implementing a primitive LSM-tree key/value store in Go

## Overview

A rudimentary write-optimised key/value store using the Go standard library, with a LSM-tree (Log-Structured Merge tree) structure. Simple set key and get key operations are accessible from HTTP endpoints. <b><i>This project is for learning only and is WIP. </i></b>

<b> LSM-tree overview </b>  
* Deal with fast sequential writes well
* Searching for existing records does not need to happen on write
* Key-value pairs initially cached in local memory `Memtable (C0)`
* `Memtable` gets flushed to disk as an immutable `SStable` (sorted string table) once full 
* Merging and compaction algorithms are used to store `SSTables` efficiently between levels of tree (`C1`-`Cn`)
* Good for write-heavy workloads such as transactional logs



## Prerequisites
- Install [Go](https://golang.org/doc/install)

## Install
#### Clone repository
`$ git clone https://github.com/natasharw/lsmtree-kvstore.git`

## Usage
#### Run the server
`$ go run main.go`
#### Set a key (WIP)
Go to `localhost:8000/set?yourkey=yourvalue` in browser
#### Get a key (WIP)
Go to `localhost:8000/get?key=yourkey` in browser


## Set key explained
* Key-value pair initially cached in local memory in a `Memtable (C0)`
* The `Memtable` is a mutable data structure which stores sorted key-value pairs
* Periodically, `Memtable (C0)` is flushed to disk, becoming an immutable sorted string table (`SSTable`)
* The flush will be triggered after a certain time period or when the cache reaches a certain size
* A `SSTable` contains data (key-value pairs) and an index (key-offset pairs)
* The flush of data from `C0` to lower components on disk (`C1`-`Cn`) will happen using a `MergeSort()` process.
* In this simple implementation, each component is designed as a sorted array (see [Future Improvements](#future-improvements))
* `SSTable` layout is periodically re-arranged using a process called compaction. This prevents reads from becoming less optimised over time as the key-value pairs become less efficiently stored

## Get key explained
* Check if the key exists in the in-memory `Memtable`
* Starting from the bottom `SSTable`, check the index then read from relevant tables
* Merge all results and return results to the browser

## Testing efficiency

## Future improvements
#### Add Bloom filters  
A Bloom filter could be added to each `SSTable` to identify the keys that probably exist in that `SSTable`. The Bloom filters could be cached in memory and would prevent unnecessary files being read when trying to retrieve a key.
#### Add DELETE key functionality
#### Concurrent execution  
Use goroutines to support multiple concurrent reads or writes
#### Redesign storage components (`C0`, `C1`-`Cn`) as individually optimised data structures
E.g. `Memtable (C0)`: sorted array -> binary/AVL tree
