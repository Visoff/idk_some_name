package endpoints

import (
	Http "idk/main/api/lib"
	api_middleware "idk/main/api/middleware"
	"idk/main/jwt"
	"io"
	"net/http"
)

func User(mux *http.ServeMux) {
	cluster := Http.Cluster(mux, "/user")
	cluster.Rest("/auth").Use(api_middleware.AllowCors).Post(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		user, err := io.ReadAll(r.Body)
		if err != nil {
			res.Status(400).Send("Provide body")
			return
		}
		token, err := jwt.Use().Encode(string(user))
		if err != nil {
			res.Status(500).Send("Could not generate jwt token")
			return
		}
		res.Status(200).Send(token)
	}).Apply()
	cluster.Rest("/").Use(api_middleware.AllowCors).Use(api_middleware.Auth).Get(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		res.Status(200).Send(r.URL.Query().Get("Authorization"))
	}).Apply()
}
