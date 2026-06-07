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

func hellowolrd(wirtter http.ResponseWriter, request *http.Request) {
	out, err := wirtter.Write([]byte("hello world!"))
	if err != nil {
		slog.Error("An error ocurred written message", "err", err)
		return
	}
	fmt.Println("bytes written:", out)
}
