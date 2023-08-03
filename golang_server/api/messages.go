package api

import (
	"encoding/json"
	"fmt"
	"idk/main/db"
	"idk/main/jwt"
	"io"
	"net/http"
	"strings"
)

func ApplyMessageHandlers() {
	Mux.HandleFunc("/api/message/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if r.Header.Get("Chat") == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Provide Headers: "Chat"`))
				return
			}
			messages, err := db.Query(`select * from "Message" where "Chat_id"='%v'`, r.Header.Get("Chat"))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			messages_json, err := json.Marshal(messages)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(messages_json)
			return
		case "POST":
			if (r.Header.Get("Authorization")) == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Authorization header(bearer)"))
				return
			}
			user_token := strings.Join(strings.Split(r.Header.Get("Authorization"), " ")[1:], " ")
			user_id, err := jwt.Use().Verify(user_token)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			if (r.Header.Get("Chat")) == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Provide Headers: "Chat"`))
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			created, err := db.Query(`insert into "Message"("author", "content", "contentType", "Chat_id") values('%v', '%v', 'text', '%v') returning *`, fmt.Sprint(user_id), string(body), r.Header.Get("Chat"))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			created_json, err := json.Marshal(created[0])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(created_json)
		}
	})
	Mux.HandleFunc("/api/message/poll/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if r.Header.Get("LastDate") == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Provide Headers: "LastDate"`))
				return
			}
			messages, err := db.Query(`select * from "Message" where "last_update" > '%v'`, r.Header.Get("LastDate"))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			if len(messages) == 0 {
				w.WriteHeader(204)
				w.Write([]byte(``))
				return
			}
			messages_json, err := json.Marshal(messages)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write(messages_json)
			return
		}
		w.Write([]byte("incorrect method"))
	})
}
