package response

func WriteServerError(w writer, err ...error) {
	writeError(w, 500, "Internal Server Error", err...)
}
