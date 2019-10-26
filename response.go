package webframework

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func TextResponse(content string, status int, headers http.Header) *http.Response {
	return &http.Response{
		Body:       ioutil.NopCloser(strings.NewReader(content)),
		StatusCode: status,
		Header:     headers,
	}
}
