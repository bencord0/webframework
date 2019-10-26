package main

import "github.com/bluele/gcache"

var cache gcache.Cache

func init() {
    cache = gcache.New(128).LRU().Build()
}
