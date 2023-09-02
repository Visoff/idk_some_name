package api

import (
	Api_lib "app/lib/api/lib"
	"net/http"
)

func Init(mux *http.ServeMux) {
	Api_lib.Rest(mux, "/").Get(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		res.Status(200).Send("hi")
	}).Apply()
}
