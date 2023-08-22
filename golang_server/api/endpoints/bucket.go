package endpoints

import (
	Http "idk/main/api/lib"
	api_middleware "idk/main/api/middleware"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Bucket(mux *http.ServeMux) {
	cluster := Http.Cluster(mux, "/bucket")
	cluster.Rest("/").Use(api_middleware.AllowCors).Use(api_middleware.Auth).Post(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		elems := strings.Split(r.URL.Path, "/")
		if elems[len(elems)-2] != "bucket" {
			res.Status(404).Send("")
			return
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			res.Status(400).Send("Provide file")
			return
		}
		filePath := filepath.Join("bucket/static/", r.URL.Query().Get("Authorization"))
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		filePath = filepath.Join(filePath, elems[len(elems)-1])
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		defer file.Close()
		defer f.Close()
		io.Copy(f, file)
		res.Status(200).Send("ok")
	}).Get(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		elems := strings.Split(r.URL.Path, "/")
		if elems[len(elems)-2] != "bucket" {
			res.Status(404).Send("")
			return
		}
		content, err := os.ReadFile(filepath.Join("bucket/static/", r.URL.Query().Get("Authorization"), elems[len(elems)-1]))
		if err != nil {
			res.Status(404).Send("Bucket not found")
			return
		}
		res.Status(200).AddHeader("Content-Type", "application/octet-stream").Send(content)
	}).Apply()
}
