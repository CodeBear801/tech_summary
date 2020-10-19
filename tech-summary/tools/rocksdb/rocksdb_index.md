# RocksDB

## Why
Very good in write-heavy workload as well as low-latency read workload

## What
- Key-Value persistent store
- Embedded storage of application, storage engine of database
- Point/Range lookup
- Optimized for flash

## Interface
- Keys and Values are byte arrays
- Keys have total order
- Update : Put/Delete/Merge
- Queries : Get/Iterator
- Consistency: atomic multi-put, multi-get, iterator, snapshot read, transactions

 