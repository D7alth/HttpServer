package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/goodbye", HandleGoodbye)
	http.HandleFunc("/hello", HanldeParametrized)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HandleRoot(w http.ResponseWriter, request *http.Request) {
	_, err := w.Write([]byte("Welcome to our webpage!\n"))
	if err != nil {
		slog.Error("An error ocurred written message", "err", err)
		return
	}
}

func HandleGoodbye(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("goodbye!\n"))
	if err != nil {
		slog.Error("An error ocurred written message", "err", err)
		return
	}
}

func HanldeParametrized(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "not implmented yet\n", http.StatusNotImplemented)
}
