package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type UserData struct {
	Name string
}

var writtenErrorMessage string = "An error ocurred written message"

func main() {
	http.HandleFunc("/{$}", HandleRoot)
	http.HandleFunc("/goodbye", HandleGoodbye)
	http.HandleFunc("/hello", HanldeParametrized)
	http.HandleFunc("/responses/{user}/hello", HandleUserHello)
	http.HandleFunc("/responses/hello", HandleUserHelloHeaders)
	http.HandleFunc("/responses/hello/json", HandleUserHelloHeaders)
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
	writeHelloUser(userList[0], w, req)
}

func HandleUserHello(w http.ResponseWriter, req *http.Request) {
	user := req.PathValue("user")
	writeHelloUser(user, w, req)
}

func HandleUserHelloHeaders(w http.ResponseWriter, req *http.Request) {
	user := req.Header.Get("user")
	if user == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	writeHelloUser(user, w, req)
}

func HandleUserHelloJson(w http.ResponseWriter, req *http.Request) {
	byteData, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "error to reading body", http.StatusBadRequest)
		return
	}
	var userData UserData
	err = json.Unmarshal(byteData, &userData)
	if err != nil {
		http.Error(w, "error unmarshiling request body", http.StatusBadRequest)
		return
	}
	writeHelloUser(userData.Name, w, req)
}

func writeHelloUser(username string, w http.ResponseWriter, _ *http.Request) {
	var out bytes.Buffer
	out.WriteString("Hello, ")
	out.WriteString(username)
	out.WriteString("!")
	_, err := w.Write(out.Bytes())
	if err != nil {
		slog.Error(writtenErrorMessage, "err", err)
		return
	}
}
