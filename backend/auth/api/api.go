package api

import (
	"net/http"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
)

func Init(mux *http.ServeMux) {
	Api_lib.Rest(mux, "/").Get(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		res.Status(200).Send("hi")
	}).Apply()
}
