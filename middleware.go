package webapi

import (
	"net/http"
)

// Middleware represents a middleware.
type Middleware interface {
	// Processing called before action handler executed.
	OnRequest(*http.Request, http.ResponseWriter) http.ResponseWriter

	// Processed called after action handler executed.
	OnResponse(*http.Request, http.ResponseWriter) http.ResponseWriter
}
