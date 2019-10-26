package webframework

import (
	"io"
	"net/http"
)

func (app application) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	urls := app.settings.Urls
	var response *http.Response
	var view ViewFunc
	var responseMiddleware []Middleware

	// Resolve which view to call
	for _, url := range urls {
		if url.Path == request.URL.Path {
			view = url.View
		}
	}

	if view == nil {
		// Can't find a view
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Apply middleware
	for _, middleware := range app.settings.Middleware {
		responseMiddleware = append([]Middleware{middleware}, responseMiddleware...)

		if middleware, ok := middleware.(RequestMiddleware); ok {
			response = middleware.ProcessView(request, view)
			// Early return in middleware if it generated a response
			if response != nil {
				break
			}
		}
	}

	if response == nil {
		// process the real request
		response = view(request)
	}

	// Apply response middleware
	for _, middleware := range responseMiddleware {
		if middleware, ok := middleware.(ResponseMiddleware); ok {
			response = middleware.ProcessResponse(request, response)
		}
	}

	// hand back to net/http
	for key, values := range response.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(response.StatusCode)

	defer response.Body.Close()
	io.Copy(w, response.Body)
}
