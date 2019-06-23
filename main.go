package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func main() {
	port := flag.Int("p", 80, "Port")
	flag.Parse()

	log.Println(fmt.Sprintf("Listening on port %v...", *port))

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
	//http.ListenAndServe(":80", nil)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", *port), nil))
}
