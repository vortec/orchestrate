Orchestrate
===========

Simple Redis PubSub to WebSocket proxy written in Go.

This is my first attempt to write anything in Go, so any feedback is greatly appreciated. Thanks!


Dependencies
------------

Orchestrate needs Redigo and the go.net WebSocket implementation.

    go get code.google.com/p/go.net/websocket
	go get github.com/garyburd/redigo/redis

Also, make sure a Redis instance is up and running:

    redis-server --loglevel verbose


Run
---

After cloning the code, just type

    go run *.go

in the orchestrate directory. Then you can access http://localhost:9000/ to try it out.
