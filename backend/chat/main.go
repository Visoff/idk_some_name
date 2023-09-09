package main

import (
	"fmt"
	"net/http"
	"strings"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
	api_middleware "github.com/Visoff/idk_some_name/golang_library/api/middleware"
	"github.com/Visoff/idk_some_name/golang_library/db"
	"github.com/Visoff/idk_some_name/golang_library/env"
	"github.com/gorilla/websocket"
)

func main() {
	err := db.Connect(db.UrlFromEnv(env.Env))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rooms := make(map[string]map[string]*websocket.Conn)

	mux := http.NewServeMux()

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	Api_lib.Rest(mux, "/").Use(api_middleware.AllowCors).Get(func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		room_id := path[len(path)-1]
		if _, ok := rooms[room_id]; !ok {
			rooms[room_id] = make(map[string]*websocket.Conn)
		}
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			ws.Close()
			return
		}
		var auth_message map[string]string
		err = ws.ReadJSON(&auth_message)
		if err != nil {
			ws.Close()
			return
		}
		auth := auth_message["auth"]
		if auth == "" {
			ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"Unauthorized"}`))
			ws.Close()
			return
		}
		rooms[room_id][auth] = ws

		messages, err := db.Query("select * from \"Message\" where \"Chat_id\" = '%v'", room_id)
		if err != nil {
			ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"`+err.Error()+`"}`))
		}
		for _, message := range messages {
			message["type"] = "message"
			ws.WriteJSON(message)
		}

		defer func() {
			ws.Close()
			delete(rooms[room_id], auth)
			if len(rooms[room_id]) == 0 {
				delete(rooms, room_id)
			}
		}()

		for {
			var message map[string]interface{}
			err := ws.ReadJSON(&message)
			if err != nil {
				return
			}
			message_type, ok := message["type"].(string)
			if !ok {
				ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"Type must be provided"}`))
			}
			switch message_type {
			case "list":
				ws.WriteJSON(rooms[room_id])
			case "message":
				var content string
				if content, ok = message["content"].(string); !ok {
					ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"Content must be provided for type message"}`))
					continue
				}
				message["from"] = auth
				for user, ws := range rooms[room_id] {
					if auth == user {
						continue
					}
					new_message := make(map[string]string)
					new_message["type"] = "message"
					new_message["content"] = content
					new_message["contentType"] = "text"
					new_message["author"] = auth
					ws.WriteJSON(new_message)
				}
				_, err := db.Query("insert into \"Message\"(\"content\", \"contentType\", \"author\", \"Chat_id\") values('%v', '%v', '%v', '%v')", content, "text", auth, room_id)
				if err != nil {
					ws.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"`+err.Error()+`"}`))
					continue
				}
			}
		}
	}).Apply()

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
