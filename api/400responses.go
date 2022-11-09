package api

// WriteNotFoundError writes a 404 response. No other information
// is passed to the client.
func WriteNotFoundError(w writer) {
	writeError(w, 404, "Not found")
}

// WriteBadRequestError writes a 400 response.
//
// errs are expected to correspond to individual
// request parameters
func WriteBadRequestError(w writer, errs ...error) {
	writeError(w, 400, "Bad Request", errs...)
}
