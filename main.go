package main

import (
	"bytes"
	"log"
	"log/slog"
	"net/http"
)

var writtenErrorMessage string = "An error ocurred written message"

func main() {
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/goodbye", HandleGoodbye)
	http.HandleFunc("/hello", HanldeParametrized)
	http.HandleFunc("/responses/{user}/hello", HandleUserHello)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	_, err := w.Write([]byte("Welcome to our webpage!\n"))
	if err != nil {
		slog.Error(writtenErrorMessage, "err", err)
		return
	}
}

func HandleGoodbye(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("goodbye!\n"))
	if err != nil {
		slog.Error(writtenErrorMessage, "err", err)
		return
	}
}

func HanldeParametrized(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	userList, ok := params["user"]
	if !ok {
		userList = append(userList, "user")
	}
	var out bytes.Buffer
	out.WriteString("Hello, ")
	out.WriteString(userList[0])
	out.WriteString("!")
	_, err := w.Write(out.Bytes())
	if err != nil {
		slog.Error(writtenErrorMessage, "err", err)
		return
	}
}

func HandleUserHello(w http.ResponseWriter, req *http.Request) {
	userValuePath := req.PathValue("user")
	var out bytes.Buffer
	out.WriteString("Hello, ")
	out.WriteString(userValuePath)
	out.WriteString("!")
	_, err := w.Write(out.Bytes())
	if err != nil {
		slog.Error(writtenErrorMessage, "err", err)
		return
	}
}
