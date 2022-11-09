package api

func WriteOk(w writer, content ...interface{}) {
	if len(content) == 1 {
		write(w, 200, content[0])
		return
	}
	write(w, 200, content)
}
