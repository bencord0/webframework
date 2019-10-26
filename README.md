# A Web Framework

An experiment to find other ways to interact with the Go http stack

## Usage

```

package main
import (
    "net/http"
    . "github.com/bencord0/webframework"
)

func index(*http.Request) *http.Response {
    return TextResponse("Hello World!\n", http.StatusOK, nil)
}

var settings := Settings{
    Addr: "0.0.0.0:8000",
    Urls: []Url{
        {"/", index},
    },
}

func main() {
    application := NewApplication(settings)
    aplication.Run()
}
```
