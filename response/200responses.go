package response

func WriteOk(w writer, content ...interface{}) {
	write(w, 200, content)
}
