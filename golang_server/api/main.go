package api

import (
	"fmt"
	"idk/main/api/endpoints"
	"net/http"
)

var Mux *http.ServeMux = http.NewServeMux()

func Init() {
	Mux := http.NewServeMux()
	endpoints.Ping(Mux)
	endpoints.Bucket(Mux)
	err := http.ListenAndServe(":8080", Mux)
	fmt.Println(err)
}
