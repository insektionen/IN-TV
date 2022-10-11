package api

import "net/http"

// NotFoundHandler responds with a 'Not found' response
var NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, NotFoundError())
})

// MethodNotAllowedHandler responds with a 'Method not allowed' response
var MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, MethodNotAllowedError())
})
