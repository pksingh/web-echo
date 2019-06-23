package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

type RequestBody struct {
	String string
	// Add other items
}
// Request is the same as http.Request minus the bits that break json.Marshall
type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	ProtoMajor       int    // 1
	ProtoMinor       int    // 0
	Header           http.Header
	Body             RequestBody
	ContentLength    int64
	TransferEncoding []string
	Host             string
	//Form url.Values
	//PostForm url.Values
	//MultipartForm *multipart.Form
	Trailer    http.Header
	RemoteAddr string
	RequestURI string
	//TLS *tls.ConnectionState
}

const megabytes = 1048576

func echo(w http.ResponseWriter, r *http.Request) {
	e := Request{}
	e.Method = r.Method
	e.URL = r.URL
	e.Proto = r.Proto
	e.ProtoMajor = r.ProtoMajor
	e.ProtoMinor = r.ProtoMinor
	e.Header = r.Header
	e.Body = RequestBody{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1*megabytes))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	e.Body.String = string(body)
	e.ContentLength = r.ContentLength
	e.TransferEncoding = r.TransferEncoding
	e.Host = r.Host
	e.Trailer = r.Trailer
	e.RemoteAddr = r.RemoteAddr
	e.RequestURI = r.RequestURI

	b, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(b))
}

func main() {
	port := flag.Int("p", 80, "Port")
	msg := flag.String("m", "Message", "-m \"Welcome!\"")
	flag.Parse()

	mh := &MsgHandler{Msg: *msg}

	log.Println(fmt.Sprintf("Listening on port %v...", *port))

	http.HandleFunc("/", hello)
	http.HandleFunc("/greet", mh.greet)
	http.HandleFunc("/echo", echo)
	//http.ListenAndServe(":80", nil)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", *port), nil))
}
