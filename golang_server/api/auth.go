package api

import (
	"idk/main/jwt"
	"net/http"
	"strings"
)

func AuthorizationJwt(w *http.ResponseWriter, r **http.Request) interface{} {
	token := strings.TrimPrefix((*r).Header.Get("Authorization"), "Bearer ")
	if token == "" {
		(*w).WriteHeader(500)
		(*w).Write([]byte("Provide Authorization token"))
		return nil
	}
	data, err := jwt.Use().Verify(token)
	if err != nil {
		(*w).WriteHeader(500)
		(*w).Write([]byte(err.Error()))
		return nil
	}
	return data
}
