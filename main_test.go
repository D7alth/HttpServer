package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHelloWorld(t *testing.T) {
	w := httptest.NewRecorder()
	hellowolrd(w, nil)
	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}
	desiredBodyMessage := []byte("hello world!")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}
