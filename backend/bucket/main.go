package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
)

func main() {
	mux := http.NewServeMux()
	Api_lib.Rest(mux, "/").Get(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		file, err := os.ReadFile(path.Join("./static/", r.URL.Path))
		if err != nil {
			res.Status(404).Send("File not found")
			return
		}
		res.Status(200).Send(file)
		w.Header().Set("Content-Type", "application/octet-stream")
	}).Post(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		url_path := strings.Split(r.URL.Path, "/")
		err := os.MkdirAll(path.Join("./static/", strings.Join(url_path[:len(url_path)-1], "/")), os.ModePerm)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("file")
		if err != nil {
			res.Status(400).Send(err.Error())
			return
		}
		defer file.Close()
		filebytes, err := io.ReadAll(file)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		err = os.WriteFile(path.Join("./static/", r.URL.Path), filebytes, os.ModePerm)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send("ok")
	}).Delete(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		err := os.Remove(path.Join("./static/", r.URL.Path))
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send("ok")
	}).Apply()
	http.ListenAndServe(":8080", mux)
}
