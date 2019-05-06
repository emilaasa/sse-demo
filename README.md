# Simple Server-Sent-Events demo

`go run main.go` and visit `localhost:8080/debug`


## Goals 
- [X] translate grpc stream to http Server-Sent-Events stream
- [X] create simple index page that uses eventsource to demo reconnect
- [X] inspect http headers 
- [X] inspect gRPC headers
- [ ] translate last event id header to grpc metadata
- [ ] update javascript page with forever-reconnect mechanism
- [ ] set retry header from the server somehow
- [ ] how to integrate with grpc-gateway

