package api_middleware

import (
	"errors"
	"fmt"
	"idk/main/jwt"
	"net/http"
	"strings"
)

func Auth(w http.ResponseWriter, r *http.Request) error {
	token := strings.TrimPrefix((*r).Header.Get("Authorization"), "Bearer ")
	if token == "" {
		w.WriteHeader(401)
		w.Write([]byte("Provide Authorization token"))
		return errors.New("Unauthorized")
	}
	data, err := jwt.Use().Verify(token)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Jwt expired"))
		return err
	}
	q := r.URL.Query()
	q.Add("Authorization", fmt.Sprint(data))
	r.URL.RawQuery = q.Encode()
	return nil
}
