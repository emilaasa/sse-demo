# Simple Server-Sent-Events demo

`go run main.go` and visit `localhost:8080/debug`


## Goals 
[x] translate grpc stream to http Server-Sent-Events stream
[x] create simple index page that uses eventsource to demo reconnect
[ ] inspect http headers 
[ ] inspect gRPC headers
[ ] translate last event id header to grpc metadata
[ ] update javascript page with forever-reconnect mechanism
[ ] set retry header from the server somehow

[ ] discuss how to integrate

