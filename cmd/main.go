package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	uploadHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		s := fmt.Sprintf("%v\n", buf.String())
		defer req.Body.Close()
		fmt.Fprintf(os.Stdout, s)
		io.WriteString(w, s)
	}

	http.HandleFunc("/v1/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":7000", nil))

}
