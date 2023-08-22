package api

import (
	"fmt"
	"idk/main/api/endpoints"
	Http "idk/main/api/lib"
	api_middleware "idk/main/api/middleware"
	"net/http"
)

var Mux *http.ServeMux = http.NewServeMux()

func Init() {
	Mux := http.NewServeMux()
	Http.Rest(Mux, "/").Use(api_middleware.AllowCors).Apply()
	endpoints.Ping(Mux)
	endpoints.Bucket(Mux)
	endpoints.User(Mux)
	err := http.ListenAndServe(":8080", Mux)
	fmt.Println(err)
}
