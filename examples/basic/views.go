package main

import (
    "fmt"
    "net/http"
    . "github.com/bencord0/webframework"
)

func index(request *http.Request) *http.Response {
    return TextResponse("hello world!", http.StatusOK, nil)
}

func health(request *http.Request) *http.Response {
    return TextResponse("healthy", http.StatusOK, nil)
}

func uuid(request *http.Request) *http.Response {
    uuid := request.Header.Get("UUID")
    if uuid != "" {
        return TextResponse(fmt.Sprintf("uuid: %s", uuid), http.StatusOK, nil)
    }

    return TextResponse("No uuid found", http.StatusInternalServerError, nil)
}

func ua(request *http.Request) *http.Response {
    useragent := request.UserAgent()
    return TextResponse(fmt.Sprintf("user agent: %s", useragent), http.StatusOK, nil)
}

func hits(request *http.Request) *http.Response {
    type HitCounter struct {
        Hits int `json:hits`
    }

    result, err := cache.Get("hits")
    if err != nil {
        return JSONResponse(HitCounter{}, http.StatusInternalServerError, nil)
    }

    hitcount, ok := result.(int)
    if !ok {
        return JSONResponse(HitCounter{}, http.StatusInternalServerError, nil)
    }

    return JSONResponse(HitCounter{
        Hits: hitcount,
    }, http.StatusOK, nil)
}
