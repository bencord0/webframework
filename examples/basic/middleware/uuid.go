package middleware

import (
    "net/http"
    "github.com/google/uuid"
    . "github.com/bencord0/webframework"
)

// UUID on request | Request middleware
// -----------------------------------------------------------------------------------------------
type UUIDify struct {}

func (m UUIDify) ProcessView(request *http.Request, view ViewFunc) *http.Response{
    newUUID := uuid.New()

    request.Header.Add("UUID", newUUID.String())

    // Pass through to the next middleware
    return nil
}
