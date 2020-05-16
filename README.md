# LSM-tree key/value store

An attempt at implementing a primitive LSM-tree key/value store in Go

## Overview

A rudimentary write-optimised key/value store using the Go standard library, with a LSM-tree (Log-Structured Merge tree) structure. <b>For learning only and is currently WIP.</b>

<b>Some details</b>
* Key-value pairs initially cached in local memory `Memtable (C0/L0)` implemented here as a sorted map of key-value pairs
* Once full, `Memtable (C0)` gets flushed to disk as a `SSTable` (Sorted String Table) file which is immutable
* The new `SSTable` gets merged into the tree which is a sequence of levels increasing in size (`L1`-`Ln`)
* Compaction process is used to merge `SSTables` to lower levels as well as periodic re-organisation within levels to maintain efficient storage
* Get requests search in `Memtable (C0)` first then sequentially down each tree level returning first value if found


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
