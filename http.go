package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func params(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	for k, v := range query {
		values := strings.Join(v[:], ", ")
		fmt.Fprintf(w, "%s -> %s\n", k, values)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/params", params)
	http.HandleFunc("/headers", headers)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8090"
	}

	filepath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}

	file := path.Base(filepath)

	fmt.Println(file)
	fmt.Printf("Listening on %s...\n", port)

	http.ListenAndServe(":"+port, nil)
}
