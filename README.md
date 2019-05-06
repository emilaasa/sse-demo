# Simple Server-Sent-Events demo

This is an exploration of translating gRPC server side streaming to
Server-Sent-Events in the browser.


To try it out:  
`go run main.go` and visit `localhost:8080/debug`

For a look at how the browser handles interupted streams, `Ctrl-C` the running
server and look at the headers in `stdout` and `developer console` in the browser.
Without any extra retry logic Firefox times out and closes the stream after just
a few seconds.

## Goals 
- [X] translate grpc stream to http Server-Sent-Events stream
- [X] create simple index page that uses eventsource to demo reconnect
- [X] inspect http headers 
- [X] inspect gRPC headers
- [X] translate last event id header to grpc metadata
- [ ] display error messages on debug page
- [ ] handle Last-Event-ID on the server side (replay events)
- [ ] set retry header from the service 
- [ ] forever-reconnect mechanism for debug page
- [ ] display Last-Event-ID on debug page
- [ ] make a demo orchestrator where server stops and then resumes
- [ ] how to integrate with grpc-gateway?

