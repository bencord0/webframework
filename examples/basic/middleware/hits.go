package middleware

import (
    "net/http"
    "sync"
    . "github.com/bencord0/webframework"
    "github.com/bluele/gcache"
)

// Hitcounter | Request middleware
// -----------------------------------------------------------------------------------------------
type HitCounter struct {
    Cache *gcache.Cache
    mutex sync.Mutex
}

func (m HitCounter) ProcessView(request *http.Request, view ViewFunc) *http.Response {
    if m.Cache == nil {
        return nil
    }

    var hitcount int
    cache := *m.Cache

    m.mutex.Lock()
    defer m.mutex.Unlock()

    result, err := cache.Get("hits")
    if err != nil {
        result = 0
    }

    hitcount, ok := result.(int)
    if !ok {
        hitcount = 0
    }

    hitcount += 1
    cache.Set("hits", hitcount)

    // Pass through to the next middleware
    return nil
}
