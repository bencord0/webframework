package webframework

import "net/http"

type Middleware interface{}
type RequestMiddleware interface {
	ProcessView(*http.Request, ViewFunc) *http.Response
}
type ResponseMiddleware interface {
	ProcessResponse(*http.Request, *http.Response) *http.Response
}

type ViewFunc = func(*http.Request) *http.Response
