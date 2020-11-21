package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// I'm ignoring I'm not checking the origin sync this is a echo server
var upgrader = websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}

	}
}

func main() {
	http.HandleFunc("/", echo)

	log.Fatal(http.ListenAndServe(":9898", nil))
}
