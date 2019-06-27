# Web Server on GoLang

## Build web-echo Server

  `$ go build`

## Running web-echo

  `$ ./web-echo` 
  `web-echo is running on 80 port.`

## Running on different Port; use -p port

  `$ ./web-echo -p 8000`                     
  `2019/06/23 23:25:28 Listening on port 8000...`

if port number not provided; default port 80 will be used

## Serving custome message; use -m "msg"

  `$ ./web-echo -p 8000 -m "my greeting msg"`

custome message will be served on '/greet' endpoint

In this case;

  `$ curl http://localhost:8000/greet`
  `my greeting msg   `                                                                                                        

If no -m provided; default message is be "Message"

  `$ curl http://localhost:8000/greet`
  `Message   `


## Echo Server : endpoint /echo

  `$ ./web-echo -p 8000 `

In this case; 

  `$ curl http://localhost:8000/echo`

  `{"Method":"GET","URL":{"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/echo","RawPath":"","ForceQuery":false,"RawQuery":"","Fragment":""},"Proto":"HTTP/1.1","ProtoMajor":1,"ProtoMinor":1,"Header":{"Accept":["*/*"],"User-Agent":["curl/7.54.0"]},"Body":{"String":""},"ContentLength":0,"TransferEncoding":null,"Host":"localhost:8000","Trailer":null,"RemoteAddr":"[::1]:53638","RequestURI":"/echo"}`                                                                                                      

If no -m provided; default message is be "Message"

  `$ curl http://localhost:8000/greet`
  `Message  `
  
## Docker Image building

  Before Building Docker image make sure two things:
  - Docker Application must be available
  - Build the OS compatiable binary
  
  Then we use the `Dockerfile` to build the Docker image.

### Build Binary for OS compatible 

  `$ GOOS=linux go build . `
  
### Build Docker Image

  `$ docker build . -f Dockerfile -t web-echo`
  
### Build Binary and Docker Image as one liner cmd

  `$ GOOS=linux go build . && docker build . -f Dockerfile -t web-echo`
  
