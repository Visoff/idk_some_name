package api

import (
	"encoding/json"
	"fmt"
	"idk/main/db"
	"io"
	"net/http"
)

func ApplyMessageHandlers() {
	Mux.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "POST":
			user_id := AuthorizationJwt(&w, &r)
			if user_id == nil {
				return
			}
			chat_id := r.Header.Get("Chat")
			if chat_id == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Provide chat id in "Chat" header`))
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Provide body"))
				return
			}
			inserted, err := db.Query(`insert into "Message"("content", "contentType", "author", "Chat_id") values('%v', 'text', '%v', '%v') returning *`, string(body), fmt.Sprint(user_id), chat_id)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			_, err = db.Query(`update "Chat" set "last_update"=now() where "id" = '%v'`, chat_id)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			inserted_json, err := json.Marshal(inserted[0])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write(inserted_json)
			return
		}
	})

	Mux.HandleFunc("/api/chat/messages/", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "GET":
			user_id := AuthorizationJwt(&w, &r)
			if user_id == nil {
				return
			}
			chat_id := r.Header.Get("Chat")
			if chat_id == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Provide chat id in "Chat" header`))
				return
			}
			req_time := r.Header.Get("LongPoll")
			var messages []map[string]interface{}
			var err error
			if req_time == "" {
				messages, err = db.Query(`select * from "Message" where "Chat_id" = '%v' limit 20`, chat_id)
			} else {
				messages, err = db.Query(`select * from "Message" where "Chat_id" = '%v' and "created_at" > '%v' limit 20`, chat_id, req_time)
			}
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			for i, message := range messages {
				messages[i]["self"] = message["author"] == fmt.Sprint(user_id)
			}
			messages_json, err := json.Marshal(messages)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			if messages_json == nil {
				w.WriteHeader(200)
				w.Write([]byte("[]"))
				return
			}
			w.WriteHeader(200)
			w.Write(messages_json)
			return
		}
	})
}
