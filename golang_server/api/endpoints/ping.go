package endpoints

import (
	Http "idk/main/api/lib"
	api_middleware "idk/main/api/middleware"
	"net/http"
)

func Ping(mux *http.ServeMux) {
	cluster := Http.Cluster(mux, "/")
	cluster.Rest("/ping").Use(api_middleware.AllowCors).Get(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		res.Status(200).Send("pong")
	}).Apply()
}
