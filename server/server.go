package server

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

func Start(port *string, specifications *[]RestRequestResponseSpec) {

	mux := http.NewServeMux()
	for hPath, hFunc := range makeHandlers(specifications) {
		mux.Handle(hPath, hFunc)
	}

	err := http.ListenAndServe(":"+*port, mux)
	if err != nil {
		log.Println("Server error", err)
		os.Exit(100)
	}

}

type PathToHandlerMap = map[string]http.HandlerFunc

func makeHandlers(specifications *[]RestRequestResponseSpec) PathToHandlerMap {
	allHandelrs := make(PathToHandlerMap)
	for _, spec := range *specifications {
		closureSpec := spec
		allHandelrs[spec.RequestSpec.Path] = func(writer http.ResponseWriter, request *http.Request) {

			requestId, _ := uuid.NewUUID()
			logRequest(&requestId, request)

			response := closureSpec.ResponseSpec.Body.Value
			responseCode := closureSpec.ResponseSpec.Code
			responseHeaders := closureSpec.ResponseSpec.Headers

			for _, header := range responseHeaders {
				writer.Header().Set(header.Key, header.Value)
			}

			writer.WriteHeader(responseCode)

			_, _ = fmt.Fprintln(writer, response)
			logResponse(&requestId, &response, &responseCode, &responseHeaders)
		}

		log.Println("Handler definition added", spec)
	}

	return allHandelrs
}
