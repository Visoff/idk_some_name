package endpoints

import (
	Http "idk/main/api/lib"
	api_middleware "idk/main/api/middleware"
	"idk/main/bucket"
	"io"
	"net/http"
	"strings"
)

func Bucket(mux *http.ServeMux) {
	cluster := Http.Cluster(mux, "/bucket")
	cluster.Rest("/").Use(api_middleware.AllowCors).Post(func(w http.ResponseWriter, r *http.Request) {
		elems := strings.Split(r.URL.Path, "/")
		if elems[len(elems)-2] != "bucket" {
			w.WriteHeader(404)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Body was not provided"))
			return
		}
		err = bucket.Write(elems[len(elems)-1], body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	}).Get(func(w http.ResponseWriter, r *http.Request) {
		elems := strings.Split(r.URL.Path, "/")
		if elems[len(elems)-2] != "bucket" {
			w.WriteHeader(404)
			return
		}
		content, err := bucket.Read(elems[len(elems)-1])
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Bucket not found"))
			return
		}
		w.WriteHeader(200)
		w.Write(content)
	}).Apply()
	cluster.Rest("/hello").Use(api_middleware.AllowCors).Get(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	}).Apply()
}
