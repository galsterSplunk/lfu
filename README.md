# LFU Cache

[![Latest Version](http://img.shields.io/github/release/mtchavez/lfu.svg?style=flat-square)](https://github.com/mtchavez/lfu/releases)
[![Build Status](https://travis-ci.org/mtchavez/lfu.svg)](https://travis-ci.org/mtchavez/lfu)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/lfu)

LFU cache in Go offering O(1) Get and Insert [outlined here](http://dhruvbird.com/lfu.pdf).

## Install

`go get -u github.com/mtchavez/lfu`

## Usage

```
package main

import (
    "fmt"
    "github.com/mtchavez/lfu"
)

func main() {
    cache := lfu.NewLFU()

    // Insert
    cache.Insert(42, []byte("user:42:user@example.com"))

    // Insert existing key
    success, err := cache.Insert(42, "Nope")
    if !success {
        fmt.Sprintf("Error inserting:", err)
    }

    var data interface{}
    var e error
    // Get
    data, e = cache.Get(42)
    fmt.Sprintf("Data for 42 is", data)

    // Get not found
    data, e = cache.Get(987654321)
    if e != nil {
        fmt.Sprintf("Error on get:", e)
    }
}

```

## Tests

`go test --cover`

## Benhcmarks

Get and insert methods are benchmarked. Results from OS X with
a 2.3 GHz Intel Core i7:

```
$ go test --bench=.*
PASS
BenchmarkInsert	 1000000	      1943 ns/op
BenchmarkGet_EmptyCache	 2000000	       768 ns/op
BenchmarkGet_AllMisses	 3000000	       534 ns/op
BenchmarkGet_AllHits	 1000000	      1118 ns/op
coverage: 100.0% of statements
```

## TODO

* Make threadsafe
* Some kind of eviction