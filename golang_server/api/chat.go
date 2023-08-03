package api

import (
	"encoding/json"
	"fmt"
	"idk/main/db"
	"idk/main/jwt"
	"net/http"
	"strings"
)

func ApplyChatHandlers() {
	Mux.HandleFunc("/api/chat", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if (r.Header.Get("Authorization")) == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Authorization header(bearer)"))
			}
			user_token := strings.Join(strings.Split(r.Header.Get("Authorization"), " ")[1:], " ")
			user_id, err := jwt.Use().Verify(user_token)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			chats, err := db.Query(`select "Chat".* from "Chat" inner join "ChatMember" on "ChatMember"."Chat_id" = "Chat".id and "ChatMember"."User_id" = '%v'`, fmt.Sprint(user_id))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			chats_json, err := json.Marshal(chats)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(chats_json)
			return
		case "POST":
			if (r.Header.Get("Authorization")) == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Authorization header(bearer)"))
			}
			user_token := strings.Join(strings.Split(r.Header.Get("Authorization"), " ")[1:], " ")
			user_id, err := jwt.Use().Verify(user_token)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			var body map[string]string
			decoder := json.NewDecoder(r.Body)
			err = decoder.Decode(&body)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			if body["name"] == "" || body["description"] == "" {
				w.WriteHeader(500)
				w.Write([]byte("Provide name and description"))
				return
			}
			created, err := db.Query(`insert into "Chat"("name") values('%v') returning *`, body["name"])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			_, err = db.Query(`insert into "ChatMember"("Chat_id", "User_id") values('%v', '%v')`, created[0]["id"], fmt.Sprint(user_id))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprint(created[0])))
		}
	})
}
