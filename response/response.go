package response

import (
	"encoding/json"
)

type writer interface {
	WriteHeader(code int)
	Write(bytes []byte) (int, error)
}

func write(w writer, code int, content ...interface{}) {
	var b []byte
	var err error
	if len(content) == 0 {
		b, err = *new([]byte), nil
	} else if len(content) == 1 {
		b, err = json.Marshal(content[0])
	} else {
		b, err = json.Marshal(content)
	}

	if err != nil {
		writeMarshalError(w, err)
		return
	}
	w.WriteHeader(code)
	w.Write(b)
}
