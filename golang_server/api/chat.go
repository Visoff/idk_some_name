package api

import (
	"encoding/json"
	"fmt"
	"idk/main/db"
	"io"
	"net/http"
)

func ApplyChatHandlers() {
	Mux.HandleFunc("/api/user/chats", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "GET":
			user_id := AuthorizationJwt(&w, &r)
			if user_id == nil {
				return
			}
			chats, err := db.Query(`select * from "Chat" inner join "ChatMember" on "ChatMember"."Chat_id" = "Chat"."id" where "ChatMember"."User_id" = '%v'`, fmt.Sprint(user_id))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			chats_json, _ := json.Marshal(chats)
			w.WriteHeader(200)
			w.Write(chats_json)
			return
		}
	})

	Mux.HandleFunc("/api/chat/", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "POST":
			user_id := AuthorizationJwt(&w, &r)
			if user_id == nil {
				return
			}
			body_json, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Provide body"))
				return
			}
			var body map[string]string
			json.Unmarshal(body_json, &body)
			if body == nil {
				w.WriteHeader(400)
				w.Write([]byte("Body must be json"))
				return
			}
			if body["name"] == "" || body["description"] == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Body must have keys: ["name", "description"]`))
				return
			}
			inserted, err := db.Query(`insert into "Chat"("name", "description") values('%v', '%v') returning *`, body["name"], body["description"])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			_, err = db.Query(`insert into "ChatMember"("User_id", "Chat_id") values('%v', '%v')`, fmt.Sprint(user_id), inserted[0]["id"])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			return
		}
	})
}
