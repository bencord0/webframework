package main

import (
    . "github.com/bencord0/webframework"
    "github.com/bencord0/webframework/examples/basic/middleware"
)

var settings = Settings{
    Addr: "0.0.0.0:8000",
    Urls: []Url{
        {"/", index},
        {"/health", health},
        {"/hits", hits},
        {"/ua", ua},
        {"/uuid", uuid},
    },
    Middleware: []Middleware{
        middleware.Pratchett{},
        middleware.Uppercase{},
        middleware.UUIDify{},
        middleware.HitCounter{Cache: &cache},
    },
}
