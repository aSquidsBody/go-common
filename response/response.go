package response

import "encoding/json"

type writer interface {
	WriteHeader(code int)
	Write(bytes []byte) (int, error)
}

func write(w writer, code int, content interface{}) {
	b, err := json.Marshal(content)
	if err != nil {
		writeMarshalError(w, err)
		return
	}
	w.WriteHeader(code)
	w.Write(b)
}
