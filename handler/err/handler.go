package err

import (
	"log"
	"net/http"
)

// Use as a wrapper around the handler functions.
type ErrorHandler func(http.ResponseWriter, *http.Request) error

// ErrorHandler implements http.Handler interface.
func (fn ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call handler function
	err := fn(w, r)
	if err == nil {
		return
	}

	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err)

	// Check if it is a ClientError.
	clientError, ok := err.(ClientError)
	if !ok {
		// If the error is not ClientError, assume that it is ServerError.
		// return 500 Internal Server Error.
		w.WriteHeader(500)
		return
	}

	// Try to get response body of ClientError.
	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}

	// Get http status code and headers.
	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}
