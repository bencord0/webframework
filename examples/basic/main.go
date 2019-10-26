package main

import (
    . "github.com/bencord0/webframework"
)

func main() {
    application := NewApplication(settings)
    application.Run()
}
