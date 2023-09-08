package main

import (
	"fmt"
	"net/http"
	"strings"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
	api_middleware "github.com/Visoff/idk_some_name/golang_library/api/middleware"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	mux := http.NewServeMux()

	rooms := make(map[string]map[string]*websocket.Conn)

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	Api_lib.Rest(mux, "/").Use(api_middleware.AllowCors).Get(func(w http.ResponseWriter, r *http.Request) {
		auth := uuid.New().String()
		path := strings.Split(r.URL.Path, "/")
		room := path[len(path)-1]
		ws, _ := upgrader.Upgrade(w, r, nil)
		ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"auth","auth":"`+auth+`"}`))
		if _, ok := rooms[room]; !ok {
			rooms[room] = make(map[string]*websocket.Conn)
		}
		rooms[room][auth] = ws
		var history []string
		for user, _ := range rooms[room] {
			history = append(history, user)
		}
		message := make(map[string]interface{})
		message["type"] = "history"
		message["users"] = history
		ws.WriteJSON(message)

		defer func() {
			ws.Close()
			delete(rooms[room], auth)
			for _, ws := range rooms[room] {
				ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"disconnect","from":"`+auth+`"}`))
			}
		}()

		for {
			var message map[string]interface{}
			err := ws.ReadJSON(&message)
			if err != nil {
				break
			}
			message["from"] = auth
			to := fmt.Sprint(message["to"])
			if to != "" {
				rooms[room][to].WriteJSON(message)
				continue
			}
			for user, ws := range rooms[room] {
				if auth == user {
					continue
				}
				err := ws.WriteJSON(message)
				if err != nil {
					delete(rooms[room], user)
				}
			}
		}
	}).Apply()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
