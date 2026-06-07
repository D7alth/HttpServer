package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", hellowolrd)
	http.HandleFunc("/goodbye", goodbye)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func hellowolrd(w http.ResponseWriter, request *http.Request) {
	wc, err := w.Write([]byte("hello world!"))
	if err != nil {
		slog.Error("An error ocurred written message", "err", err)
		return
	}
	fmt.Println("bytes written:", wc)
}

func goodbye(w http.ResponseWriter, request *http.Request) {
	wc, err := w.Write([]byte("goodbye!"))
	if err != nil {
		slog.Error("An error ocurred written message", "err", err)
		return
	}
	fmt.Print("bytes written:", wc)
}
