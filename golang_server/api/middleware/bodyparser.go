package api_middleware

import (
	"io"
	"net/http"
)

func BodyParser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return nil
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Please provide body"))
		return err
	}
	q := r.URL.Query()
	q.Add("body", string(body))
	r.URL.RawQuery = q.Encode()
	return nil
}
