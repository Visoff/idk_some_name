package endpoints

import (
	Http "idk/main/api/lib"
	api_middleware "idk/main/api/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Bucket(mux *http.ServeMux) {
	cluster := Http.Cluster(mux, "/bucket")
	cluster.Rest("/dir/").Use(api_middleware.AllowCors).Use(api_middleware.Auth).Get(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		user := r.URL.Query().Get("Authorization")
		url_path := strings.Split(r.URL.Path, "/")
		var path string
		for i, el := range url_path {
			if el == "dir" {
				path = strings.Join(url_path[i+1:], "/")
				break
			}
		}
		path = filepath.Join("bucket/static", user, path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		dir, err := os.ReadDir(path)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		var resp []map[string]interface{}
		for _, file := range dir {
			f := make(map[string]interface{})
			f["IsDir"] = file.IsDir()
			f["Name"] = file.Name()
			resp = append(resp, f)
		}
		res.Status(200).Send(resp)
	}).Apply()
	cluster.Rest("/file/").Use(api_middleware.AllowCors).Use(api_middleware.Auth).Get(func(w http.ResponseWriter, r *http.Request) {
		res := Http.NewBetterResponseWriter(w)
		user := r.URL.Query().Get("Authorization")
		url_path := strings.Split(r.URL.Path, "/")
		var path string
		for i, el := range url_path {
			if el == "file" {
				path = strings.Join(url_path[i+1:], "/")
				break
			}
		}
		path = filepath.Join("bucket/static", user, path)
		content, err := os.ReadFile(path)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send(content)
	}).Apply()
}
