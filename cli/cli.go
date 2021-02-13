package main

import (
	"bufio"
	"net/http"
	"os"
)

func main() {
	http.Post(
		"http://172.0.0.1:7000/v1/upload",
		"Content-Type: text/html; charset=utf-8",
		bufio.NewReader(os.Stdout),
	)
}
