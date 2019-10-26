package main

import (
    . "github.com/bencord0/webframework"
)

var application = NewApplication(settings)

func main() {
    application.Run()
}
