package server

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

func logRequest(uuid *uuid.UUID, request *http.Request) {
	log.Printf("> Request %s from remote address %s, %s %s%s headers %s; content length %d; reqest body %s; ",
		*uuid,
		request.RemoteAddr,
		request.Method,
		request.Host,
		request.URL,
		request.Header,
		request.ContentLength,
		request.Body)
}

func logResponse(uuid *uuid.UUID, response *string, code *int, headers *[]StringKVSpec) {
	log.Printf("< Response to request=%s code=%d headers=%v\n%s", *uuid, *code, *headers, *response)
}
