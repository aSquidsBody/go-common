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

	b, err = arrayToBytes(content)
	if err != nil {
		writeMarshalError(w, err)
		return
	}

	w.WriteHeader(code)
	w.Write(b)
}

func arrayToBytes(content []interface{}) (b []byte, err error) {
	switch len(content) {
	case 0:
		b, err = nil, nil
	case 1:
		b, err = json.Marshal(content[0])
	default:
		b, err = json.Marshal(content)
	}

	return
}
