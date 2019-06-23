package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

type MsgHandler struct {
	Msg string
}

func (mh *MsgHandler) greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", mh.Msg)
}

func main() {
	port := flag.Int("p", 80, "Port")
	msg := flag.String("m", "Message", "-m \"Welcome!\"")
	flag.Parse()

	mh := &MsgHandler{Msg: *msg}

	log.Println(fmt.Sprintf("Listening on port %v...", *port))

	http.HandleFunc("/", hello)
	http.HandleFunc("/greet", mh.greet)
	//http.ListenAndServe(":80", nil)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", *port), nil))
}
