package api

import (
	"fmt"
	"net/http"
)

var Mux *http.ServeMux = http.NewServeMux()

func Init() {
	Mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	})
	ApplyMessageHandlers()
	ApplyUserHandlers()
	ApplyChatHandlers()
	ApplyConference()
	err := http.ListenAndServe(":8080", Mux)
	fmt.Println(err)
}
