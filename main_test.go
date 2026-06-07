package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHwlloWorld(t *testing.T) {
	w := httptest.NewRecorder()
	hellowolrd(w, nil)
	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body)
	}
}
