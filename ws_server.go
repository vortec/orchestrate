package main

// WebSocket Server
// 
// Acts as a hub for WSConnections after you called Run().
// Contains a map of active connections and provides two channels,
// one for registering and one for unregistering WSConnections. 
type WSServer struct {
	connections map[*WSConnection] bool
	register chan *WSConnection
	unregister chan *WSConnection
}

// Start waiting for register or unregister commands
func (wss *WSServer) Run() {
	for {
		select {
		case wsc := <- wss.register:
			wss.connections[wsc] = true
		case wsc := <- wss.unregister:
			delete(wss.connections, wsc)
		}
	}
}

// Fake singleton
var wss = WSServer {
	connections: make(map[*WSConnection] bool),
	register: make(chan *WSConnection),
	unregister: make(chan *WSConnection),
}


// Wrapper function for wss.Run
func runWSServer() {
	wss.Run()
}
