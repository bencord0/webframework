package webframework

import (
    "bytes"
    "encoding/json"
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

func JSONResponse(data interface{}, status int, headers http.Header) *http.Response {
    out, err := json.Marshal(data)
    if err != nil {
        return &http.Response{
            StatusCode: http.StatusInternalServerError,
        }
    }

    return &http.Response{
        Body: ioutil.NopCloser(bytes.NewBuffer(out)),
        StatusCode: status,
        Header: headers,
    }
}
