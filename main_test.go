package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	w := httptest.NewRecorder()

	HandleRoot(w, nil)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("Welcome to our webpage!\n")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandleGoodbye(t *testing.T) {
	w := httptest.NewRecorder()

	HandleGoodbye(w, nil)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("goodbye!\n")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandleParametrized(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?user=alberth", nil)

	w := httptest.NewRecorder()

	HanldeParametrized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("Hello, alberth!")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandlerParametrizedWithNoParameters(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	w := httptest.NewRecorder()

	HanldeParametrized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("Hello, user!")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandlerParametrizedWithNoValidParameters(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?foo=bar", nil)

	w := httptest.NewRecorder()

	HanldeParametrized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("Hello, user!")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandleUserHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/responses/alberth/hello", nil)
	req.SetPathValue("user", "alberth")

	w := httptest.NewRecorder()

	HandleUserHello(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("Hello, alberth!")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandleUserHelloHeaders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/responses/hello", nil)
	req.Header.Set("user", "alberth")

	w := httptest.NewRecorder()

	HandleUserHelloHeaders(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("Hello, alberth!")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandleUserHelloHeadersWithNoParameters(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/responses/hello", nil)

	w := httptest.NewRecorder()

	HandleUserHelloHeaders(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("bad request\n")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}

func TestHandleUserHelloHeadersWithNoValidParameters(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/responses/hello", nil)
	req.Header.Set("someKey", "key")

	w := httptest.NewRecorder()

	HandleUserHelloHeaders(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("Bad respose code, excpeted %v but was %v\n Body: %s",
			desiredCode, w.Code, w.Body.String())
	}

	desiredBodyMessage := []byte("bad request\n")
	if !bytes.Equal(desiredBodyMessage, w.Body.Bytes()) {
		t.Errorf("The message body does't seem correct, body: %s", w.Body.String())
	}
}
