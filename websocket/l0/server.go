// Copyright 2024 tabuyos. All rights reserved.
//
// @author tabuyos
// @since 2024/08/26
// @description description
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var log = slog.Default()
var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info("upgrade:", "error", err.Error())
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Error("read:", "error", err.Error())
			break
		}
		log.Info(fmt.Sprintf("recv: %s", message))
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Error("write:", "error", err.Error())
			break
		}
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Error(err.Error())
	}
}