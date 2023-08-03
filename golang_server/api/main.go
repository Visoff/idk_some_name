package api

import (
	"fmt"
	"net/http"
)

var Mux *http.ServeMux = http.NewServeMux()

func Init() {
	ApplyMessageHandlers()
	ApplyUserHandlers()
	ApplyChatHandlers()
	err := http.ListenAndServe(":8080", Mux)
	fmt.Println(err)
}
