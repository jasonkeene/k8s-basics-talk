package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/crash", crashHandler)

	addr := os.Getenv("ADDR")
	log.Print("listening on: ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func helloHandler(rw http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(rw, "hello from: "+os.Getenv("HOSTNAME"))
}

func crashHandler(http.ResponseWriter, *http.Request) {
	log.Fatal("crashing")
}
