package api

import (
	"idk/main/db"
	"idk/main/jwt"
	"io"
	"net/http"
)

func ApplyUserHandlers() {
	Mux.HandleFunc("/api/user/sign", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Body = user access token"))
				return
			}
			user, err := db.Query(`select * from "User" where "clerk_id" = '%v'`, string(body))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			if len(user) == 0 {
				user, err = db.Query(`insert into "User"("clerk_id", "username") values('%v', '') returning *`, string(body))
			}
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			token, err := jwt.Use().Encode(user[0]["id"])
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(token))
			return
		}
	})
}
