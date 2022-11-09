package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter struct {
	code  int
	bytes []byte
}

func newMockWriter() *mockWriter {
	return &mockWriter{
		code:  0,
		bytes: []byte{},
	}
}

func (m *mockWriter) Write(bytes []byte) (int, error) {
	m.bytes = bytes
	return 0, nil
}

func (m *mockWriter) WriteHeader(code int) {
	m.code = code
}

func TestWrite(t *testing.T) {
	mock := newMockWriter()
	code := 100
	content := map[string]string{"test": "value"}
	bytes, _ := json.Marshal(content)

	write(mock, code, content)
	assert.Equal(t, code, mock.code)
	assert.Equal(t, bytes, mock.bytes)
}

func TestWriteArray(t *testing.T) {
	mock := newMockWriter()
	code := 100
	content := map[string]string{"test": "value"}
	bytes, _ := json.Marshal([]map[string]string{content, content})

	write(mock, code, content, content)
	assert.Equal(t, code, mock.code)
	assert.Equal(t, bytes, mock.bytes)
}

func TestWriteEmpty(t *testing.T) {
	mock := newMockWriter()
	code := 100

	write(mock, code)
	assert.Equal(t, code, mock.code)
	assert.Equal(t, "", string(mock.bytes))
}

func TestWriteWithMarshalError(t *testing.T) {
	mock := newMockWriter()
	code := 100
	content := make(chan int)
	write(mock, code, content)

	var res ErrorResponse
	json.Unmarshal(mock.bytes, &res)

	assert.Equal(t, 500, mock.code)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, "Internal Server Error", res.Message)
}

func TestWriteError(t *testing.T) {
	mock := newMockWriter()
	code := 100
	message := "test message"
	err1 := errors.New("error 1")
	err2 := fmt.Errorf("error 2")
	writeError(mock, code, message, err1, err2)

	var res ErrorResponse
	json.Unmarshal(mock.bytes, &res)
	assert.Equal(t, code, mock.code)
	assert.Equal(t, code, res.Code)
	assert.Equal(t, message, res.Message)
	assert.Equal(t, len(res.Errors), 2)
	assert.Equal(t, err1.Error(), res.Errors[0])
	assert.Equal(t, err2.Error(), res.Errors[1])
}
