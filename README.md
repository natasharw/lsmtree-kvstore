# LSM-tree key/value store

An attempt at implementing a primitive LSM-tree key/value store in Go

## Overview


A rudimentary write-optimised key/value store using the Go standard library, with a LSM-tree (Log-Structured Merge tree) structure. <b>For learning only and is currently WIP.</b>


<b>Some details</b>
* Key-value pairs initially cached in local memory `Memtable (L0)`
* Once full, `Memtable (L0)` gets flushed to disk as a `SSTable` (sorted string table) which is immutable
* A new `SStable` gets merged into the tree which is a sequence of levels increasing in size (`L1`-`Ln`)
* `SSTables` on disk are periodically re-arranged by compaction
* Get requests search in `Memtable (L0)` first then sequentially down each tree level returning first value if found


## Install
```
$ git clone https://github.com/natasharw/lsmtree-kvstore.git
```
```
$ cd lsmtree-kvstore
```

## Use

#### Run
```
$ go run cmd/server/main.go
```
#### Set key
```
$ set <yourkey> <yourvalue>
```
#### Get key
```
$ get <yourkey>
```
#### Exit
```
$ exit
```
