package main

import (
	"fmt"
	"net/http"

	handel "ascii-art-web/Handler"
)

type Data struct {
	Result string
}

func main() {
	http.HandleFunc("/", handel.Home)
	http.HandleFunc("/ascii-art", handel.Ascii_art)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
