package main

import (
    "encoding/json"
    "log"
    "code.google.com/p/go.net/websocket"
    "github.com/garyburd/redigo/redis"
    "github.com/nu7hatch/gouuid"
)

// Websocket Message
//
// Used for JSON conversion.
// action = SUBSCRIBE|UNSUBSCRIBE|PUBLISH
// channel = Redis channel
// data = Message to be sent
type WSMessage struct {
    Action string `json:"action"`
    Channel string `json:"channel"`
    Data string `json:"data"`
}


// Websocket Connection
//
// Handles incoming and outcoming websocket data by communicating
// with Redis via its PubSub commands.
type WSConnection struct {
    uuid string
    socket *websocket.Conn
    publish *redis.PubSubConn
    subscribe *redis.PubSubConn
}

// Register at WSServer and connect to Redis.
func (wsc *WSConnection) Initialize() {
    wss.register <- wsc

    uuid, err := uuid.NewV4()
    if err == nil {
        log.Println("Initialize", uuid.String())
        wsc.uuid = uuid.String()
        wsc.publish = wsc.MakeRedisConnection()
        wsc.subscribe = wsc.MakeRedisConnection()
        message := WSMessage {"CONNECT", wsc.uuid, wsc.uuid}
        wsc.SendWebsocket(message)
    }
}

// Unregister from WSServer and disconnect from Redis.
func (wsc *WSConnection) Uninitialize() {
    log.Println("Uninitialize")
    wss.unregister <- wsc

    wsc.publish.Close()
    wsc.subscribe.Close()
}

// Read from Websocket (do nothing)
func (wsc *WSConnection) ReadWebsocket() {
    for {
        var json_data []byte

        // Receive data from Websocket
        err := websocket.Message.Receive(wsc.socket, &json_data)
        if err != nil {
            return
        }
    }
}

// Send to Websocket
func (wsc *WSConnection) SendWebsocket(message WSMessage) {
    json_data, err := json.Marshal(message)
    if err == nil {
        websocket.Message.Send(wsc.socket, string(json_data))
    }
}

// Proxy incoming data from Redis to Websocket.
func (wsc *WSConnection) ProxyRedisSubscribe() {
    for {
        switch reply := wsc.subscribe.Receive().(type) {
        case redis.Message:
            message := WSMessage {"PUBLISH", reply.Channel, string(reply.Data)}
            wsc.SendWebsocket(message)
        case redis.Subscription:
            message := WSMessage {"SUBSCRIBE", reply.Channel, ""}
            wsc.SendWebsocket(message)
        case error:
            return
        }
    }
}

// Establish a connection to Redis via Redigo PubSubConn
func (wsc *WSConnection) MakeRedisConnection() *redis.PubSubConn {
    c, err := redis.Dial("tcp", ":6379")
    if err != nil {
        log.Fatal(err)
    }
    return &redis.PubSubConn{c}
}





func handleWSConnection(socket *websocket.Conn) {
    wsc := &WSConnection {socket: socket}
    defer wsc.Uninitialize()
    wsc.Initialize()
    go wsc.ProxyRedisSubscribe()
    wsc.ReadWebsocket()
}
