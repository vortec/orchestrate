package main

import (
	"log"
	"net/http"
	"text/template"
	"code.google.com/p/go.net/websocket"
)

var homeTemplate = template.Must(template.ParseFiles("html/index.html"))

// Return HTML page
func handleWeb(conn http.ResponseWriter, req *http.Request) {
	homeTemplate.Execute(conn, req.Host)
}

// Open HTTP server, return HTML page on / and create new WSConnection when
// user establishes a WebSocket connection at /ws.
func serveWeb() {
	http.HandleFunc("/", handleWeb)
	http.Handle("/ws", websocket.Handler(handleWSConnection))
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
