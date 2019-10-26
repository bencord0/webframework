package middleware

import (
    "net/http"
)

// GNU Terry Pratchett Middleware | Response middleware
// -----------------------------------------------------------------------------------------------
type Pratchett struct {}

func (m Pratchett) ProcessResponse(request *http.Request, response *http.Response) *http.Response {
    if response.Header == nil {
        header := make(map[string][]string)
        response.Header = header
    }

    response.Header["X-Clacks-Overhead"] = []string{"GNU Terry Pratchett"}
    return response
}

