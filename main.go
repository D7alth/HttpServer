package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", hellowolrd)
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
