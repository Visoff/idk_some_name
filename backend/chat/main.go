package main

import (
	"fmt"
	"net/http"
	"strings"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
	api_middleware "github.com/Visoff/idk_some_name/golang_library/api/middleware"
	"github.com/Visoff/idk_some_name/golang_library/db"
	"github.com/Visoff/idk_some_name/golang_library/env"
	"github.com/Visoff/idk_some_name/golang_library/jwt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Rooms struct {
	this map[string]map[string]*websocket.Conn
}

func (rooms *Rooms) Join(room_id string, auth string, ws *websocket.Conn) roomWithUser {
	user_id := uuid.New().String()
	if room, ok := rooms.this[room_id]; ok {
		room[user_id] = ws
	} else {
		new_room := make(map[string]*websocket.Conn)
		new_room[user_id] = ws
		rooms.this[room_id] = new_room
	}
	return roomWithUser{room_id: room_id, user_id: user_id, auth: auth, rooms: *rooms}
}

type roomWithUser struct {
	room_id string
	user_id string
	auth    string
	rooms   Rooms
}

func (r *roomWithUser) Leave() {
	delete(r.rooms.this[r.room_id], r.user_id)
	if len(r.rooms.this[r.room_id]) == 0 {
		delete(r.rooms.this, r.room_id)
	}
}

func (r *roomWithUser) emit(f func(ws *websocket.Conn) error) error {
	for user_id, ws := range r.rooms.this[r.room_id] {
		if r.user_id == user_id {
			continue
		}
		err := f(ws)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	err := db.Connect(db.UrlFromEnv(env.Env))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rooms := Rooms{make(map[string]map[string]*websocket.Conn)}

	mux := http.NewServeMux()

	Api_lib.Rest(mux, "/").Use(api_middleware.AllowCors).Get(func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Unable to create WebSocket connection"))
			return
		}
		path := strings.Split(r.URL.Path, "/")
		room_id := path[len(path)-1]

		room := rooms.Join(room_id, "", ws)

		defer room.Leave()
		var message map[string]string
		for {
			err := ws.ReadJSON(&message)
			if err != nil {
				break
			}
			if message["type"] != "auth" && room.auth == "" {
				ws.WriteMessage(websocket.TextMessage, []byte("Please authorize"))
				continue
			}
			switch message["type"] {
			case "message":
				new_message := make(map[string]string)
				new_message["type"] = "message"
				new_message["from"] = room.auth
				new_message["content"] = message["content"]
				err := room.emit(func(ws *websocket.Conn) error {
					return ws.WriteJSON(new_message)
				})
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
			case "auth":
				auth, err := jwt.Use().Decode(message["token"])
				if err != nil {
					ws.WriteMessage(websocket.TextMessage, []byte("Incorrect token"))
					continue
				}
				room.auth = auth.(string)
				ws.WriteMessage(websocket.TextMessage, []byte("Success"))
			case "reg":
				auth, err := jwt.Use().Encode(message["token"])
				if err != nil {
					ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
					continue
				}
				ws.WriteMessage(websocket.TextMessage, []byte(auth))
			}
		}
	}).Apply()
	http.ListenAndServe(":8080", mux)
}
