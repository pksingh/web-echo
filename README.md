# Web Server on GoLang

## Build web-echo Server

  $ go build

## Running web-echo

  $ ./web-echo 
  web-echo is running on 80 port.

## Running on different Port; use -p port
  $ ./web-echo -p 8000                     
  2019/06/23 23:25:28 Listening on port 8000...

if port number not provided; default port 80 will be used

## Serving custome message; use -m "msg"
  $ ./web-echo -p 8000 -m "my greeting msg"

custome message will be served on '/greet' endpoint

In this case;

  $ curl http://localhost:8000/greet
  my greeting msg                                                                                                           

If no -m provided; default message is be "Message"

  $ curl http://localhost:8000/greet
  Message   
