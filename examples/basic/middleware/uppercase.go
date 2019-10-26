package middleware

import (
    "io"
    "io/ioutil"
    "strings"
    "net/http"
)

// Uppercase response | Request middleware, returning a response
// -----------------------------------------------------------------------------------------------
type Uppercase struct {}

func (m Uppercase) ProcessResponse(request *http.Request, response *http.Response) *http.Response{
    var buf strings.Builder
    defer response.Body.Close()
    io.Copy(&buf, response.Body)

    // Modify the response
    caps := strings.Title(buf.String())
    body := ioutil.NopCloser(strings.NewReader(caps))
    response.Body = body

    return response
}
