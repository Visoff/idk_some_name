package api

import (
	"encoding/json"
	"fmt"
	"idk/main/db"
	"idk/main/jwt"
	"net/http"
	"strings"
)

func ApplyUserHandlers() {
	Mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			if r.Header.Get("username") == "" || r.Header.Get("password") == "" {
				w.WriteHeader(400)
				w.Write([]byte(`Provide Headers: "username" and "password"`))
				return
			}
			inserted, err := db.Query(`select "id" from "User" where "username" = '%v'`, r.Header.Get("username"))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			toSend, err := jwt.Use().Encode(inserted[0]["id"])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(toSend))
			return
		case "GET":
			if r.Header.Get("Authorization") == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Authorization header(bearer)"))
				return
			}
			token := strings.Join(strings.Split(r.Header.Get("Authorization"), " ")[1:], " ")
			user_id, err := jwt.Use().Verify(token)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			user, err := db.Query(`select * from "User" where "id" = '%v'`, fmt.Sprint(user_id))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			user_json, err := json.Marshal(user[0])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprint(err)))
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(user_json)
			return
		}
	})
}
