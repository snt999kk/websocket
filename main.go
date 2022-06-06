package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/abc", echo)
	http.HandleFunc("/main.html", serveHTML)
	http.HandleFunc("/main.js", serveJs)
	http.ListenAndServe(":9191", nil)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main.html")
}

func serveJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main.js")
}

func echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)

	for {
		mt, message, err := connection.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			fmt.Println("error:", err)
			fmt.Println("type:", mt)
			break
		}
		response := string(message) + "!!!"
    	connection.WriteMessage(websocket.TextMessage, []byte(response))
	}
}